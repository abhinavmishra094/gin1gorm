package routes

import (
	"gin1gorm/models"
	"gin1gorm/util"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func GetDir() func(c *gin.Context) {
	return func(c *gin.Context) {
		var path = os.Getenv("TV_DB")
		var directory []models.Directory
		ch := make(chan string)
		go util.Tv_db_login(ch)
		wg.Add(1)
		log.Println("path: ", path)
		bearerToken := "Bearer " + <-ch
		files, err := ioutil.ReadDir(path)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			for i, file := range files {
				log.Println(file.Name())
				ch := make(chan []models.File)
				ch1 := make(chan string)
				wg.Add(1)
				go util.GetSeries(file.Name(), bearerToken, ch1)
				go util.GetFiles(wg, path+file.Name(), ch)

				d := models.Directory{ID: i, Name: file.Name(), Path: path + file.Name(), Files: <-ch}
				directory = append(directory, d)
			}
			wg.Wait()
			c.JSON(200, directory)
		}

	}
}
