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
    def __init__(self, url):
        self.payloads = get_payloads_from_vectors()
        self.url = url
        self.codeguardianurl = os.getenv("Codeguarian")
        self.params = self.get_params()
        self.base_url = get_base_url(self.url)
      

    def setup_windows(self):
        query_window = self.driver.current_window_handle
        return query_window

    def get_params(self):
        source = urllib.request.urlopen(self.codeguardianurl+"/Codeguardian.json").read()
        source = json.loads(source)
        params = {}
        for i in source["vulns"]:
            if "Reflected XSS" in  i["type"]:
                s = i["source"]["value"]
                a = s.replace("'","")
                params[a] = ""
        return params
    
    def get_hash(self):
        source = urllib.request.urlopen(self.codeguardianurl+"/Codeguardian.json").read()
        source = json.loads(source)

        return source["hash"]

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
                    break

        return

s = Scanner("http://codeguardian-serv/index.php")
s.setup()

while True:
    if os.getenv("GITHASH") != s.get_hash(): #changed
        print("Changed")
        os.environ["GITHASH"] = s.get_hash()
        s.run_on_url()
    print("Not changed")
    time.sleep(5)