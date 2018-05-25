package photoboo

import (
	"bufio"
	"io/ioutil"
	"strings"
)

const (
	_domain         = "http://192.168.0.10"
	_imagePath      = "/DCIM/100OLYMP"
	_imageList      = "/get_imglist.cgi"
	_imageThumbnail = "/get_thumbnail.cgi"
)

//Image defines an image
type Image struct {
	directory, filename, size, attribute, data, time string
}

func imageList() []Image {
	res, body, errors := request.Get(_domain + _imageList).
		Query("DIR=" + _imagePath).
		End()
	catchHTTPError("", res, errors)

	//directory | filename | size | attribute | date | time
	scanner := bufio.NewScanner(strings.NewReader(body))

	//new list
	images := make([]Image, 0)

	for scanner.Scan() {
		txt := scanner.Text()
		vals := strings.Split(txt, ",")

		if len(vals) == 6 {
			images = append(images, Image{vals[0], vals[1], vals[2], vals[3], vals[4], vals[5]})
		}
	}

	return images
}

func imageThumbnail(filename string) {
	res, body, errors := request.Get(_domain + _imageThumbnail).
		Query("DIR=" + _imagePath + "/" + filename).
		End()
	catchHTTPError("", res, errors)

	ioutil.WriteFile("something.jpg", []byte(body), 0666)
}

//Connection: Keep-Alive
//User-Agent: OI.Share v2
//Get Image Thumbnail
// > GET /get_thumbnail.cgi?DIR=/DCIM/100OLYMP/P5150144.ORF HTTP/1.1
//Get Image List
// > GET /switch_cammode.cgi?mode=play HTTP/1.1 (optional)
// > GET /get_imglist.cgi?DIR=/DCIM/100OLYMP HTTP/1.1
//Get Connection Mode (used in checking connection)
// > GET /get_connectmode.cgi HTTP/1.1
//Take Photo
// > GET /switch_cammode.cgi?mode=shutter HTTP/1.1
// > GET /exec_takemotion.cgi?com=starttake HTTP/1.1
// > get last taken image
// > GET /exec_takemisc.cgi?com=getrecview HTTP/1.1
// > ...
//Live view
// > GET /switch_cammode.cgi?mode=play HTTP/1.1
// > GET /switch_cammode.cgi?mode=rec&lvqty=0640x0480 HTTP/1.1
// > GET /exec_takemisc.cgi?com=startliveview&port=28488 HTTP/1.1
// > LINK: https://stackoverflow.com/questions/27176523/udp-in-golang-listen-not-a-blocking-call?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
// > GET /exec_takemisc.cgi?com=stopliveview HTTP/1.1
//Set Auto Focus Point
// > GET /exec_takemotion.cgi?com=assignafframe&point=0331x0120 HTTP/1.1
//Set Property
// > POST /set_camprop.cgi?prop=set&propname=shutspeedvalue HTTP/1.0
// > <xml><set><value>15</value></set>
//Turn Off
// > GET /exec_pwoff.cgi HTTP/1.1
