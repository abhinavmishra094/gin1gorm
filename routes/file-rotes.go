package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin1gorm/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

func GetDir() func(c *gin.Context) {
	return func(c *gin.Context) {
		var directory []models.Directory
		chanel := make(chan string)
		wg.Add(1)
		go tv_db_login(chanel)
		bearerToken := "Bearer " + <-chanel
		log.Println("login", bearerToken)
		files, err := ioutil.ReadDir("/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			for i, file := range files {
				log.Println(file.Name())
				ch := make(chan []models.File)
				ch1 := make(chan string)
				wg.Add(1)
				go getSeries(file.Name(), bearerToken, ch1)
				go getFiles("/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows/"+file.Name(), ch)

				d := models.Directory{ID: i, Name: file.Name(), Path: "/run/user/1000/gvfs/smb-share:server=synology1.local,share=abhinavssynologyvol1/TV Shows/" + file.Name(), Files: <-ch}
				directory = append(directory, d)
			}
			wg.Wait()
			c.JSON(200, directory)
		}

	}
}

func tv_db_login(channel chan<- string) {
	values := map[string]string{"apikey": "9782b4fb-c262-4143-b983-acfc19c679e8",
		"pin": "7HHCJSWV"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://api4.thetvdb.com/v4/login", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var res models.ThetvdbLogin
	r := res.UnMarshal(string(body))
	channel <- r.Data.Token
	close(channel)
}

func getSeries(name string, bearerToken string, ch1 chan<- string) {
	url := fmt.Sprintf("https://api4.thetvdb.com/v4/search?q=%s", strings.Replace(name, " ", "%20", -1))
	log.Printf("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Add("Authorization", bearerToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	ch1 <- string(body)
	close(ch1)

}
func getFiles(path string, ch chan<- []models.File) {
	files, err := ioutil.ReadDir(path)
	var Files []models.File
	defer wg.Done()
	if err != nil {
		log.Println(err.Error())
	} else {
		for i, file := range files {
			if !file.IsDir() {
				f := models.File{ID: i, Name: file.Name(), Path: path + "/" + file.Name()}
				Files = append(Files, f)
			}
		}
		ch <- Files
	}
	close(ch)
}
