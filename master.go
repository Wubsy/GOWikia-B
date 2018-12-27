package GOWikia_Wubsy

import (
	"net/url"
	"errors"
)

func NewClient(url string)(*WikiaApi, error){
	valid, err := isValidUrl(url)
	if !valid {
		return nil, err
	}
	return &WikiaApi{url}, nil
}

func isValidUrl(u string) (bool, error) {
	ur, err := url.Parse(u)
	if err != nil {
		return false, err
	}
	if ur.Scheme == "http" || ur.Scheme == "https" {
		return true, nil
	}
	return false, errors.New("Invalid protocol provided, only http and https are allowed")
}
