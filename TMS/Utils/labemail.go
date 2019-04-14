package utils
//
//import (
//	"gopkg.in/gomail.v2"
//	"waterdroplet/Service/config"
//	"waterdroplet/Service/common"
//)
//
//func SendMail(To []string, Cc []string, Bcc []string, Subject string, Body string) error {
//	m := gomail.NewMessage()
//	m.SetHeader("From", config.Cfg.EmailUser)
//	receiver := make([]string, len(To))
//	for i, to := range To {
//		receiver[i] = m.FormatAddress(to, to)
//	}
//	if len(receiver) > 0 {
//		m.SetHeader("To", receiver[:]...)
//	}
//
//	receiver = make([]string, len(Cc))
//	for i, cc := range Cc {
//		receiver[i] = m.FormatAddress(cc, cc)
//	}
//	if len(receiver) > 0 {
//		m.SetHeader("Cc", receiver[:]...)
//	}
//
//	receiver = make([]string, len(Bcc))
//	for i, bcc := range Bcc {
//		receiver[i] = m.FormatAddress(bcc, bcc)
//	}
//	if len(receiver) > 0 {
//		m.SetHeader("Bcc", receiver[:]...)
//	}
//
//	m.SetHeader("Subject", Subject)
//	m.SetBody("text/html", Body)
//	//m.Attach("/home/Alex/logcat.jpg")
//
//	d := gomail.NewDialer(config.Cfg.SmtpServer, config.Cfg.SmtpPort, config.Cfg.EmailUser, config.Cfg.EmailPassword)
//
//	err := d.DialAndSend(m)
//	if err != nil {
//		common.Info("Email send fail:", err)
//	}
//
//	return err
//}