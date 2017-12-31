package entities

import "time"

type UserInfo struct {
    UID        int   `orm:"id,auto-inc,type=INT(10)"` //语义标签
    UserName   string
    DepartName string
    CreateAt   *time.Time `orm:"name=created" json:",omitempty"`
}

func NewUser(username string, departname string) *UserInfo {
	var user UserInfo
	user.UserName = username
	user.DepartName = departname
	t := time.Now()
	user.CreateAt = &t
	return &user
}
