package main

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"github.com/tidwall/gjson"
)

func getLatAndLonFromAddress(address string) (lat float64, lon float64){
	var Url *url.URL
	Url, err := url.Parse(apiDefaultUrl)
	if err != nil {
		panic("boom")
	}

	parameters := url.Values{}
	parameters.Add("query",address)
	Url.RawQuery = parameters.Encode()

	geocodeRequest, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		panic(err)
	}
	geocodeRequest.Header.Add("X-Naver-Client-Id", clientID)
	geocodeRequest.Header.Add("X-Naver-Client-Secret", clientSecret)

	client := &http.Client{}
	response, err := client.Do(geocodeRequest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)
	str := string(bytes)

	latitude := gjson.Get(str, "result.items.0.point.y")
	longitude := gjson.Get(str, "result.items.0.point.x")

	return latitude.Float(), longitude.Float()
}
