package main

import (
	"encoding/json"
	"fmt"
	//"github.com/imdario/mergo"
)

// Sto0 ...
type Sto0 struct {
	Nick    string `json:"nick,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Country string `json:"country,omitempty"`
}

// Sto1 ...
type Sto1 struct {
	Nick    string `json:"nick"`
	Avatar  string `json:"avatar"`
	Country string `json:"country"`
}

func main() {
	st0 := &Sto0{Nick: "n0", Avatar: "a"}
	st1 := &Sto0{Nick: "n1", Country: "c"}
	b, _ := json.Marshal(st1)
	json.Unmarshal(b, st0)
	fmt.Println(st0) //&{n1 a c}

	st2 := &Sto1{Nick: "n2", Avatar: "a"}
	st3 := &Sto1{Nick: "n3", Country: "c"}
	b, _ = json.Marshal(st3)
	json.Unmarshal(b, st2)
	fmt.Println(st2) //&{n3  c}
}
