package main

import (
	"os"
	"log"
	"strings"
	"io/ioutil"
	"path/filepath"
	"github.com/billyct/diff-csv"
)

func main() {
	log.Println("正在读取文件...")
	files, err := ioutil.ReadDir(basePath())
	checkError(err)

	var csvFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
			csvFiles = append(csvFiles, filepath.Join(basePath(), file.Name()))
		}
	}

	log.Println("正在比对文件...")
	err = diff_csv.DiffCsv(csvFiles[0], csvFiles[1], filepath.Join(basePath(), "output.csv"))
	checkError(err)
}

// https://blog.csdn.net/skh2015java/article/details/78515002
func basePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	checkError(err)
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

// simple panic error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
