package nominetuk

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func getFrameStringByFrameName(fileName string) (string, error) {
	xml, err := ioutil.ReadFile(getFramePath(fileName))
	return string(xml), err
}

// Get file path by frame name
func getFramePath(fileName string) string {
	path, _ := filepath.Abs(FrameDir)
	fmt.Println(path)
	path = path + "/" + fileName
	return path
}
