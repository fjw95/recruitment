package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	sourcePath = "/source"
	targetPath = "/target"
)

func compare(sourceMapFiles map[string]string, targetMapFIles map[string]string) {

	resultMap := []string{}

	var isExist bool
	for kSource, vSource := range sourceMapFiles {
		isExist = false
		for kTarg, vTarg := range targetMapFIles {
			if kSource == kTarg && vSource == vTarg {
				isExist = true
				break
			}
		}
		if isExist == false {
			message := vSource + "/" + kSource + " New"
			resultMap = append(resultMap, message)
		}
	}

	for kTarg, vTarg := range targetMapFIles {
		isExist = false
		for kSource, vSource := range sourceMapFiles {
			if kTarg == kSource && vTarg == vSource {
				isExist = true
				break
			}
		}
		if isExist == false {
			message := kTarg + "/" + vTarg + " Delete"
			resultMap = append(resultMap, message)
		}
	}

	for _, v := range resultMap {
		fmt.Println(v)
	}

}

func getMapFiles(source string, sourceFolder string, level int, mapFiles map[string]string) map[string]string {

	var sourceFiles []os.FileInfo

	if level > 0 && sourceFolder != "" {
		source, _ = filepath.Abs(source + "/" + sourceFolder)
		sourceFiles, _ = ioutil.ReadDir(source)
	} else {
		source, _ = filepath.Abs("." + source)
		sourceFiles, _ = ioutil.ReadDir(source)
		mapFiles = map[string]string{}
	}

	for _, file := range sourceFiles {
		if file.IsDir() {
			level++
			mapFiles = getMapFiles(source, file.Name(), level, mapFiles)
		} else {
			if level > 0 {
				mapFiles[file.Name()] = sourceFolder
			} else {
				mapFiles[file.Name()] = "Root"
			}

		}
	}
	return mapFiles
}

func main() {

	sourceMapFiles := map[string]string{}
	targetMapFIles := map[string]string{}

	sourceMapFiles = getMapFiles(sourcePath, "", 0, sourceMapFiles)
	targetMapFIles = getMapFiles(targetPath, "", 0, targetMapFIles)

	compare(sourceMapFiles, targetMapFIles)
	/*for k, v := range targetMapFIles {
		fmt.Println(k + v)

	}
	for k, v := range sourceMapFiles {
		fmt.Println(k + v)

	}*/
	/*files, err := ioutil.ReadDir("./recruitment")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}*/
}
