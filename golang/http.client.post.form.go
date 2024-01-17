package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
		r.ParseForm()
		fmt.Printf("post form:%v\n", r.PostForm) // 打印form数据
		// 2. 请求类型是application/json时从r.Body读取数据
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read request.Body failed, err:%v\n", err)
			return
		}
		fmt.Printf("body data:%s\n", string(b))
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		w.Write(dump)
	}))
	defer ts.Close()
	// 表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=test&age=18"
	// json
	//contentType := "application/json"
	//data := `{"name":"test","age":18}`
	resp, err := http.Post(ts.URL, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
