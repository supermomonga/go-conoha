package identity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Version struct {
	ID    string `json:"id"`
	Links []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
		Type string `json:"type"`
	} `json:"links"`
	MediaTypes []struct {
		Base string `json:"base"`
	} `json:"media-types"`
	Status  string    `json:"status"`
	Updated time.Time `json:"updated"`
}

type Versions []Version

func (v *Versions) UnmarshalJSON(b []byte) (err error) {
	tmp := struct {
		Versions struct {
			Values []Version
		}
	}{}
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		log.Fatal(err)
	}
	vs := Versions(tmp.Versions.Values)
	fmt.Println(vs)
	v = &vs
	return
}

func GetVersions() Versions {
	url := "https://identity.tyo1.conoha.io/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusMultipleChoices {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var versions Versions
	err = json.Unmarshal(body, &versions)
	if err != nil {
		log.Fatal(err)
	}

	return versions

}
