package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/blend/go-sdk/logger"
	"github.com/blend/go-sdk/web"
)

func main() {
	log := logger.All(logger.OptPath("echo"))

	app := web.New(web.OptConfigFromEnv(), web.OptLog(log))
	app.GET("/", func(r *web.Ctx) web.Result {
		return web.Text.Result("echo")
	})


	ticker := time.NewTicker(3000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			{
				_, err := client.Get("https://anderjaska-echo2.blend.svc.cluster.local/status")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
