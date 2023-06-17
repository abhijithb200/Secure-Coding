import sys, requests, json, urllib, os,time
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoAlertPresentException, NoSuchElementException, ElementNotInteractableException, TimeoutException
from selenium.webdriver.common.keys import Keys
from selenium.common.exceptions import StaleElementReferenceException
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions
import json
import json,urllib.request
from email_send import email_sender
import subprocess
import mysql.connector

def get_payloads_from_vectors():
    payloads = []

    with open('constants/vectors.txt', 'r', encoding = "utf-8") as vector_file:
        for vector in vector_file.readlines():
            payloads.append(vector)


    return payloads

def encode_url(url, params):
    params_encoded = urllib.parse.urlencode(params)
    full_url = url + "?" + params_encoded
    return full_url

def get_base_url(url):
    base_url = url.split('?')[0]
    return base_url


class Scanner:
    def __init__(self):
        self.payloads = get_payloads_from_vectors()
        self.url = "http://localhost/dashboard/phpinfo.php"
        self.codeguardianurl = "http://localhost/dashboard/"
        
        self.base_url = get_base_url(self.url)
        self.source = self.set_source()

    def set_source(self):
        source = urllib.request.urlopen(self.codeguardianurl+"/Codeguardian.json").read()
        source = json.loads(source)

        return source

class XSS_Scanner(Scanner):
    def __init__(self) :
        super().__init__()
        self.params = self.get_params()
        

    def setup_windows(self):
        query_window = self.driver.current_window_handle
        return query_window

    def get_params(self):
        params = {}
        for i in self.source["vulns"]:
            if "Reflected XSS" in  i["type"]:
                s = i["source"]["value"]
                a = s.replace("'","")
                a = s.replace('"',"")
                params[a] = ""
        
        return params
    
    def get_hash(self):

        return self.source["hash"]

    def query_scanner(self, payload):
        for param in self.params.keys():
            previous_value = self.params[param]
            self.params[param] = payload
            target_url = encode_url(self.base_url, self.params)
            self.raw_params = urllib.parse.urlencode(self.params)


            self.driver.get(target_url)
            try:
                WebDriverWait(self.driver, 1).until(expected_conditions.alert_is_present())
                self.driver.switch_to.alert.accept()
                return target_url
            except TimeoutException:
                pass

    def setup(self):
        options = webdriver.ChromeOptions()
        options.add_argument('--headless')
        options.add_argument('--no-sandbox')
        self.driver = webdriver.Chrome(options=options)
        self.query_window = self.setup_windows()

    def run_on_url(self):
        print('[*] Running XSS Scan [*]')

        for payload in self.payloads:
            if payload != "\n":
                payload = payload[:-1]
                self.driver.switch_to.window(self.query_window)
                
                vuln_url = self.query_scanner(payload)
                if vuln_url:
                    print("Detected",vuln_url)
                    # email_sender(vuln_url) 
                    break

        return
    
class SQL_Scanner(Scanner):

    def __init__(self):
        super().__init__()
        self.params = self.get_params()
        self.sqlmapPath = "C://Users//abhij//sqlmap//sqlmap.py"
        self.dbconn = mysql.connector.connect(host = "localhost", user = "root",passwd = "",database = "test")
        self.tablename,self.tablefields = self.get_dbdetails()
        self.setup()

    def checkTableExists(self):
        dbcur = self.dbconn.cursor()
        dbcur.execute("SHOW tables;")
        result = dbcur.fetchall()
        for x in result:
            if x[0]==self.tablename:
                return True
        dbcur.close()
        return False
    
    def checkReqFields(self):
        dbcur = self.dbconn.cursor()
        dbcur.execute("SHOW COLUMNS FROM persons;")
        result = dbcur.fetchall()

        fields_present = []
        fields_not_present = []
        for x in result:
            fields_present.append(x[0].lower())

        for i in self.tablefields:
            if i not in fields_present:
                fields_not_present.append(i)

        dbcur.close()
        return fields_not_present


        
    

    def get_dbdetails(self):
        for i in self.source["vulns"]:
            if "SQL Injection" in  i["type"]:
                dbdetails = i["source"]["dbdetails"]
                for k,v in dbdetails.items():
                    return k,[i.lower() for i in v]

    def setup(self):
        '''
        if table exist - check is the table contain the required fields - and any data
        if table not exist - create a table with fields and data
        if table exist - and have required fields - add data
        '''
        dbcur = self.dbconn.cursor()
        if self.checkTableExists():
            c = self.checkReqFields()
            if c != [] : # fields not preset in table
                
                for i in c:
                    dbcur.execute(f"ALTER TABLE {self.tablename} \
                                  ADD {i} VARCHAR(100);",)
                    
                query = 'INSERT INTO persons '+'(' +','.join(c)+')'+' VALUES '+ '(' +','.join([ "\""+5*"A"+"\"" for i in range(len(c))])+');'

                dbcur.execute(query)
                self.dbconn.commit()
                dbcur.close()
                print("Added fields")
        else:   #if table not present
            query = f'CREATE TABLE {self.tablename} ' + ' (' + ' VARCHAR(255), '.join(self.tablefields) + " VARCHAR(255));"
            dbcur.execute(query)

            query = f'INSERT INTO {self.tablename} '+'(' +','.join(self.tablefields)+')'+' VALUES '+ '(' +','.join([ "\""+5*"A"+"\"" for i in range(len(self.tablefields))])+');'
            print(query)
            dbcur.execute(query)
            self.dbconn.commit()
            print("Table created")
            dbcur.close()


    def run_on_url(self):
        sqlmap_cmd = ["python",self.sqlmapPath, "-u", "http://localhost/test/sql.php?id=2", "-p", "id", "--batch","--flush-session"]
        result = subprocess.run(sqlmap_cmd, capture_output=True, text=True)
        output = result.stdout.strip()
        if "does not seem to be injectable" not in output:
            print("Vulnerability found",output)
        else:
            print("No vulnerability")


        pass

    def get_params(self):
        params = {}
        for i in self.source["vulns"]:
            if "SQL Injection" in  i["type"]:
                s = i["source"]["value"]
                a = s.replace("'","")
                a = s.replace('"',"")
                params[a] = ""

        return params

# xss = XSS_Scanner()
# xss.setup()
# xss.run_on_url()

sql = SQL_Scanner()
sql.run_on_url()

  

# while True:
#     if os.getenv("GITHASH") != s.get_hash(): #changed
#         print("Changed")
#         os.environ["GITHASH"] = s.get_hash()
#         s.run_on_url()
#     print("Not changed")
#     time.sleep(5)

