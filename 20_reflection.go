package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type response struct {
	Item   string `json: "item"`
	Album  string
	Title  string
	Artist string
}

type respWrapper struct {
	response
}

// this works with normal unmarshalling since it contains the same field types as the response struct
// var j1 = `{
// 	"item": "album",
// 	"alum": "Dark side of the moon"
// }`

// this doesn't work since album and song are not fields in the response struct
var j1 = `{
	"item": "album",
	"album": {"title": "Dark side of the moon"}
}`

var j2 = `{
	"item": "song",
	"song": {"title": "Bella Donna", "artist": "Stevie Nicks"}
}`

func main() {
	var resp1 respWrapper
	var err error

	if err = resp1.UnmarshalJSON(j2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp1.response) // "%#v" formats the output in go syntax
}

// customer unmarshaller
func (r *respWrapper) UnmarshalJSON(input string) error {
	var raw map[string]interface{}
	var resp response

	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return err
	}

	fmt.Printf("%#v\n", raw)


	item, ok := raw["item"].(string)
	if !ok {
		return errors.ErrUnsupported
	}
	
	resp = response {
		Item: item,  
   }

	switch raw["item"] {
	case "album": 
	    fmt.Println("album type")
		album, ok := raw["album"].(map[string]interface{})
		if !ok {
			return errors.ErrUnsupported
		}

		title, ok := album["title"].(string)
		if !ok {
			return errors.ErrUnsupported
		}

		resp.Title = title

	case "song":
		fmt.Println("song type")
		
		song, ok := raw["song"].(map[string]interface{})
		if !ok {
			return errors.ErrUnsupported
		}

		title, ok := song["title"].(string)
		if !ok {
			return errors.ErrUnsupported
		}

		artist, ok := song["artist"].(string)
		if !ok {
			return errors.ErrUnsupported
		}

		resp.Title = title
		resp.Artist = artist		

     default: 
		return errors.ErrUnsupported
	}

	r.response = resp

	return nil
}

// type assertion is use to convert from interface to concrete type
