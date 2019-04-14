package Utils

import (
	"TMS/Comment"
	"strconv"
)



var RoleName=map[string]int{
	"Tester":Comment.TESTER,
	"TestLeader":Comment.TESTLEADER,
	"TestTeamer":Comment.TESTTEAMER,
	"SoftLeader":Comment.SOFTWARELEADER,
	"TestManager":Comment.TESTMANAGER,
}

var teamName=map[string]int{
	"MTK_GROUP":0,
	"RTK_GROUP":1,
	"MSTAR_GROUP":2,
	"Certificaty_GROUP":3,
	"Config_GROUP":4,
	"Multiple_GROUP_1":5,
	"Multiple_GROUP_2":6,
	"AUTO_GROUP":7,
}

type PriorityConstructor struct{

	MTK_GROUP 			int		`项目一组 `
	RTK_GROUP 			int		`项目二组`
	MSTAR_GROUP 		int		`项目三组`
	Certificaty_GROUP 	int		`认证组`
	Config_GROUP 		int		`配置集成组`
	Multiple_GROUP_1 	int		`惠州一组`
	Multiple_GROUP_2 	int		`惠州二组`
	AUTO_GROUP 			int		`自动化组`
}
func(this *PriorityConstructor)GenCode()(code int){

	if this.MTK_GROUP!=0{
		code+=this.MTK_GROUP
	}
	if this.RTK_GROUP!=0{
		code+=this.RTK_GROUP*10
	}

	if this.MSTAR_GROUP!=0{
		code+=this.MSTAR_GROUP*100
	}

	if this.Certificaty_GROUP!=0{
		code+=this.MSTAR_GROUP*1000
	}
	if this.Config_GROUP!=0{
		code+=this.Config_GROUP*10000
	}
	if this.Multiple_GROUP_1!=0{
		code+=this.Multiple_GROUP_1*100000
	}
	if this.Multiple_GROUP_2!=0{
		code+=this.Multiple_GROUP_2*1000000
	}

	if this.AUTO_GROUP!=0{
		code+=this.AUTO_GROUP*1000000
	}
	return
}
func(this *PriorityConstructor)ParseCode(code int){
		Code:=strconv.Itoa(code)
		for x,y:=range Code{
			v,_:=strconv.Atoi(string(y))
			switch x {
				case 0:
					this.MTK_GROUP=v
				case 1:
					this.RTK_GROUP=v
				case 2:
					this.MSTAR_GROUP=v
				case 3:
					this.Certificaty_GROUP=v
				case 4:
					this.Config_GROUP=v
				case 5:
					this.Multiple_GROUP_1=v
				case 6:
					this.Multiple_GROUP_2=v
				case 7:
					this.AUTO_GROUP=v

			}
		}
}