// This code can be used to check if your gophish campaign email will be picked up by the IMAP monitor.
// e.g:
/*
	➜  go run imapRidChecker.go ~/Downloads/email/Fw_\ Outlook\ size\ limit\ exceeded.\ outlook.com.eml 
	[+] Checking email from 'Bob <bob@internet.com>' with subject 'Fw: Outlook size limit exceeded.'
	[+] This is a campign email. Found the following gophish rid parameters: 5Stzk5o
*/

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/jordan-wright/email"
)

var goPhishRegex = regexp.MustCompile("((\\?|%3F)rid(=|%3D)(3D)?([A-Za-z0-9]{7}))")

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please supply .eml file")
		os.Exit(1)
	}
	rawBodyStream, err := os.Open(os.Args[1])
	m, err := email.NewEmailFromReader(rawBodyStream)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file. Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[+] Checking email from '%s' with subject '%s'\n", m.From, m.Subject)
	rids, err := matchEmail(m) // Search email Text, HTML, and each attachment for rid parameters

	if len(rids) > 0 {
		fmt.Println("[+] This is a campign email. Found the following gophish rid parameters:")
		for rid, _ := range rids {
			fmt.Println(rid)
		}
	} else {
		fmt.Println("[!] No rid parameters found, not a campaign email!")
	}
}

func checkRIDs(em *email.Email, rids map[string]bool) {
	// Check Text and HTML
	emailContent := string(em.Text) + string(em.HTML)
	for _, r := range goPhishRegex.FindAllStringSubmatch(emailContent, -1) {
		newrid := r[len(r)-1]
		if !rids[newrid] {
			rids[newrid] = true
		}
	}
}

// returns a slice of gophish rid paramters found in the email HTML, Text, and attachments
func matchEmail(em *email.Email) (map[string]bool, error) {
	rids := make(map[string]bool)
	checkRIDs(em, rids)

	// Next check each attachment
	for _, a := range em.Attachments {
		ext := filepath.Ext(a.Filename)
		if a.Header.Get("Content-Type") == "message/rfc822" || ext == ".eml" {

			// Let's decode the email
			rawBodyStream := bytes.NewReader(a.Content)
			attachmentEmail, err := email.NewEmailFromReader(rawBodyStream)
			if err != nil {
				return rids, err
			}

			checkRIDs(attachmentEmail, rids)
		}
	}

	return rids, nil
}
