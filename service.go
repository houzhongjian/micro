package micro

import (
	"log"
	"net/http"
	"path"
	"reflect"
	"strings"
)

//Register .
func (*Service) Register() {
}

//Run .
func (m *Service) Run() {
	if err := http.ListenAndServe(m.addr, m); err != nil {
		log.Panic(err)
	}
}

//.
func (m *Service) RunTLS() {
	if err := http.ListenAndServe(m.addr, m); err != nil {
		log.Panic(err)
	}
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
	m.handlerValue.Elem().FieldByName("Ctx").Set(reflect.ValueOf(Context{
		ResponseWriter: w,
		Request:        r,
		Host:           r.Host,
		IP:             r.RemoteAddr,
	}))

	//调用方法.
	v := m.handlerValue.MethodByName(fn)
	v.Call(nil)
}
