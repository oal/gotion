package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type image struct {
	Path     string
	Datetime time.Time
}

type video struct {
	Path     string
	Datetime time.Time
}

var images []image
var videos []video

func main() {
	loadFiles()
	fmt.Println(images)
	fmt.Println(videos)
}

// loadFiles loads all images when rascam is started.
func loadFiles() {
	files, err := ioutil.ReadDir("./motion")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		name := file.Name()
		stringTime := strings.Split(name, "-")[1]
		parsedTime, err := time.Parse("20060102150405", stringTime)
		if err != nil {
			log.Println("Unable to parse time")
		}

		if name[len(name)-3:] == ".jpg" {
			images = append(images, image{
				Path:     fmt.Sprintf("/motion/%v", file.Name()),
				Datetime: parsedTime,
			})
		} else {
			videos = append(videos, video{
				Path:     fmt.Sprintf("/motion/%v", file.Name()),
				Datetime: parsedTime,
			})
		}
	}
}
