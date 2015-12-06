package main

import (
	log "github.com/Sirupsen/logrus"

	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	//	"strconv"
)

//
//{
//"geometry": {
//	"x": 2,
//	"y": 51
//	},
//	"attributes": {
//			"parkid": "karolis@rusenas2.com0 51",
//			"lat": 4,
//			"lng": 51,
//			"host_name": "karolis@rusenas2.com",
//			"spaces": 3
//			}
//}
//

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PlacePayload struct {
	Parkid string  `json:"parkid"` // hostemail+lat+lon
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Host   string  `json:"host_name"` // host(owner)
	Spaces int     `json:"spaces"`
}

type EsriFeatureNode struct {
	Geometry   Coordinates  `json:"geometry"`
	Attributes PlacePayload `json:"attributes"`
}

func (h *HTTPClientHandler) addEsriNode(place HostingPlace) (*http.Response, error) {
	c := h.http

	parkid := fmt.Sprintf("%s%d%d", place.Host, int(place.Lat), int(place.Long))

	payload := PlacePayload{
		Parkid: parkid,
		Lat:    place.Lat,
		Lng:    place.Long,
		Host:   place.Host,
		Spaces: place.Space,
	}

	coords := Coordinates{
		X: place.Lat,
		Y: place.Long,
	}

	finalPayload := EsriFeatureNode{Geometry: coords, Attributes: payload}

	bts, err := json.Marshal(finalPayload)

	if err != nil {
		log.Error(err.Error())
	}

	//	fullurl := fmt.Sprintf("%s%s", AppConfig.ESRIEndpoint, string(bts))

	log.WithFields(log.Fields{
		"body":     string(bts),
		"endpoint": AppConfig.ESRIEndpoint,
		//		"fullurl":  fullurl,
	}).Info("Adding esri node")

	//	req, err := http.NewRequest("POST", fullurl, nil)
	req, err := http.NewRequest("POST", AppConfig.ESRIEndpoint, bytes.NewBuffer(bts))
	req.Header.Set("Content-Type", "application/json")
	//	req.Header.Set("Content-Type", "application/html")
	resp, err := c.HTTPClient.Do(req)

	b := bufio.NewScanner(req.Body)

	bodyStr := b.Text()

	log.WithFields(log.Fields{
		"esriStatus": resp.StatusCode,
		"esriBody":   bodyStr,
	}).Info("Got response from esri")

	return resp, err
}
