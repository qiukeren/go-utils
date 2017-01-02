package utils

import (
	"github.com/endeveit/enca"
)

func EncodingTest(content *[]byte) (encoding string, err error) {
	encoding = ""
	defer func() {
		if err := recover(); err != nil {

		}
	}()

	analyzer, err := enca.New("zh")
	if err == nil {
		encoding, err = analyzer.FromBytes(*content, enca.NAME_STYLE_ICONV)
		defer analyzer.Free()
	}
	return
}
