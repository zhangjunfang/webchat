package path

import (
	"fmt"
	"os"
	"runtime"
)

//返回当前工作路径
func WorkPath() (dir string, err error) {
	return os.Getwd()
}

//返回当前文件
func CurrentFilePath() (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(0)
}
func Debug() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
}
