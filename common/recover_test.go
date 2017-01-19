package utils

import (
	"fmt"
	"testing"
)

func fn(err interface{}) {
	fmt.Println("recover", err)
}

func TestRecover(t *testing.T) {
	defer Recover(fn)
	panic("lalala")
}
