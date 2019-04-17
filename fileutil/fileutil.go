package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// SafeWriteFileOrExit write content to file unless it already exists, it exists in this case
func SafeWriteFileOrExit(filename string, fileContent []byte) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, fileContent, 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("file %s already exists\n", filename)
		os.Exit(0)
	}
}
