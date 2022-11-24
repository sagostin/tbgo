package sbc

import (
	"encoding/json"
	"strings"
)

// file_dbs

// TBFileDBs client
type TBFileDBs struct {
	Client
	fileDbPath string
}

// TBFileDBs constructor (from Client)
func (c Client) TBFileDBs(fileDb string) TBFileDBs {
	return TBFileDBs{
		Client: c,
		// todo why do they do /file_dbs/File_DB... can you even have multiple?? wot? or is it config based?
		fileDbPath: fileDb,
	}
}

func (c TBFileDBs) Get(config string) (*Nap, error) {
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+c.fileDbPath, nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c TBFileDBs) GetCustomFileNames(config string) ([]string, error) {
	d := make(map[string]json.RawMessage)
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+c.fileDbPath+"/custom_files", nil, &d)
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

func (c TBFileDBs) GetRouteDefNames(config string) ([]string, error) {
	// dec0de_routedef%2Ecsv
	d := make(map[string]json.RawMessage)
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+c.fileDbPath+"/routesets_definitions/", nil, &d)
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

func (c TBFileDBs) GetDigitMapsNames(config string) ([]string, error) {
	d := make(map[string]json.RawMessage)
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+c.fileDbPath+"/routesets_digitmaps/", nil, &d)
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

func (c TBFileDBs) GetDigitMap(config string, dMap string) (*TBFile, error) {
	var file *TBFile
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_digitmaps/"+strings.ReplaceAll(dMap, ".", "%2E"),
		nil, &file)
	if err != nil {
		return file, err
	}
	return file, nil
}

func (c TBFileDBs) UpdateDigitMap(config string, digitMap string) error {
	err := c.Client.Request("PUT", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_digitmaps/"+strings.ReplaceAll(digitMap, ".", "%2E"),
		digitMap, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c TBFileDBs) GetRouteDef(config string, rDef string) (*TBFile, error) {
	var file *TBFile
	// File_DB is default?
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/"+strings.ReplaceAll(rDef, ".", "%2E"),
		nil, &file)
	if err != nil {
		return file, err
	}
	return file, nil
}

func (c TBFileDBs) CreateRouteDef(config string, rDef string, file TBFile) error {
	// File_DB is default?
	err := c.Client.Request("POST", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/"+strings.ReplaceAll(rDef, ".", "%2E"),
		file, nil)
	if err != nil {
		return err
	}
	return nil
}

//todo update routedef??

/*
func (c TBFileDBs) GetRadiusDirectories(config string) (*Nap, error) {
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/File_DB/radius_dictionaries", nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c TBFileDBs) GetLawEnforcementTargets(config string) (*Nap, error) {
	// todo why do they do /file_dbs/File_DB... can you even have multiple?? wot? or is it config based?
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/File_DB/law_enforcement_targets", nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}*/

// TODO
/*"radius_dictionaries": {},
"law_enforcement_targets": {},
"name": "File_DB",*/
