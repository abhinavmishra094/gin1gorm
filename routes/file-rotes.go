package routes

import (
	"gin1gorm/models"
	"io/ioutil"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func GetDir() func(c *gin.Context) {
	return func(c *gin.Context) {
		var directory []models.Directory

		files, err := ioutil.ReadDir("/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			for i, file := range files {
				log.Println(file.Name())
				ch := make(chan []models.File)
				wg.Add(1)
				go getFiles("/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows/"+file.Name(), ch)
				d := models.Directory{ID: i, Name: file.Name(), Path: "/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows/" + file.Name(), Files: <-ch}
				directory = append(directory, d)
			}
			wg.Wait()
			c.JSON(200, directory)
		}

	}
}

func getFiles(path string, ch chan<- []models.File) {
	files, err := ioutil.ReadDir(path)
	var Files []models.File
	defer wg.Done()
	if err != nil {
		log.Println(err.Error())
	} else {
		for i, file := range files {
			f := models.File{ID: i, Name: file.Name(), Path: path + "/" + file.Name()}
			Files = append(Files, f)
		}
		ch <- Files
	}
	close(ch)
}
