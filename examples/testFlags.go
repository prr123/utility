// program to test Flags
// author: prr, azul software
// date:22 May 2023
//
//
package main

import (
	"os"
	"fmt"

	util "github.com/prr123/utility/utilLib"
)

func main() {

	tstStr := []string{"program", "/dbg", "/test=hello"}
	flags:=[]string{"dbg","test"}

	fmt.Printf("*** cli %d ***\n", len(tstStr))
	for i:=0; i< len(tstStr); i++ {
		fmt.Printf("cli flag[%d]: %s\n", i+1, tstStr[i])
	}

	fmt.Printf("*** flags %d ***\n", len(flags))
	for i:=0; i< len(flags); i++ {
		fmt.Printf("flag[%d]: %s\n", i+1, flags[i])
	}

	flagMap, err := util.ParseFlags(tstStr, flags)
	if err != nil {
		fmt.Printf("error util.ParseFlags: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("*** found flags %d ***\n", len(flagMap))
	for k, v :=range flagMap {
		fmt.Printf("k: %s v: %s\n", k, v)
	}

}

