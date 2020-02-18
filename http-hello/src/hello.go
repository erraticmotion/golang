// go run hello.go &
// curl localhost:4000
// -- To kill background http server, get the PID
// ps
// kill -9 <PID>
// ps

package main

import (
	"fmt"
	"log"
	"net/http"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, g)
}

func main() {
	err := http.ListenAndServe("localhost:4000", greeting("Hello Go\n"))
	if err != nil {
		log.Fatal(err)
	}
}
