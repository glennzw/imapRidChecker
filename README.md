Small utility to test the gophish IMAP campaign detection logic. eg:
 
```bash
➜  go run imapRidChecker.go ~/Downloads/email/Fw_\ Outlook\ size\ limit\ exceeded.\ outlook.com.eml 
[+] Checking email from 'Jonathan McCarthy <user@internet.com>' with subject 'Fw: Outlook size limit exceeded.'
[+] This is a campign email. Found the following gophish rid parameters: 5Stzk5o

➜  src go run imapRidChecker.go ~/Downloads/email/FW_\ Outlook\ size\ limit\ exceeded.\ Desktop.eml    
[+] Checking email from 'Jonathan McCarthy <user@internet.com>' with subject 'FW: Outlook size limit exceeded.'
[+] This is a campign email. Found the following gophish rid parameters: 5Stzk5o

➜  src go run imapRidChecker.go ~/Downloads/Help\ strengthen\ the\ security\ of\ your\ Google\ Account.eml
[+] Checking email from 'Google <no-reply@accounts.google.com>' with subject 'Help strengthen the security of your Google Account'
[!] No rid parameters found, not a campaign email!
```
