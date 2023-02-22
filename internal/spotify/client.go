package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type HttpClient struct {
	Client      http.Client
	BaseUrl     string
	AccessToken string
	Owner       string
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	return c.Client.Do(req)
}

// method must have no type parameters -> can't use receiver :( hence passing client ref as param
func GetApiStruct[T any](c *HttpClient, url string) (T, error) {
	var resStruct T

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resStruct, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	res, err := c.Do(req)
	if err != nil {
		return resStruct, err
	}

	if res.StatusCode != http.StatusOK {
		return resStruct, errors.Errorf("API request to %s failed %d: ", url, res.StatusCode)
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(err.Error())
	}

	if err = json.Unmarshal(body, &resStruct); err != nil {
		return resStruct, err
	}

	return resStruct, nil
}
