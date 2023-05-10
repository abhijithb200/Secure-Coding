import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.mime.base import MIMEBase 
from email import encoders 

def email_sender(link):
    #The mail addresses and  password
    sender_address ="abhijithbinoy11@gmail.com"                  
    sender_pass ="qrxvgdyumisgekbk"                                       
    reciver_mail="gokulgr555@gmail.com"

    # list of reciver email_id to the mail 
                                                   
    #[item for item in input("Enter Receiver Mail Address :- ").split()] this is used to take user input of receiver mail id


    # getting length of list 

   
    # Iterating the index 
    # same as 'for i in range(len(list))' 

    # Here we iterate the loop and send msg one by one to the reciver


    message = MIMEMultipart()
    message['From'] = sender_address
    message['To'] =  reciver_mail             #  Pass Reciver Mail Address
    message['Subject'] =  'Vulnerability Found'       #The subject line
    

    mail_content = f'''Hello,
    XSS Vulnerability Found on : {link}'''


    #The body and the attachments for the mail
    message.attach(MIMEText(mail_content, 'plain'))
    s = smtplib.SMTP('smtp.gmail.com', 587) 
    s.starttls() 
    s.login(sender_address, sender_pass) 
    text = message.as_string()
    s.sendmail(sender_address, reciver_mail, text) 
    s.quit() 
    print('Mail Sent to',reciver_mail)


