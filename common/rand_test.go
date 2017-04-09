package common

import (
	"testing"
)

func TestRandString(t *testing.T) {
	if RandString(13, "i") != "iiiiiiiiiiiii" {
		t.Error("not equal")
	}
	if len(RandString(13, "i")) != 13 {
		t.Error("not equal")
	}
}

func BenchmarkRandString(b *testing.B) {
	RandString(32, "oafuoixgunr0tw908209485t28c29quxiosjgiohao")
}
