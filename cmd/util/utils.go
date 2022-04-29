package util

import (
	"os"
)

func CreateFolder(name string) (bool, error) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
