package main

import "github.com/mmcloughlin/geohash"

func getGeoHash(lat float64, lon float64) string {
	return geohash.Encode(lat, lon)
}
