package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
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

	caCert, err := ioutil.ReadFile("/var/tls-certs/internal-cert-root-ca/tls.crt")
	if err != nil {
		log.Fatal(err)
	}

	extCACert, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	caCertPool.AppendCertsFromPEM(extCACert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	ticker := time.NewTicker(3000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			{

				inStart := time.Now().UnixNano()
				_, err := client.Get("https://echo.blend/status")
				if err != nil {
					log.Fatal(err)
				}
				inEnd := time.Now().UnixNano()

				outStart := time.Now().UnixNano()
				_, err = client.Get("https://echo.test.k8s.centrio.com/status")
				if err != nil {
					log.Fatal(err)
				}
				outEnd := time.Now().UnixNano()

				log.Infof("\nhttps://echo.test.k8s.centrio.com/status: %d ms \nhttps://echo.blend/status: %d ms\n", (outEnd-outStart)/1000, (inEnd-inStart)/1000)

			}
		}
	}
}
