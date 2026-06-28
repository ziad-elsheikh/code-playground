package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main(){
	wwc();
}

func wwc(){
	if len(os.Args) < 2{
		fmt.Println("wwc <file>")
	}

	var lc, wc, bc int

	for _, fname := range os.Args[1:]{
		file, err := os.Open(fname)

		if err != nil{
			fmt.Fprintln(os.Stderr, err)
				continue
		}

		scanner := bufio.NewScanner(file)

		for{
			byt := scanner.Bytes()
			
			bc++

			if byt == []byte(10){
				lc++
			}

			isSpace := true
			if isSpace{
				wc++
			}
		}
		fmt.Printf("%7d %7d %7d %s", lc, wc, bc, fname)
		file.Close()
	}
}

func wwc2(){
	if len(os.Args) < 2{
			fmt.Println("wwc <file>")
	}

	var lc, wc, bc int
	
	for _, file := range os.Args[1:]{
			fd, err := os.Open(file)

			if err != nil{
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			scanner := bufio.NewScanner(fd)
			
			for scanner.Scan(){
				s := scanner.Text()
				
				wc += len(strings.Fields(s))
				bc += len(s)
				lc++
			}
			fmt.Printf("%3d %3d %3d %s\n", lc, wc, bc, file)
			fd.Close()
			
	 }
}
