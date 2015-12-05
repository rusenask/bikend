package main

import (
	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//[
//	{
//	"geometry" : {"x" : -118.15, "y" : 33.80},
//	"attributes" : {
//					"OWNER" : "Joe Smith",
//					"VALUE" : 94820.37,
//					"APPROVED" : true,
//					"LASTUPDATE" : 1227663551096
//					}
//	},
//	{
//	"geometry" : { "x" : -118.37, "y" : 34.086 },
//	"attributes" : {
//					"OWNER" : "John Doe",
//					"VALUE" : 17325.90,
//					"APPROVED" : false,
//					"LASTUPDATE" : 1227628579430
//					}
//					}
//]

type coordinates struct {
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
	Geometry   coordinates  `json:"geometry"`
	Attributes PlacePayload `json:"attributes"`
}

func (h *HTTPClientHandler) addEsriNode(place HostingPlace) error {
	c := h.http

	parkid := fmt.Sprint("%s%f%f", place.Host, place.Lat, place.Long)
	payload := &PlacePayload{
		Parkid: parkid,
		Lat:    place.Lat,
		Lng:    place.Long,
		Host:   place.Host,
		Spaces: place.Space,
	}

	bts, err := json.Marshal(payload)

	if err != nil {
		log.Error(err.Error())
	}

	req, err := http.NewRequest("POST", AppConfig.ESRIEndpoint, bytes.NewBuffer(bts))
	req.Header.Set("Content-Type", "application/json")
	_, err = c.HTTPClient.Do(req)

	return err
}
