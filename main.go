package main

import (
	"net/http"
	"time"

	"github.com/blend/go-sdk/logger"
)

func main() {
	log := logger.All(logger.OptPath("echo"))
	ticker := time.NewTicker(3000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			{

				inStart := time.Now().UnixNano()
				_, err := http.Get("https://echo.blend/status")
				if err != nil {
					log.Fatal(err)
				}
				inEnd := time.Now().UnixNano()

				outStart := time.Now().UnixNano()
				_, err = http.Get("https://echo.test.k8s.centrio.com/status")
				if err != nil {
					log.Fatal(err)
				}
				outEnd := time.Now().UnixNano()

				log.Infof("\nhttps://echo.test.k8s.centrio.com/status: %d ms \nhttps://echo.blend/status: %d ms\n", (outEnd-outStart)/1000, (inEnd-inStart)/1000)

			}
		}
	}
}
