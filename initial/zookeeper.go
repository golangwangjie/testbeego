package initial

import (
	"fmt"

	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var (
	ZKConn *zk.Conn
)

func InitZookeeper() {
	var hosts = []string{"192.168.56.1:2181", "192.168.56.1:2182", "192.168.56.1:2183"}
	var err error
	ZKConn, _, err = zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("zk connect ...", ZKConn.Server())
	// defer zkConn.Close()
	return
}
