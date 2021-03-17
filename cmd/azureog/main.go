package main

import (
	"flag"
	"log"

	"github.com/bketelsen/azureog"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var (
		fontSize          = flag.Float64("fontSize", 140, "font fontSize in points")
		fontPath          = flag.String("fontPath", "assets/FiraSans-Light.ttf", "filename of the ttf font")
		backgroundImgPath = flag.String("bgImg", "assets/og-standard.png", "image to use as background")
		title             = flag.String("title", "Nothing To See Here, Move Along", "text to print on the image")
		outputPath        = flag.String("output", "og-image.png", "output path for the resulting image")
	)
	flag.Parse()
	img, err := azureog.TextOnImg(
		azureog.Request{
			BgImgPath: *backgroundImgPath,
			FontPath:  *fontPath,
			FontSize:  *fontSize,
			Text:      *title,
		},
	)
	if err != nil {
		return err
	}

	if err := azureog.Save(img, *outputPath); err != nil {
		return err
	}

	log.Println("image saved on [", *outputPath, "]")
	return nil
}
