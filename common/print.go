package common

import (
	"github.com/MouseSun/goprint"
)

//P is a quick function to print struct.
func P(title string, c interface{}) {
	goprint.P(title, c)
}
