// go env - check GOPATH
// go get github.com/jessevdk/go-flags
// go run flags.go -s true --name hoop

package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name"     default:"World" description:"A name to say hello to."`
	Spanish bool   `short:"s" long:"spanish"                  description:"Use Spanish language."`
}

func main() {
	flags.Parse(&opts)

	if opts.Spanish == true {
		fmt.Printf("Hoola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}