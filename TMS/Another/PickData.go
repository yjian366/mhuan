package Another

import (
	"TMS/Bin"
	"TMS/Role"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetAllTester()([]Component.Tester){

	value,err:=os.Open("another/SQAUsers.txt")

	if err!=nil{
		fmt.Println(err)
		return nil
	}
	abc:=bufio.NewReader(value)
	tmp:=make([]byte,1024)

	users:=[]Component.Tester{}


	for err!=io.EOF{

		tmp,_,err=abc.ReadLine()

		if string(tmp)!=""{
			record:=strings.Split(string(tmp),",")
			user:=Component.Tester{}
			for x,y:=range record{
				switch x {
				case 0:
					user.Department=y
				case 1:
					if y=="MTK_GROUP" {
						user.Team = "项目一组"
						user.Priority=1
					}
					if y=="RTK_GROUP"{
						user.Team="项目二组"
						user.Priority=10
					}
					if y=="MSTAR_GROUP"{
						user.Team="项目三组"
						user.Priority=100
					}

					if y=="CERTIFICATY_GROUP"{
						user.Team="认证组"
						user.Priority=1000
					}

					if y=="CONFIG_GROUP"{
						user.Team="配置集成组"
						user.Priority=10000
					}

					if y=="MULTIPLE1_GROUP"{
						user.Team="惠州一组"
						user.Priority=100000
					}

					if y=="MULTIPLE2_GROUP"{
						user.Team="惠州二组"
						user.Priority=1000000
					}
					if y=="AUTO_GROUP"{
						user.Team="自动化组"
						user.Priority=10000000
					}
					if y=="ALL"{
						user.Team="无"
						user.Priority=11111111
					}
				case 2:
					user.Name=strings.TrimSpace(y)
				case 3:
					user.Email=strings.TrimSpace(y)
				case 4:
					user.Title=strings.TrimSpace(y)
					if user.Title=="Tester"{
						user.Priority=user.Priority*1
					}

					if user.Title=="TestLeader"{
						user.Priority=user.Priority*3
					}
					if user.Title=="TestTeamer"{
						user.Priority=user.Priority*5
					}

					if user.Title=="SoftLeader"{
						user.Priority=user.Priority*7
					}

					if user.Title=="TestManager"{
						user.Priority=user.Priority*9
					}
				}
			}
			users=append(users,user)

		}

	}
	return users

}

func WriteDB(){
	data:=GetAllTester()

	fmt.Println(len(data))
	if data!=nil{

		if len(data)!=0{
			for _,v:=range data{
				fmt.Println(v)
				e:=Bin.UserAssisant.Add(v)
				fmt.Println(e)
			}


		}else{
			return
		}

	}else{
		return
	}
}