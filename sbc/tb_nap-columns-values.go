package sbc

import (
	"encoding/json"
	"strings"
)

// TBNapColumnsValues client
type TBNapColumnsValues struct {
	Client
}

// TBNapColumnsValues constructor (from Client)
func (c Client) TBNapColumnsValues() TBNapColumnsValues {
	return TBNapColumnsValues{
		Client: c,
	}
}

/*
"routesets_definition": {},
"route_groups": {},
"routesets_digitmap": {},
"weight": {},
"black_white_list": {},
"called_pre_remap": {},
"priority": {},
*/

func (c TBNapColumnsValues) GetAllValues(config string) (*Nap, error) {
	d := make(map[string]json.RawMessage)

	err := c.Client.Request("GET", "/configurations/"+config+"/nap_columns_values/", nil, &d)
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

	for _, n := range k {
		err := c.Client.Request("GET", "/configurations/"+config+"/nap_columns_values/"+n, nil, nil)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
