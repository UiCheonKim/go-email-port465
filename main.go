package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	subj := "test1"
	body := "test2"

	from := "drooom@naver.com"
	to := "drooom@naver.com"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "servername" // ex) mail.corgi.co.kr:465

	host := "host" // ex) mail.corgi.co.kr

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}
	/*
	  // Auth
	  if err = c.Auth(auth); err != nil {
	      log.Panic(err)
	  }
	*/
	if err := c.Auth(smtp.PlainAuth("", "username", "password", host)); err != nil {
		log.Panic(err)
	}

	// To && From
	/*
	  if err = c.Mail(from.Address); err != nil {
	      log.Panic(err)
	  }

	  if err = c.Rcpt(to.Address); err != nil {
	      log.Panic(err)
	  }
	*/

	if err = c.Mail("drooom@naver.com"); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt("drooom@naver.com"); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	err = c.Quit()
	if err != nil {
		log.Panic(err)
	}
}
