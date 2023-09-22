package helpers

import (
	"path"
	"runtime"
	"strconv"
)

// CheckErr error panic
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetProjectPath get project root dir
func GetProjectPath() string {

	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}

// StringToInt
func StringToInt(value string) int {
	num, _ := strconv.Atoi(value)
	return num
}
