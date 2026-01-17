package main

import (
	"fmt"

	"github.com/fivemeepo/mainrepo/lib"
	"github.com/fivemeepo/mainrepo/subrepo"
)

func main() {
	lib.Hello()
	fmt.Println("main repo version: v6")
	fmt.Println(subrepo.Hello())
}
