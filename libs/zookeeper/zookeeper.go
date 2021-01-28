package zookeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/go-zookeeper/zk"
)

func Connection() {
	ip := os.Getenv("SERVER")

	c, _, err := zk.Connect([]string{ip}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("test %+v %+v\n", children, stat)
	e := <-ch
	fmt.Printf("%+v\n", e)
}
