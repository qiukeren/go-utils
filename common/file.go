package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Touch(filename string) {
	exec.Command("touch", filename).Run()
}

func ReadToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}
