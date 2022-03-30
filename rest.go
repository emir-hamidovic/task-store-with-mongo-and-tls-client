package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	certPem := flag.String("cert", "cert.pem", "certificate file for checking server certificates")
	clientCertPem := flag.String("clientcert", "clientcert.pem", "certificate file for client")
	clientKeyPem := flag.String("clientkey", "clientkey.pem", "key file for client")
	flag.Parse()

	clientCert, err := tls.LoadX509KeyPair(*clientCertPem, *clientKeyPem)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := os.ReadFile(*certPem)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatalf("unable to parse cert from %s", *certPem)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

	r, err := client.Get("https://localhost:3333/task")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	html, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", r.Status)
	fmt.Println(string(html))
}
