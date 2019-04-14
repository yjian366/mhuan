package Component

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Tester struct{
	User
	Name string
	Department string
	Team string
	Title string
	Priority int64
	Mobile string
}

func(this *Tester)GetTasks(){

	sess,err:=mgo.Dial("localhost")
	if err!=nil{
		fmt.Println(err)
		return
	}

	c:=sess.DB("TMS").C("Tasks")
	len,err:=c.Count()
	fmt.Println(len)
	var content Task
	c.Find(bson.M{"name":bson.M{"$eq":this.Name}}).One(&content)

	fmt.Println(content)


}

func(this *Tester)UpdateTasks(){

}

func(this *Tester)QueryTask(){

}