package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
)

func main() {
	f, err := exec.LookPath("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	// 用户所在的所有的组的id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)

	c := make(chan os.Signal, 0)
	signal.Notify(c)
	// Block until a signal is received.
	s1 := <-c
	fmt.Println("Got signal:", s1)
}