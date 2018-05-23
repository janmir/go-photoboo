package photoboo

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	dropbox "github.com/tj/go-dropbox"
)

const (
	_debug = true
)

var (
	server *gin.Engine
)

func init() {
	logger("init")

	//Create a new dropbox object.
	box = dropbox.New(dropbox.NewConfig(_token))

	//Init server
	server = gin.Default()

	//debugging
	if !_debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func startServer() {
	//Interface to Olympus
	server.Static("/", "./www")

	//For testing
	server.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//listen and serve on 0.0.0.0:8080
	server.Run(":80")
}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func catchHTTPError(expectedPath string, res gorequest.Response, err []error) error {
	//Standard errors
	if len(err) > 0 {
		for _, e := range err {
			catch(e)
		}
	}

	//http request errors
	if res.StatusCode != http.StatusOK {
		return errors.New("status code:" + res.Status)
	}
	path := res.Request.URL.Path
	if expectedPath != "" && path != expectedPath {
		return errors.New("request unsuccessful:" + path)
	}

	return nil
}

func getCurrDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if _debug {
		return "C:\\Users\\jp\\go\\src\\github.com\\janmir\\go-photoboo\\"
	}
	return dir
}

func logger(arg string) {
	if _debug {
		log.Println(arg)
	}
}
