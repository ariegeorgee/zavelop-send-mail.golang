package main

func main(){

	// SETTING MAIL SMTP SERVER
	mail := NewAccount("YOUR EMAIL","YOUR PASSWORD")

	// SETING RECEPIENT
	// SAMPLE Receiver := []string{"abc@sample.com"}
	Receiver := []string{"MAIL TO"}

	Subject := "Percobaan mengirim Email menggunakan GoLang"
	message := `
				<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
				<html>
				<head>
				<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
				</head>
				<body>This is the body<br>
				<div class="moz-signature"><i><br>
				<br>
				Regards<br>
				Ari Ardiansyah<br>
				081287814488<br>
				ariardiansyah470@gmail.com<br>
				<i></div>
				</body>
				</html>
				`
	// SET EMAIL ATTRIBUTE
	bodyMessage := mail.KirimHtmlEmail(Receiver,Subject,message)

	// KIRIM EMAIL
	mail.SendMail(Receiver,Subject,bodyMessage)
}