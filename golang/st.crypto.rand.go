import (
	"encoding/base64"
	"crypto/rand"
	"io"
)

// RandInt64 ...
func RandInt64(x int) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(x)))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// sessionId函数用来生成一个session ID，即session的唯一标识符
func sessionId() string {
	b := make([]byte, 32)
	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//将生成的随机数b编码后返回字符串,该值则作为session ID
	// url.QueryEscape(sessionId) //对sessionId进行转码使之可以安全的用在URL查询里
	// url.QueryUnescape(encodedSessionId) //将QueryEscape转码的字符串还原
	return base64.URLEncoding.EncodeToString(b)
}
