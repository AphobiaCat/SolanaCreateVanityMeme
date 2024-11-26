package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func DBG_LOG(v ...interface{}) {

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Failed to get caller information")
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("Failed to get function information")
		return
	}

	path := file
	filename := filepath.Base(path)

	var outputStr string = "[info] file[" + filename + "]\t| func[" + fn.Name() + "]\t| line[" + convertToString(line) + "]\t| log:"

	for _, val := range v {
		outputStr += convertToString(val)
	}

	fmt.Println(outputStr)

}

func DBG_ERR(v ...interface{}) {

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Failed to get caller information")
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		fmt.Println("Failed to get function information")
		return
	}

	path := file
	filename := filepath.Base(path)

	var outputStr string = "[error] file[" + filename + "]\t| func[" + fn.Name() + "]\t| line[" + convertToString(line) + "]\t| log:"

	for _, val := range v {
		outputStr += convertToString(val)
	}

	fmt.Println(outputStr)
}

func convertToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
