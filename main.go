package main

import (
	"flag"
	"fmt"

	"./runtime"
)

func main() {

	svcFlag := flag.String("service", "", "Control the system service.")
	runtimeFlag := flag.Bool("runtime", false, "Execute runtime from command line")
	flag.Parse()
	fmt.Println(*svcFlag)

	if len(*svcFlag) != 0 || *runtimeFlag {
		runtime.Exec(svcFlag)
	}

	fmt.Println("open gui")
}
