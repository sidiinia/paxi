package dlog

import (
	"log"
	"io/ioutil"
)

const DLOG = true

func Printf(format string, v ...interface{}) {
	if !DLOG {
		return
	}
	log.Printf(format, v...)
}

func Println(v ...interface{}) {
	if !DLOG {
		return
	}
	log.Println(v...)
}


func PrintOnFile(filename string){
	d1 := []byte("hello\ngo\n")
	ioutil.WriteFile("/tmp/" + filename, d1, 0644)
}
