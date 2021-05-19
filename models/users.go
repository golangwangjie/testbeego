package models

import (
	"fmt"
	"golangwangjie/initial"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // string默认长度为255, 使用这种tag重设。
}

func (User) TableName() string {
	return "users"
}

func NewUser() *User {
	return &User{}
}

func IndexUsers() {
	ifhas := initial.MysqlDb.HasTable(&User{})
	if ifhas {
		return
	}
	fmt.Println("mysql create table users ...")
	initial.MysqlDb.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}

func AddUserInfo() {
	user := User{Name: "wangjie1", Age: 18, Birthday: time.Now()}
	prikey := initial.MysqlDb.NewRecord(user) // => 主键为空返回`true`
	fmt.Println("................... = ", prikey)
	initial.MysqlDb.Create(&user)
	prikey1 := initial.MysqlDb.NewRecord(user) // => 创建`user`后返回`false`
	fmt.Println("................... = ", prikey1)
}

func GetUserInfo(id int) *User {
	list := NewUser()
	initial.MysqlDb.First(list)
	return list
}

func UpdateUserInfo(name string, age int) {
	list := NewUser()
	initial.MysqlDb.Model(list).Where("name = ?", name).Update("age", age)
}

func DelUserInfo(name string) {
	list := NewUser()
	initial.MysqlDb.Where("name = ?", name).Delete(list)
}
