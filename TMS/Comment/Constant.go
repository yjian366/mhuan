package Comment

import (
	"fmt"
	"log"
	"os"
	"time"
)

type tmsLog struct{
	logFile *log.Logger
	logShow *log.Logger
}

var GlobalLog tmsLog

func init(){
	GlobalLog=tmsLog{}
	f,err:=os.Create("TMS"+time.Now().Format("20060102150405")+".log")
	if err!=nil{
		log.SetPrefix("TMS:")
		log.SetFlags(log.LstdFlags|log.Llongfile)
		log.Println(err)
		return
	}

	GlobalLog.logFile=log.New(f,"TMS:",log.LstdFlags|log.Llongfile)
	GlobalLog.logShow=log.New(os.Stdout,"TMS:",log.LstdFlags|log.Llongfile)
}

func(this *tmsLog)Debug(args string){
	GlobalLog.logFile.SetPrefix(fmt.Sprintf("TMS-%s:","[Debug]"))
	GlobalLog.logFile.Println(args)

	GlobalLog.logShow.SetPrefix(fmt.Sprintf("TMS-%s:","[Debug]"))
	GlobalLog.logShow.Println(args)
}

func(this *tmsLog)Error(args string){
	GlobalLog.logFile.SetPrefix(fmt.Sprintf("TMS-%s:","[Error]"))
	GlobalLog.logFile.Println(args)

	GlobalLog.logShow.SetPrefix(fmt.Sprintf("TMS-%s:","[Error]"))
	GlobalLog.logShow.Println(args)
}

func(this *tmsLog)Info(args string){
	GlobalLog.logFile.SetPrefix(fmt.Sprintf("TMS-%s:","[Info]"))
	GlobalLog.logFile.Println(args)

	GlobalLog.logShow.SetPrefix(fmt.Sprintf("TMS-%s:","[Info]"))
	GlobalLog.logShow.Println(args)
}

func(this *tmsLog)Panic(args string){
	GlobalLog.logFile.SetPrefix(fmt.Sprintf("TMS-%s:","[Panic]"))
	GlobalLog.logFile.Panic(args)

	GlobalLog.logShow.SetPrefix(fmt.Sprintf("TMS-%s:","[Panic]"))
	GlobalLog.logShow.Panic(args)
}

func(this *tmsLog)Fatal(args string){
	GlobalLog.logFile.SetPrefix(fmt.Sprintf("TMS-%s:","[Fatal]"))
	GlobalLog.logFile.Fatalln(args)

	GlobalLog.logShow.SetPrefix(fmt.Sprintf("TMS-%s:","[Fatal]"))
	GlobalLog.logShow.Fatalln(args)
}