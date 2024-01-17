package main

import (
	"fmt"
	"net/http"
)

// transport实现了RoundTripper接口，该接口只有一个方法RoundTrip()，故transport的入口函数就是RoundTrip()。
// transport的主要功能其实就是缓存了长连接，用于大量http请求场景下的连接复用，减少发送请求时TCP(TLS)连接建立的时间损耗，同时transport还能对连接做一些限制，如连接超时时间，每个host的最大连接数等。transport对长连接的缓存和控制仅限于TCP+(TLS)+HTTP1，不对HTTP2做缓存和限制。
//
// tranport包含如下几个主要概念：
//
// 连接池：在idleConn中保存了不同类型(connectMethodKey)的请求连接(persistConn)。当发生请求时，首先会尝试从连接池中取一条符合其请求类型的连接使用
// readLoop/writeLoop：连接之上的功能，循环处理该类型的请求(发送request，返回response)
// roundTrip：请求的真正入口，接收到一个请求后会交给writeLoop和readLoop处理。
// 一对readLoop/writeLoop只能处理一条连接，如果这条连接上没有更多的请求，则关闭连接，退出循环，释放系统资源

// LoggingRoundTripper This type implements the http.RoundTripper interface
type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

// RoundTrip ...
func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	// Do "before sending requests" actions here.
	fmt.Printf("Sending request to %v\n", req.URL)
	// Send the request, get the response (or the error)
	res, e = lrt.Proxied.RoundTrip(req)
	// Handle the result.
	if e != nil {
		fmt.Printf("Error: %v", e)
	} else {
		fmt.Printf("Received %v response\n", res.Status)
	}
	return
}

func main() {
	httpClient := &http.Client{
		Transport: LoggingRoundTripper{http.DefaultTransport},
		//Transport: LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
	httpClient.Get("https://example.com/")
}
