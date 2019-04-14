package mail

import (
	"crypto/tls"
)
type MailMessage struct{
	To struct{
		Mail string
		Name string
	}
	Subject string

	Body struct{
		TaskName string
		Date string
		TaskCreater string
	}
}
type Mail struct {

	account	string
	password string
	host string
	port int
	mess *Message
	dial *Dialer

	message struct{
		from    struct{
			mail string
			name string
		}
		to  struct{
			mail string
			name string
		}
		cc          []string
		bcc         []string
		subject     string
		body        string
		contentType string
	}
}
func(this *Mail)SetAccount(account string){
	this.account=account
}

func(this *Mail)GetAccount() string{
	return this.account
}

func(this *Mail)SetPassword(pass string){
	this.password=pass
}

func(this *Mail)GetPassword()string{
	return this.password
}

func(this *Mail)SetHost(host string){
	this.host=host
}

func(this *Mail)GetHost()string{
	return this.host
}

func(this *Mail)SetPort(port int){
	this.port=port
}

func(this *Mail)GetPort()int{
	return this.port
}

func(this *Mail)SetToMail(mail string){
	this.message.to.mail=mail
}
func(this *Mail)GetToMail()string{
	return this.message.to.mail
}

func(this *Mail)SetToName(name string){
	this.message.to.name=name
}

func(this *Mail)GetToName()string{
	return this.message.to.name
}

func(this *Mail)GetSubject()string{
	return this.message.subject
}

func(this *Mail)SetSubject(subject string){
	this.message.subject=subject
}

func(this *Mail)GetBody()string{
	return this.message.body
}

func(this *Mail)SetBody(body string){
	this.message.body=body
}

func(this *Mail)GetContentType()string{
	return this.message.contentType
}

func(this *Mail)SetContentType(contenttype string){
	this.message.contentType=contenttype
}

func New()*Mail{

	var m Mail
	m.account="rd.shared@tcl.com"
	m.password="tmt#0228"
	m.message.from.mail=m.account
	m.message.from.name="TMS"
	m.message.contentType="text/html"
	m.mess=NewMessage()
	m.host="mail.tcl.com"
	m.port=25
	m.dial=NewDialer(m.host,m.port,m.account,m.password)
	return &m
}

func(this *Mail)dialAndSend()(err error){
	if this.dial==nil{
		this.dial=NewDialer(this.host,this.port,this.account,this.password)
	}
	this.dial.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err = this.dial.DialAndSend(this.mess)
	return
}

func(this *Mail)Send()(err error){

	this.mess.SetAddressHeader("From",this.message.from.mail,this.message.from.name)
	this.mess.SetAddressHeader("To",this.message.to.mail,this.message.to.name)
	this.mess.SetHeader("Subject",this.message.subject)
	this.mess.SetBody(this.message.contentType,this.message.body)
	err=this.dialAndSend()
	return
}