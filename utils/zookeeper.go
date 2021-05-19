package utils

import (
	"fmt"
	"golangwangjie/initial"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// 增
func ZKadd(path string, data []byte) {
	fmt.Println("zk set data ... = ", string(data))
	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	s, err := initial.ZKConn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建失败: %v\n", err)
		return
	}
	fmt.Printf("创建: %s 成功 ", s)
}

// 查
func ZKget(path string) {
	data, _, err := initial.ZKConn.Get(path)
	if err != nil {
		fmt.Printf("查询%s失败, err: %v\n", path, err)
		return
	}
	fmt.Printf("%s 的值为 %s\n", path, string(data))
}

// 删改与增不同在于其函数中的version参数,其中version是用于 CAS支持
// 可以通过此种方式保证原子性
// 改
func ZKmodify(path string, newdata []byte) {
	// new_data := []byte("hello zookeeper")
	_, sate, _ := initial.ZKConn.Get(path)
	_, err := initial.ZKConn.Set(path, newdata, sate.Version)
	if err != nil {
		fmt.Printf("数据修改失败: %v\n", err)
		return
	}
	fmt.Println("数据修改成功")
}

// 删
func ZKdel(path string) {
	_, sate, _ := initial.ZKConn.Get(path)
	err := initial.ZKConn.Delete(path, sate.Version)
	if err != nil {
		fmt.Printf("数据删除失败: %v\n", err)
		return
	}
	fmt.Println("数据删除成功")
}

//分布式锁
func ZKLock(path string) {
	acls := zk.WorldACL(zk.PermAll)
	l := zk.NewLock(initial.ZKConn, path, acls)
	err := l.Lock()
	if err != nil {
		panic(err)
	}
	fmt.Println("lock succ, do your business logic")
	time.Sleep(time.Duration(2) * time.Second)
	// do some thing
	l.Unlock()
	fmt.Println("unlock succ, finish business logic")
}

//发布与订阅（配置中心）

// zk 订阅函数
func ZKsubscribe(e <-chan zk.Event) {
	for {
		select {
		case event := <-e:
			{
				fmt.Println("path:", event.Path)
				fmt.Println("type:", event.Type.String())
				fmt.Println("state:", event.State.String())

				if event.Type == zk.EventNodeCreated {
					fmt.Printf("has node[%s] create\n", event.Path)
					_, _, e, _ = initial.ZKConn.ExistsW("/watchzk")
				} else if event.Type == zk.EventNodeDeleted {
					fmt.Printf("has new node[%s] delete\n", event.Path)
					_, _, e, _ = initial.ZKConn.ExistsW("/watchzk")
				} else if event.Type == zk.EventNodeDataChanged {
					fmt.Printf("has node[%s] data changed", event.Path)
					_, _, e, _ = initial.ZKConn.ExistsW("/watchzk")
				}
			}
		}
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}

// zk 发布函数
func ZKpublish(path string, data []byte) {
	ZKadd(path, data)
	time.Sleep(time.Duration(2) * time.Second)
	ZKdel(path)
}
