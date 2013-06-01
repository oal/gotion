package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type configuration struct {
	MaxSize int64
}

var config configuration

func loadConfig() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &config)
}

type file struct {
	Path     string
	Datetime time.Time
	Size     int64
	IsVideo  bool
}

type files []file

var motionFiles files

func (f *files) Size() (size int64) {
	for _, file := range *f {
		size += file.Size
	}
	size /= 1000000 // To MB
	return
}

func main() {
	loadConfig()
	loadFiles()

	fmt.Println(motionFiles.Size())
	fmt.Println("Max size (MB): ", config.MaxSize)
}

// loadFiles loads all images when rascam is started.
func loadFiles() {
	fileList, err := ioutil.ReadDir("motion")
	if err != nil {
		log.Println(err)
	}

	for _, f := range fileList {
		name := f.Name()
		var stringTime string
		if name[len(name)-4:] == ".jpg" {
			stringTime = strings.Split(name, "-")[1]
		} else {
			stringTime = strings.Split(name, ".")[0]
			if stringTime[2] == '-' {
				stringTime = strings.Split(stringTime, "-")[1]
			}
		}

		parsedTime, err := time.Parse("20060102150405", stringTime)
		if err != nil {
			log.Printf("Unable to parse time: %v\n", stringTime)
		}

		isVideo := true
		if name[len(name)-4:] == ".jpg" {
			isVideo = false
		}
		motionFiles = append(motionFiles, file{
			Path:     fmt.Sprintf("motion/%v", f.Name()),
			Datetime: parsedTime,
			Size:     f.Size(),
			IsVideo:  isVideo,
		})
	}

	currentSize := motionFiles.Size()
	if currentSize > config.MaxSize {
		overflow := currentSize - config.MaxSize
		var removed int64

		for _, f := range motionFiles {
			err := os.Remove(f.Path)
			if err != nil {
				panic(err)
			}

			removed += f.Size
			if removed >= overflow {
				break
			}
		}

		log.Println("Removed", removed)

	}
}
