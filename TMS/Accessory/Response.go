package Utils

import (
	json2 "encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http"
)

type Answer struct{

	Code 		int				`json:"code"`
	Descript	string			`json:"descript"`
	More      	interface{} 	`json:"more"`
	status		int
}

func(this Answer)SetCode(arg int){
	this.Code=arg
}

func(this Answer)SetDescript(arg string){
	this.Descript=arg
}

func(this Answer)SetDetail(arg interface{}){
	this.More=arg
}

func(this Answer)GenMap()(h gin.H){
	value,err:=json.Marshal(this)
	if err!=nil{
		h=gin.H{"reason":"the Answer Ouput Mehtod Marshal data error"}
		return
	}


	err=json2.Unmarshal(value,&h)
	if err!=nil{
		h=gin.H{"reason":"the Answer Ouput Mehtod Marshal data error"}
		return
	}
	return

}

func(this Answer)OutputJson(c *gin.Context){
	if this.status==0{
		this.status=http.StatusOK
	}

	c.JSON(this.status,this.GenMap())
}
func(this Answer)OutputString(c *gin.Context){
	if this.status==0{
		this.status=http.StatusOK
	}

	c.JSON(this.status,this.Descript)

}

func(this Answer)OutputXML(c *gin.Context){
	if this.status==0{
		this.status=http.StatusOK
	}

	value,err:=xml.Marshal(this)
	if err!=nil{
		c.XML(this.status,`<map>the Answer Ouput Mehtod Marshal data error</map>`)

	}
	var v string
	err=xml.Unmarshal(value,&v)
	if err!=nil {
		c.XML(this.status, `<map>the Answer Ouput Mehtod Marshal data error</map>`)
		c.XML(this.status,`<map>the Answer Ouput Mehtod Marshal data error</map>`)
	}
	c.XML(this.status,this)
}
