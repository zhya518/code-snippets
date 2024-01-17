package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func rewindTest() {
	var req http.Request
	req.Header = make(http.Header)
	req.Header.Add("a", "a1")
	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte("abc")))

	for i := 0; i < 10; i++ {
		bs := readRewindRequest(&req)
		fmt.Println(i, string(bs))
	}
}

func readRewindRequest(r *http.Request) []byte {
	var b []byte
	if r.Body != nil {
		b, _ = ioutil.ReadAll(r.Body)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return b
}

// curl -i -X POST --data '{"username":"xyz","password":"xyz"}' http://localhost:8080
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := &bytes.Buffer{}
		//_, err := io.Copy(buf, r.Body)
		tee := io.TeeReader(r.Body, buf)
		body, err := ioutil.ReadAll(tee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("body", string(body))
		bodyCopy := bytes.NewReader(buf.Bytes())
		bodyCopy.Seek(0, 0)
		r.Body = io.NopCloser(bodyCopy)
		//body := readRewindRequest(r)
		payload, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("dump request", string(payload))
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
