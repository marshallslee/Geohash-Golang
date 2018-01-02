package main

import "fmt"

func main() {
	filePath := "/home/marshall/GoLandProjects/Geo/address.txt"
	address := getAddressFromFile(filePath)

	lat, lon := getLatAndLonFromAddress(address)

	fmt.Println("Latitude: ", lat)
	fmt.Println("Longitude: ", lon)

	hash := getGeoHash(lat, lon)

	fmt.Println(hash)
}
