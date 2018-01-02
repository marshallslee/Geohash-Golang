package main

import (
	"io/ioutil"
	"log"
)

func getAddressFromFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Failed reading data from file %s", err)
	}

	address := string(data)
	return address
}