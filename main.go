package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type file struct {
	Path     string
	Datetime time.Time
	Size     int64
	IsVideo  bool
}

type files []file

func (f *files) Size() (size int) {
	for _, file := range f {
		size += file.Size
	}
	return
}

func main() {
	loadFiles()
	fmt.Println(files)
	fmt.Println(files.Size())
}

// loadFiles loads all images when rascam is started.
func loadFiles() {
	fileList, err := ioutil.ReadDir("./motion")
	if err != nil {
		log.Println(err)
	}

	for _, f := range fileList {
		name := f.Name()
		stringTime := strings.Split(name, "-")[1]
		parsedTime, err := time.Parse("20060102150405", stringTime)
		if err != nil {
			log.Println("Unable to parse time")
		}

		isVideo := true
		if name[len(name)-3:] == ".jpg" {
			isVideo = false
		}
		files = append(files, file{
			Path:     fmt.Sprintf("/motion/%v", f.Name()),
			Datetime: parsedTime,
			Size:     f.Size(),
		})
	}
}
