package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func getDirList(context *gin.Context) {

	files, err := ioutil.ReadDir("/usr/bin")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {

        type fullDirList struct {
            Name        string `json:"name`
            Size        int64 `json:"size"`
            Attributes  os.FileMode  `json:"attributes"`
        }
    
        var dirList = []fullDirList{
            {Name: file.Name(), Size: file.Size(), Attributes: file.Mode()},
        }

        fmt.Println(file.Name(), file.Size() / 1000, file.Mode())
        context.IndentedJSON(http.StatusOK, dirList)
    }
}

func main() {
    router := gin.Default()
    router.GET("/", getDirList)
    router.Run(":8080")
}

