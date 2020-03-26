// go build -o hello
// go run hello.go
package main

import (
    "io"
	"log"
	"fmt"
	"runtime"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, getOsPlatform())
}

func getOsPlatform() string {
	const os, arch = runtime.GOOS, runtime.GOARCH
	return fmt.Sprintf("{'os': '%s', 'platform': '%s'}", os, arch)
}