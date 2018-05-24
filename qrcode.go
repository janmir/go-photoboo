package photoboo

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func makeQR(filename, data string) error {
	//make qr
	jpg, err := qrcode.Encode(data, qrcode.Medium, 256)
	catch(err)

	//make the image
	outFile, err := os.Create(filename)
	catch(err)

	defer outFile.Close()
	logger("Saving image to: " + filename)

	canvas, _, err := image.Decode(bytes.NewReader(jpg))
	catch(err)

	return jpeg.Encode(outFile, canvas, &jpeg.Options{Quality: 100})
}
