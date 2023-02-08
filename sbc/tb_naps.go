package sbc

import (
	"encoding/json"
	"strings"
)

// TBNaps client
type TBNaps struct {
	Client
}

// TBNaps constructor (from Client)
func (c Client) TBNaps() TBNaps {
	return TBNaps{
		Client: c,
	}
}

func (c TBNaps) GetNap(config string, napName string) (*Nap, error) {
	var nap *Nap

	err := c.Client.Request("GET", "/configurations/"+config+"/naps/"+napName+"/", nil, &nap)
	if err != nil {
		return nil, err
	}

	return nap, nil
}

func (c TBNaps) GetNames(config string) ([]string, error) {
	d := make(map[string]json.RawMessage)

	err := c.Client.Request("GET", "/configurations/"+config+"/naps/", nil, &d)
	if err != nil {
		return nil, err
	}

	// thank u stackoverflow: https://stackoverflow.com/questions/17452722/how-to-get-the-key-value-from-a-json-string-in-go
	// a string slice to hold the keys
	k := make([]string, len(d)-1)
	// iteration counter
	i := 0
	// copy c's keys into k
	for s := range d {
		if strings.Contains(s, "meta") {
			continue
		}
		k[i] = s
		i++
	}

	return k, nil
}

func (c TBNaps) CreateNap(config string, nap Nap) error {
	err := c.Client.Request("POST", "/configurations/"+config+"/naps/", nap, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c TBNaps) UpdateNap(config string, nap Nap) error {
	err := c.Client.Request("PUT", "/configurations/"+config+"/naps/"+nap.Name, nap, nil)
	if err != nil {
		return err
	}
	return nil
}
