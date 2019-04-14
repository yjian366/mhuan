package Utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var g_loger *log.Logger

func init() {

	f,err:=os.Create("TMS"+time.Now().Format("20060102150405")+".log")
	if err!=nil{
		log.SetPrefix("TMS:")
		log.SetFlags(log.LstdFlags|log.Llongfile)
		log.Println(err)
		return
	}
	w:=io.MultiWriter(f,os.Stdout)

	g_loger = log.New(w, "[TMS] ", log.LstdFlags|log.Llongfile)


}

func GetLoger() *log.Logger{
	return g_loger
}

func Debug(format string,args ...interface{}) {
	if g_loger != nil {
		g_loger.SetPrefix(fmt.Sprintf("TMS-%s:","[Debug]"))
		g_loger.Output(2, fmt.Sprintf(format, args...))
	}
}

func Info(format string,args ...interface{}) {
	if g_loger != nil {
		g_loger.SetPrefix(fmt.Sprintf("TMS-%s:","[Infos]"))
		g_loger.Output(2, fmt.Sprintf(format, args...))
	}
}

func Error(format string, args ...interface{}) {
	if g_loger != nil {
		g_loger.SetPrefix(fmt.Sprintf("TMS-%s:","[Error]"))
		g_loger.Output(2, fmt.Sprintf(format, args...))
	}
}
func Warn(format string,args ...interface{}) {
	if g_loger != nil {
		g_loger.SetPrefix(fmt.Sprintf("TMS-%s:","[Warns]"))
		g_loger.Output(2, fmt.Sprintf(format, args...))
	}
}
