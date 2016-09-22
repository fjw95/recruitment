package main

import (
	"fmt"
)

func main() {
	flag.IntVar(&maxGoroutinesSpawn, "max", 1, "max goroutines")
	flag.StringVar(&targetFile, "target", "./output.txt", "target save file")
	flag.Parse()

	if rootUrl == "" {
		fmt.Println("Cannot nil URL Parameter")
		os.Exit(-1)
	} else {
		getSiteURL(rootUrl, maxGoroutinesSpawn, targetFile)
	}
}
