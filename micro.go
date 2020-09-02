package micro

import (
	"go.etcd.io/etcd/clientv3"
	"reflect"
)

type Options struct {
	Name    string
	Etcd    []string
	Handler interface{}
}

type Service struct {
	httpAddr     string
	rpcAddr      string
	etcdAddr     []string
	name         string
	handlerType  reflect.Type
	handlerValue reflect.Value
	etcd         *clientv3.Client
}

func NewServer(opt *Options) Engine {
	return &Service{
		name:         opt.Name,
		handlerValue: reflect.ValueOf(opt.Handler),
		handlerType:  reflect.TypeOf(opt.Handler),
		etcdAddr:     opt.Etcd,
	}
}
