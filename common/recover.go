package common

func Recover(fn func(interface{})) {
	if err := recover(); err != nil {
		fn(err)
	}
}
