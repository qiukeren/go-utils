package encoding

import (
	"gopkg.in/iconv.v1"
)

func ConvString(from, to, content string) (string, error) {
	cd, err := iconv.Open(to, from)
	if err != nil {
		return "", err
	}
	defer cd.Close()

	gbk := cd.ConvString(content)
	return gbk, nil
}
