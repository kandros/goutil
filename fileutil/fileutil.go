package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// SafeWriteFile write content to file unless it already exists, it exists in this case
func SafeWriteFile(filename string, fileContent []byte) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, fileContent, 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("file %s already exists\n", filename)
		return err
	}
	return nil
}
