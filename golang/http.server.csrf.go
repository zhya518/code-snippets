package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

func main() {
	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/", IndexHandler)

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))

	http.ListenAndServe(":8000", CSRF(muxServer))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 获取token值
	token := csrf.Token(r)
	// 将token写入到header中
	w.Header().Set("X-CSRF-Token", token)
	fmt.Fprintln(w, "hello world.Go")
}
