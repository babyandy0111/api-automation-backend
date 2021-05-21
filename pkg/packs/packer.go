package packs

import (
	"github.com/gobuffalo/packr/v2"
)

/**
 * 直接把 config 整個目錄打包
 */
func InitConfig() *packr.Box {
	boxname := "config"
	boxpath := "../../config"
	return packr.New(boxname, boxpath)
}

func InitTestFile() *packr.Box {
	boxname := "test"
	boxpath := "../../test"
	return packr.New(boxname, boxpath)
}

func PackTestCovertFile() ([]byte, error) {
	testFilePath := "coverage.html"
	content, err := InitTestFile().Find(testFilePath)
	return content, err
}
