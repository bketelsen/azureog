package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bketelsen/azureog"
)

var assets azureog.Assets

func init() {
	assets = azureog.Assets{
		BgImgPath: "assets/og-standard.png",
		FontPath:  "assets/FiraSans-Light.ttf",
		FontSize:  60,
	}
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/ogimage", assets.Serve)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
