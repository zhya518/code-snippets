package main

// https://cloud.tencent.com/developer/section/1141539

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Vector类型具有未导出的字段，包无法访问。
// 因此，我们编写了一个BinaryMarshal/BinaryUnmarshal方法对来允许我们
// 使用gob包发送和接收类型。 这些接口是
// 在“encoding”包中定义。
// 我们可以等效地使用本地定义的GobEncode/GobDecoder
// 接口。
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	// 一个简单的编码：纯文本。
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary修改接收器，因此必须使用指针接收器。
func (v *Vector) UnmarshalBinary(data []byte) error {
	// 一个简单的编码：纯文本。
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// 此示例传输实现自定义编码和解码方法的值。
func main() {
	var network bytes.Buffer // 替代（Stand-in）网络。

	// 创建编码器并发送值。
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// 创建解码器并接收值。
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)

}
