package initial

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlDb *gorm.DB
	err     error
)

//初始化数据连接
func InitMysql() {
	fmt.Println("Init mysql database ...")
	//数据库参数获取
	dbUser := beego.AppConfig.String("mysqluser")
	dbPassword := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlurls")
	dbPort := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	//mysql 注册
	MysqlDb, err = gorm.Open("mysql", dsn)
	checkErr(err)
	MysqlDb.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error Init mysql database ...", err)
		// panic(err)
	}
}
