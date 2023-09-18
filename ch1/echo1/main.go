// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

func main() {
	start := time.Now()

	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}

	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %v\n", elapsed)
}

//!-
