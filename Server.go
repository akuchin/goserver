package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type PrintHandler struct{}

func (h PrintHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	writer := w
	fmt.Fprintf(writer, "%s %s %s\n", req.Method, req.URL.Path, req.Proto)

	for k, v := range req.Header {
		fmt.Fprintf(writer, "%s : %s\n", k, v)
	}

	io.Copy(writer, req.Body)
	name, _ := os.Hostname()
	fmt.Fprintf(writer, "\nHostname: %s\n", name)

	fmt.Fprintf(writer, "%s %s\n", req.Host, req.RemoteAddr)
	fmt.Fprintf(writer, "Time: %f\n", time.Since(start).Seconds())
}

func main() {
	var h PrintHandler
	log.Fatal(http.ListenAndServe(":9001", h))
}
