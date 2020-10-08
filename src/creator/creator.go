package creator

import (
	"fmt"
	"io/ioutil"
)

func CreateStucture() error {
	filesFromDir, err := ioutil.ReadDir(".")
	if err != nil {
		return fmt.Errorf("error read directories config %v: ", err)
	}
	fmt.Println(filesFromDir)
	return nil
}
