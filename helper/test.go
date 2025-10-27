package helper

import (
	"runtime"
)

func GetOsName() string  {
	return runtime.GOOS
}