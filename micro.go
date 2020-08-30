package micro

import (
	"reflect"
)

type Options struct {
	Name    string
	Addr    string
	Handler interface{}
}

type Service struct {
	addr         string
	name         string
	handlerType  reflect.Type
	handlerValue reflect.Value
}

func NewServer(opt *Options) Engine {
	return &Service{
		addr:         opt.Addr,
		name:         opt.Name,
		handlerValue: reflect.ValueOf(opt.Handler),
		handlerType:  reflect.TypeOf(opt.Handler),
	}
}
