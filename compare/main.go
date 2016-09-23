package main

import (
	"crypto/md5"
	"fmt"
	_ "io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	sourcePath = "/source"
	targetPath = "/target"
)

func compareContentFile(source string, target string) bool {
	dataSource, _ := ioutil.ReadFile(source)
	dataTarget, _ := ioutil.ReadFile(target)

	if md5.Sum(dataSource) != md5.Sum(dataTarget) {
		return true
	}
	return false
}

func compareFile(sourceMapFiles map[string]string, targetMapFIles map[string]string) {
	resultMap := []string{}

	for kSourceFile, vSourceFile := range sourceMapFiles {
		for kTargFile, vTargFile := range targetMapFIles {
			if kSourceFile == kTargFile && vSourceFile == vTargFile {
				if vSourceFile == "Root" {
					source, _ := filepath.Abs("." + sourcePath + "/" + kSourceFile)
					target, _ := filepath.Abs("." + targetPath + "/" + kTargFile)
					if compareContentFile(source, target) {
						message := kSourceFile + " Modified"
						resultMap = append(resultMap, message)
					}
				} else {
					source, _ := filepath.Abs("." + sourcePath + "/" + vSourceFile + "/" + kSourceFile)
					target, _ := filepath.Abs("." + sourcePath + "/" + vSourceFile + "/" + kSourceFile)
					if compareContentFile(source, target) {
						message := kSourceFile + "/" + vSourceFile + " Modified"
						resultMap = append(resultMap, message)
					}
				}
			}
		}
	}
	for _, v := range resultMap {
		fmt.Println(v)
	}
}

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
			message := vTarg + "/" + kTarg + " Delete"
			resultMap = append(resultMap, message)
		}
	}

	for _, v := range resultMap {
		fmt.Println(v)
	}

}

func getMapFiles(root string, source string, sourceFolder string, level int, mapFiles map[string]string) map[string]string {

	var sourceFiles []os.FileInfo

	if level > 0 && sourceFolder != "" && root != "" {
		source, _ = filepath.Abs(source + "/" + sourceFolder)
		sourceFiles, _ = ioutil.ReadDir(source)
	} else {
		source, _ = filepath.Abs("." + source)
		root = source
		sourceFiles, _ = ioutil.ReadDir(source)
		mapFiles = map[string]string{}
	}

	for _, file := range sourceFiles {
		if file.IsDir() {
			level++
			mapFiles = getMapFiles(root, source, file.Name(), level, mapFiles)
		} else {
			if level > 0 {
				valSource, _ := filepath.Rel(root, source)
				mapFiles[file.Name()] = valSource
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

	sourceMapFiles = getMapFiles("", sourcePath, "", 0, sourceMapFiles)
	targetMapFIles = getMapFiles("", targetPath, "", 0, targetMapFIles)

	compare(sourceMapFiles, targetMapFIles)
	compareFile(sourceMapFiles, targetMapFIles)
}
