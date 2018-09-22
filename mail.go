package main

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

const (
	SMTPServer = "smtp.gmail.com"
)

type Account struct{
	User string
	Password string
}

func NewAccount(Username, Password string) Account{
	return Account{Username,Password}
}

func (account Account) SendMail(Dest []string,Subject, bodyMessage string){
	msg := "From : " + account.User + "\n" +
			"To : " + strings.Join(Dest, ",") + "\n" +
			"Subject :" + Subject + "\n" + bodyMessage

	err := 	smtp.SendMail(SMTPServer+":587",
			smtp.PlainAuth("",account.User,account.Password,SMTPServer),
			account.User, Dest, []byte(msg))

	if err != nil{
		fmt.Printf("Terjadi kesalahan pada smtp : %s",err)
		return
	}

	fmt.Println("Email berhasil terkirim")

}

func (account Account) WriteEmail(dest []string,contentType,subject,bodyMessage string) string{
	header := make(map[string]string)
	header["From"] = account.User

	penerima := ""

	for _,user := range dest{
		penerima = penerima + user
	}

	header["To"] = penerima
	header["Subject"] = subject
	header["MIME-version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"",contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key,value := range header{
		message += fmt.Sprintf("%s: %s\r\n",key,value)
	}

	var encodedMessage bytes.Buffer

	isiPesan := quotedprintable.NewWriter(&encodedMessage)
	isiPesan.Write([] byte(bodyMessage))
	isiPesan.Close()

	message += "\r\n" + encodedMessage.String()

	return message
	
}

func (account *Account) KirimHtmlEmail(dest []string, subject,bodyMessage string) string{
	return account.WriteEmail(dest,"text/html",subject,bodyMessage)
}

func (account *Account) KirimTextEmail(dest []string, subject,bodyMessage string) string{
	return account.WriteEmail(dest,"text/plain",subject,bodyMessage)
}