package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	//"io"
)

func main(){
	grep()	
}

func grep(){
	target := os.Args[1]
	for _, fname := range os.Args[2:]{
		file, err := os.Open(fname)
		
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		
		for scanner.Scan(){
			if strings.Contains(scanner.Text(), target){
				fmt.Println(scanner.Text())
			}
		}
		file.Close()
	}
}

