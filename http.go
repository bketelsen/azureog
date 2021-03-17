package azureog

import (
	"bytes"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
	Text      string
}

type Assets struct {
	BgImgPath string
	FontPath  string
	FontSize  float64
}

//Serve writes on the image referenced on Assets.BgImgPath with the font set on Assets.FontPath and
//the size Assets.FontSize the text present on the http.Request.
func (a *Assets) Serve(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("title")
	size := r.URL.Query().Get("size")
	fontSize := 60.0
	var err error
	if size == "" {
		fontSize = a.FontSize
	} else {
		fontSize, err = strconv.ParseFloat(size, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	img, err := TextOnImg(
		Request{
			BgImgPath: a.BgImgPath,
			FontPath:  a.FontPath,
			FontSize:  fontSize,
			Text:      text,
		},
	)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("couldn't write image to response")
	}
}
