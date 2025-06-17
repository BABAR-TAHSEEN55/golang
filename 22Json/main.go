package main

import (
	"encoding/json"
	"fmt"
)

// NOTE : You can change the key pair by using `json:"anynameyouwant"`
// NOTE: : And using "-" in password hidews it
type course struct {
	Name     string
	Price    string
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Welcome to Handling JSON")
	// EncodeJSON()
	DecodeJSON()
}
func EncodeJSON() {
	NewJSON := []course{
		{
			"Mern-Stack",
			"299",
			"www.youtube.com",
			"abc123",
			[]string{"Web-dev"},
		},

		{
			"React",
			"199",
			"www.youtube.com",
			"bdc1222",
			[]string{"Render"},
		},
	}
	finalJSON, err := json.MarshalIndent(NewJSON, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\t", finalJSON)

}
func DecodeJSON() {
	JsonFromWeb := []byte(`
        {
                "Name": "React",
                "Price": "199",
                "Platform": "www.youtube.com",
                "tags": ["Render"]
        }
	`)

	var StoreJson course
	CheckValid := json.Valid(JsonFromWeb)
	if CheckValid {
		fmt.Println("Json Is Valid")
		json.Unmarshal(JsonFromWeb, &StoreJson)
		fmt.Printf("%#v", StoreJson) // FOr printing this interface you gotta use %#v
	} else {
		fmt.Println("JSON was not Valid")

	}
}
//TODO : Do a Map 
