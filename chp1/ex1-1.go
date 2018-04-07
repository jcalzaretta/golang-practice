// modify the echo program to print Arg[0] as well
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
