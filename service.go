package micro

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"net/http"
	"path"
	"reflect"
	"strings"
)

//Register 向etcd注册服务.
func (m *Service) Register() error {
	//连接etcd.
	client, err := clientv3.New(clientv3.Config{
		Endpoints: m.etcdAddr,
	})
	if err != nil {
		return err
	}
	m.etcd = client

	//创建租约.
	lease := clientv3.NewLease(client)
	_, err = lease.Grant(context.Background(), 60)
	if err != nil {
		return err
	}

	_, err = m.etcd.Put(context.Background(), fmt.Sprintf("/%s/%s", "http", m.name), m.httpAddr)
	if err != nil {
		return err
	}

	return nil
}

//RunHTTP 运行http服务.
func (m *Service) RunHTTP(addr string) *Service {
	m.httpAddr = addr

	go func() {
		if err := http.ListenAndServe(m.httpAddr, m); err != nil {
			panic(err)
		}
	}()

	return m
}

//RunRPC 运行grpc服务.
func (m *Service) RunRPC(addr string) *Service {
	m.rpcAddr = addr

	go func() {
		//todo 启动grpc
	}()

	return m
}

func (m *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//路由处理.
	urlParams := strings.Split(path.Clean(r.RequestURI), "?")
	urlItem := strings.Split(urlParams[0], "")
	if urlItem[0] == "/" {
		urlItem[0] = ""
	}
	if urlItem[len(urlItem)-1] == "/" {
		urlItem[len(urlItem)-1] = ""
	}
	url := strings.Join(urlItem, "")
	arr := strings.Split(url, "/")

	fn := ""
	if url == "" {
		fn = "Index"
	} else {
		for _, v := range arr {
			data := strings.Split(v, "")
			data[0] = strings.ToUpper(data[0])

			fn += strings.Join(data, "")
		}
	}

	//验证方法是否存在.
	fn = "Action" + fn
	_, ok := m.handlerType.MethodByName(fn)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("404"))
		return
	}

	//结构体赋值.
	m.handlerValue.Elem().FieldByName("ctx").Set(reflect.ValueOf(Context{
		ResponseWriter: w,
		Request:        r,
		Host:           r.Host,
		IP:             r.RemoteAddr,
	}))

	//调用方法.
	v := m.handlerValue.MethodByName(fn)
	v.Call(nil)
}
