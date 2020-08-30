package main

import (
	"pandaschool.net/demo/micro"
	"pandaschool.net/demo/micro/example/user/handler"
)

func main() {
	opt := micro.Options{
		Name:    "user",
		Addr:    "127.0.0.1:9111",
		Handler: &handler.User{},
	}

	engine := micro.NewServer(&opt)
	engine.Run()
}
