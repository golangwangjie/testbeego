package logs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

var Loginfo *logs.BeeLogger

func init() {
	BConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("config init error:", err)
		return
	}

	logConf := make(map[string]interface{})
	filestr := strings.Split(BConfig.String("log::log_path"), ".")
	confile := filestr[0] + "_" + time.Now().Format("2006-01-02") + "." + filestr[1]
	logConf["filename"] = confile
	level, _ := BConfig.Int("log::log_level")
	logConf["level"] = level

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}

	Loginfo = logs.NewLogger(10000)                      // 创建一个日志记录器，参数为缓冲区的大小
	Loginfo.SetLogger(logs.AdapterFile, string(confStr)) // 设置日志记录方式：控制台记录
	Loginfo.SetLevel(logs.LevelDebug)                    // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	Loginfo.EnableFuncCallDepth(true)                    // 输出log时能显示输出文件名和行号（非必须）
	Loginfo.Async()                                      //设置异步输出
}
