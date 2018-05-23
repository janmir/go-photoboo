package photoboo

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	//fmt wouldnt let me
	_ "image/png"
)

func frame(mainImage, frameImage, outputImage string) {
	//get image again
	frame := readImage(frameImage)
	photo := readImage(mainImage)

	canvas := image.NewRGBA(frame.Bounds())
	canvasBounds := canvas.Bounds()
	draw.Draw(canvas, canvasBounds, photo, image.ZP, draw.Src)
	draw.Draw(canvas, canvasBounds, frame, image.ZP, draw.Over)

	outFilename := outputImage
	outFile, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	log.Print("Saving image to: ", outFilename)

	//Resize
	//outImage := resize.Resize(500, 0, canvas, resize.NearestNeighbor)

	jpeg.Encode(outFile, canvas, &jpeg.Options{Quality: 100})
}

func readImage(file string) image.Image {
	imageFile, err := os.Open(file)
	catch(err)
	defer imageFile.Close()

	image, _, err := image.Decode(imageFile)
	catch(err)

	return image
}
