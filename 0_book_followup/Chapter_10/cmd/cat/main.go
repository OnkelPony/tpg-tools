package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("smazat.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = write(file)
	if err != nil {
		log.Fatal(err)
	}
}

type safeWriter struct {
	w     io.Writer
	Error error
}

func (sw *safeWriter) Write(data []byte) {
	if sw.Error != nil {
		return
	}
	_, err := sw.w.Write(data)
	if err != nil {
		sw.Error = err
	}
}

func write(w io.Writer) error {
	data := []byte{1, 0, 8}
	sw := safeWriter{w: w}
	sw.Write(data)
	sw.Write(data)
	sw.Write(data)
	return sw.Error
}
