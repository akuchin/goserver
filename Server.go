package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type PrintHandler struct{}

func (h PrintHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	writer := w

	request := fmt.Sprintf("%s %s %s\n", req.Method, req.URL, req.Proto)
	prn(request, writer, os.Stdout)

	for k, v := range req.Header {
		header := fmt.Sprintf("%s : %s\n", k, v)
		prn(header, writer, os.Stdout)
	}
	body, _ := io.ReadAll(req.Body)
	//_, _ = io.Copy(writer, req.Body)
	name, _ := os.Hostname()

	prn(fmt.Sprintf("\nHostname: %s\n", name), writer, os.Stdout)
	prn(fmt.Sprintf("%s %s\n", req.Host, req.RemoteAddr), writer, os.Stdout)
	prn(fmt.Sprintf("\nBody: %s\n", body), writer, os.Stdout)
	prn(fmt.Sprintf("Time: %f\n", time.Since(start).Seconds()), writer, os.Stdout)

}

func prn(s string, writers ...io.Writer) {
	for _, w := range writers {
		_, _ = fmt.Fprintf(w, s)
	}
}

func main() {
	var h PrintHandler
	ln, _ := net.Listen("tcp", ":9001")
	log.Fatal(http.Serve(ln, h))
	//log.Fatal(http.ListenAndServe(":9001", h))
}
