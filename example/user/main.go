package main

import (
	"log"
	"micro"
	"user/handler"
)

func main() {
	opt := micro.Options{
		Name:    "user",
		Etcd:    []string{"127.0.0.1:2379"},
		Handler: &handler.User{},
	}

	engine := micro.NewServer(&opt)
	_ = engine.RunHTTP(":9111")
	_ = engine.RunRPC(":9112")
	log.Println("over")
}
