package main

import (
	"fmt"
	"log"
	"os"
)

func printColor(color string, msg string) {
	colorData := [...][2]string{
    {"reset", "\033[0m"},
		{"black", "\033[30m"},
		{"red", "\033[31m"},
		{"green", "\033[32m"},
		{"yellow", "\033[33m"},
		{"blue", "\033[34m"},
		{"magenta", "\033[35m"},
		{"cyan", "\033[36m"},
		{"white", "\033[37m"},
	}

	for _, colorIter := range colorData {
		if color == colorIter[0] {
			fmt.Print(colorIter[1] + msg)
		}
	}

}

func main() {
	dirname := "./"
	var isFile bool

	fileColor := "cyan"
	dirColor := "magenta"

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			isFile = false
			printColor(dirColor, "DIR   ")
		} else {
			isFile = true
			printColor(fileColor, "FILE  ")
		}

		fileInfo, _ := file.Info()
		fileSize := fileInfo.Size()
		var sizeStr string
		if fileSize >= 1e9 {
			printColor("red", "")
			sizeStr = fmt.Sprintf("%.2f GB", float64(fileSize)/(1e9))
		} else if fileSize >= 1e6 {
			printColor("yellow", "")
			sizeStr = fmt.Sprintf("%.2f MB", float64(fileSize)/(1e6))
		} else if fileSize >= 1e3 {
			printColor("green", "")
			sizeStr = fmt.Sprintf("%.2f KB", float64(fileSize)/(1e3))
		} else {
			printColor("blue", "")
			sizeStr = fmt.Sprintf("%.1d B", fileSize)
		}

		fmt.Printf("%-10s", sizeStr)

		if isFile {
			printColor(fileColor, "")
		} else {
			printColor(dirColor, "")
		}

		fmt.Printf("%s\n", file.Name())
		printColor("reset", "")
	}
}
