package utils
//
//import (
//	"crypto/md5"
//	"io"
//	"fmt"
//)
//
//func GetEncryptPassword(id string, email string, password string) string {
//	m := md5.New()
//	str := id + password + email + password;
//	io.WriteString(m, str)
//
//	newPwd := fmt.Sprintf("%x", m.Sum(nil))
//	return newPwd
//}
