package deal

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

type GUID struct{}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func (self *GUID) GetGuid() string {
	b := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//return GetMd5String(base64.URLEncoding.EncodeToString(b))
	return base64.URLEncoding.EncodeToString(b)
}

/*func main() {
	str := GetGuid()
	fmt.Println(str)
}*/
