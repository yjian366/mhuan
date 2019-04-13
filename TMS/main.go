package main

import (
	"TMS/Accessory/Mail"
	"TMS/Rounter"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)


type MyHandler struct{}

func (this MyHandler)ServeHTTP(w http.ResponseWriter,r *http.Request) {

	if r.Method==http.MethodGet{
		w.Write([]byte("this is GET request"))
	}else{
		w.Write([]byte("the not GET request"))
	}
}
type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}
func  sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		log.Println(c.Request.Method)
		log.Println(c.Request.Header)
		log.Println(c.Keys)

		log.Println("the before request")
		c.Next()
		log.Println("the after request")
		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}


type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
type authTicketEntry struct {
	TickerName string `json:"ticketName"`
	TicketValue string `json:"ticketValue"`
}

type authUser struct {
	Uid string `json:"uid"`
	Mail string `json:"mail"`
	Mobile string `json:"mobile"`
}

type authBody struct {
	TicketEntry authTicketEntry `json:"ticketEntry"`
	User authUser `json:"user"`
}
type authRes struct {
	ErrorCode int `json:"ErrorCode"`
	ErrorDescription string `json:"ErrorDescription"`
	Body authBody `json:"Body"`
}
func LoginByRd(username string, password string) (mail string, err error) {
	url := "http://rd.tmt.tcl.com/api/v2/login"
	var jsonStr = []byte(`{"UserName":"` + username + `","Password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	res := authRes{}
	json.Unmarshal(body, &res)
	if res.Body.User.Mail == "" {
		return "", errors.New("empty email")
	} else {
		return res.Body.User.Mail, nil
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")

			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}


		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

func TestMail(){
	m := mail.NewMessage()
	m.SetAddressHeader("From", "j.yang@tcl.com", "杨建")
	m.SetAddressHeader("To", "j.yang@tcl.com", "段平")
	m.SetHeader("Subject", "TMS邮件测试")
	m.SetBody("text/html", `<h1 style='color:red'>hello world</h1>
							<button onclick="alert('fuck')">click me </button>
							<button onclick="alert('this')">click me</button>
							<a style='color:green' href="www.baidu.com">this is </a>
							`)

	d := mail.NewDialer("mail.tcl.com", 25, "j.yang@tcl.com", "loves521.")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("***%s\n", err.Error())
	}

}
func main() {

	router:=gin.Default()
	router.LoadHTMLGlob(filepath.Join(`.\Test\login.html`))
	router.Use(Cors())
	Rounter.InitRouter(router)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.Run(":8081")







	//var abc Utils.PriorityConstructor
	//
	//abc.MTK_GROUP=Comment.TESTER
	//abc.RTK_GROUP=Comment.TESTLEADER
	//abc.MSTAR_GROUP=Comment.TESTLEADER
	//fmt.Println(abc.GenCode())
	//
	//var adb Utils.PriorityConstructor
	//adb.ParseCode(abc.GenCode())
	//
	//fmt.Println(adb.RTK_GROUP)
	//fmt.Println(adb.MTK_GROUP)
	//fmt.Println(adb.MSTAR_GROUP)
	//fmt.Println(adb.Certificaty_GROUP)









}

	//x:=Utils.Token{}
	//y:=Utils.Token{}
	//
	//x.Data.Email="j.yang"
	//x.Data.Time=time.Now()
	//x.Data.Priority=777
	//
	//r:=gin.Default()
	//r.GET("/abc", func(c *gin.Context) {
	//	fmt.Println(x.WriteToCookie(c))
	//})
	//
	//r.GET("/adb", func(c *gin.Context) {
	//
	//	fmt.Println(y.ReadFromCookie(c))
	//
	//	log.Println(y)
	//})
	//
	//r.GET("/c", func(c *gin.Context) {
	//
	//	fmt.Println(y.ClearCookie(c))
	//})
	//r.GET("/u", func(c *gin.Context) {
	//
	//	fmt.Println(y.UpdateToken(c))
	//})
	//
	//r.Run(":8080")




