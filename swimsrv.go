package main

import (
    "encoding/json"
    "fmt"
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "os"
)

type Objects struct {
    Pools []Pool `json:"pools"`
}

/*
   {
       "title": "Edmonds Community Centre",
       "place": "Edmonds Community Centre",
       "tags": [
       "Southeast",
       "Burnaby"
       ],
       "region": "Burnaby",
       "latitude": 49.2294908,
       "longitude": -123.0025753,
       "information": "Edmonds Community Centre",
       "telephone": "+1 (604) 297-4838",
       "url": "http://www.romanbaths.co.uk",
       "visited": true,
       "address": "4603 Kingsway, Burnaby, BC V5H, Canada"
   },
*/
type Pool struct {
    Title       string   `json:"title"`
    Place       string   `json:"place"`
    Tags        []string `json:"tags"`
    Region      string   `json:"region"`
    Latitude    float64  `json:"latitude"`
    Longitude   float64  `json:"longitude"`
    Information string   `json:"information"`
    Tel         string   `json:"tel"`
    Url         string   `json:"url"`
    Visited     bool     `json:"visited"`
    Address     string   `json:"address"`
}

func GetLocations(w rest.ResponseWriter, req *rest.Request) {
    w.WriteJson(&spools)
}

var spools Objects

func main() {
    handler := rest.ResourceHandler{}
    handler.SetRoutes(
        &rest.Route{"GET", "/pools", GetLocations},
    )

    file, e := os.Open("burnaby.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    jsonParser := json.NewDecoder(file)
    if err := jsonParser.Decode(&spools); err != nil {
        fmt.Printf("Error parsing file: %v\n", err.Error())
    }

    fmt.Printf("Results: %v\n", spools)

    http.ListenAndServe(":8080", &handler)
}
