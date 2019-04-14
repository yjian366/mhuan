package Component

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type TestLeader struct {
	Tester
}

func(this *TestLeader)CreateTask(args Task){

	sess,err:=mgo.Dial("localhost")

	defer sess.Close()
	if err!=nil{
		fmt.Println(err)
		return
	}

	c:=sess.DB("TMS").C("Tasks")
	fmt.Println(c.Insert(&args))

}

func(this *TestLeader)DeleteTask(){

}
