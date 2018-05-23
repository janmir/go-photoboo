package photoboo

import (
	"fmt"
	"testing"
	"time"

	"github.com/parnurzeal/gorequest"
)

var (
//wg = sync.WaitGroup
)

func TestFrame(t *testing.T) {
	fmt.Println("Frame Test")

	//do the image conversion
	frame("picture.jpg", "frame.png", "out.jpg")
}

func TestDropbox(t *testing.T) {
	logger("Dropbox-Folder Test")

	//do the image conversion
	logger(upload(getCurrDir() + "out.jpg"))
}

func TestServer(t *testing.T) {

	go startServer()
	time.Sleep(time.Second * 10)

	request := gorequest.New()
	res, body, errs := request.Post("http://0.0.0.0:8080/ping").End()
	catchHTTPError("", res, errs)

	logger(body)
}

func TestStartServer(t *testing.T) {
	startServer()
}
