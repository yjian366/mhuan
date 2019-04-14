package Rounter

import (
	"TMS/Bin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine){



	api:=g.Group("/api")
	v1:=api.Group("/v1")
	v1.POST("/login", Bin.Login)
	v1.GET("/logout",Bin.Logout)

	user:=v1.Group("/user")
	user.GET("department/:department",Bin.QueryDepartmentMember)
	user.GET("department/:department/role/:role",Bin.QueryDepartmentRoleMember)
	user.GET("team/:team",Bin.QueryTeamMember)
	user.GET("team/:team/role/:role",Bin.QueryTeamRoleMember)
	user.GET("personal/:userInfo",Bin.QueryUser)


	task:=v1.Group("/task")
	task.POST("/create",Bin.CreateTask)
	task.GET("/all",Bin.GetAllTasks)
	task.GET("/creater/:name",Bin.GetTasksByCreater)
	task.GET("/executor/:name",Bin.GetTasksByExecutor)
	task.GET("/module/:name",Bin.GetTasksByModule)
	task.GET("/system_id/:id",Bin.GetTasksBySystemID)
	task.GET("/name/:name",Bin.GetTasksByTaskName)
	task.GET("/type/:name",Bin.GetTasksByTaskType)
	task.GET("/status/:status",Bin.GetTasksByTaskStatus)

	task.GET("/great_time/:time",Bin.GetTasksByTaskGreatTime)
	task.GET("/less_time/:time",Bin.GetTasksByTaskLessTime)
	task.GET("/great_equal_time/:time",Bin.GetTasksByTaskGreatETime)
	task.GET("/less_equal_time/:time",Bin.GetTasksByTaskLessETime)


	mail:=v1.Group("/mail")
	mail.POST("create",Bin.SendMail)



	g.Handle("POST","/createTask", Bin.CreateTaskNoAuth)
	task.POST("/testcreate", func(c *gin.Context) {
		v,err:=c.GetRawData()
		if err!=nil{
			fmt.Println(err)
		}
		s,_:=Bin.GBKEncode(v)
		fmt.Println("212312")
		fmt.Println(string(s))
		c.String(200,string(s))
		return
	})

	}



