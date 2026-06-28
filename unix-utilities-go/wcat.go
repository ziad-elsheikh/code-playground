package main

import (
	"fmt"
	"os"
	"io"
)

func main(){
	fmt.Println("method1")
	// cat1()
	cat2()
}


func cat1(){
	for _, fname := range os.Args[1:]{
		file, err := os.ReadFile(fname)
		
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(string(file))
	}
}


func cat2(){ // using OS && IO 
	files := os.Args[1:]
	
		for _, file := range files { // looping on all files
			data, err := os.Open(file) // open and read in once
		
			if err != nil{
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			if _, err := io.Copy(os.Stdout, data); err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			data.Close()
	}
}
