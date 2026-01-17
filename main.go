package main

import (
	"fmt"

	"github.com/fivemeepo/mainrepo/lib"
	"github.com/fivemeepo/mainrepo/subrepo"
)

func main() {
	lib.Hello()
	fmt.Println("main repo version: v1.1")
	subrepo.Hello()
}
