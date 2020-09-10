package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("https://rs.aspsp.ob.forgerock.financial:443")
	log.Printf("forwarding to -> %s%s\n", target.Scheme, target.Host)

	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// https://stackoverflow.com/questions/38016477/reverse-proxy-does-not-work
		// https://blog.semanticart.com/2013/11/11/a-proper-api-proxy-written-in-go/
		// https://forum.golangbridge.org/t/explain-how-reverse-proxy-work/6492/7
		// https://stackoverflow.com/questions/34745654/golang-reverseproxy-with-apache2-sni-hostname-error

		// log.Println("req.Host=", req.Host)
		// log.Println("req.URL.Host=", req.URL.Host)
		req.Host = req.URL.Host

		proxy.ServeHTTP(w, req)
	})

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}
