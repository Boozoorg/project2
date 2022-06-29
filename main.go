package main

import (
	"log"
	"os"
	"time"
)

func main() {
	Minute := 10

	file, err := os.Create("text/log.txt")
	if err != nil {
		log.Print(err)
		return
	}
	
	go Tack(file, Minute)
	go Tick(file, Minute)

	defer func() {
		cerr := file.Close()
		if cerr != nil {
			log.Print(cerr)
		}
	}()

	time.Sleep(time.Second * time.Duration(Minute))

	if err != nil {
		log.Print(err)
		return
	}
}

func Tick(file *os.File, Minute int) {
	for i := 0; i <= Minute; i++ {
		file.Write([]byte("tick\n"))
		<-time.After(time.Second)
	}
}

func Tack(file *os.File, Minute int) {
	for i := 0; i <= Minute; i++ {
		if i%5 == 0 {
			<-time.After(time.Second * 5)
			file.Write([]byte("tack\n"))
		}
	}
}
