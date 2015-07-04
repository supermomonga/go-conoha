package identity

import (
	"encoding/json"
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

type Metadata struct {
	Roles   []string `json:"roles"`
	IsAdmin int      `json:"is_admin"`
}

type User struct {
	Name       string        `json:"name"`
	Roles      Roles         `json:"roles"`
	ID         string        `json:"id"`
	RolesLinks []interface{} `json:"roles_links"`
	Username   string        `json:"username"`
}

type Role struct {
	Name string `json:"name"`
}

type Roles []Role

type Service struct {
	Name           string        `json:"name"`
	Type           string        `json:"type"`
	EndpointsLinks []interface{} `json:"endpoints_links"`
	Endpoints      []struct {
		Publicurl string `json:"publicURL"`
		Region    string `json:"region"`
	} `json:"endpoints"`
}

type Services []Service

type Tenant struct {
	Description   string `json:"description"`
	DomainID      string `json:"domain_id"`
	Sjc1ImageSize string `json:"sjc1_image_size"`
	Sin1ImageSize string `json:"sin1_image_size"`
	ID            string `json:"id"`
	Tyo1ImageSize string `json:"tyo1_image_size"`
	Enabled       bool   `json:"enabled"`
	Name          string `json:"name"`
}

type Token struct {
	AuditIds []string  `json:"audit_ids"`
	Tenant   Tenant    `json:"tenant"`
	ID       string    `json:"id"`
	Expires  time.Time `json:"expires"`
	IssuedAt string    `json:"issued_at"`
}

type Tokens []Token

type Client struct {
	Token      string
	endpoint   string
	apiVersion string
	client     http.Client
}

func NewClient(token string) Client {
	cli := Client{
		endpoint:   "https://identity.tyo1.conoha.io/",
		apiVersion: "v2.0",
		client:     http.Client{},
	}
	return cli
}

func (c *Client) GetVersion(apiVersion string) Version {
	url := c.endpoint + apiVersion

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	whole := struct {
		Version Version
	}{}
	err = json.Unmarshal(body, &whole)
	if err != nil {
		log.Fatal(err)
	}

	return whole.Version

}

func (c *Client) GetVersions() Versions {
	url := c.endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	res, err := c.client.Do(req)
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

	whole := struct {
		Versions struct {
			Values Versions
		}
	}{}
	err = json.Unmarshal(body, &whole)
	if err != nil {
		log.Fatal(err)
	}

	return whole.Versions.Values

}

func (c *Client) GetTokens(username, password, torentID string) Tokens {
	var tokens Tokens
	// params := struct {
	// 	auth struct {
	// 		passwordCredentials struct {
	// 			username string
	// 			password string
	// 		}
	// 	}
	// 	torentId string
	// }{}

	// whole := struct {
	// 	Access struct {
	// 		Metadata       Metadata `json:"metadata"`
	// 		User           User     `json:"user"`
	// 		Servicecatalog Services `json:"serviceCatalog"`
	// 		Token          Token    `json:"token"`
	// 	} `json:"access"`
	// }{}

	return tokens
}
