package main

import "log"

func main() {
	log.SetPrefix("lada: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)
	log.Println("Hello, World!")
}
