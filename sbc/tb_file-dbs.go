package sbc

import (
	"encoding/json"
	"github.com/gocarina/gocsv"
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

func (c TBFileDBs) GetDigitMap(config string, dMap string) ([]TBDigitMap, error) {
	var file *TBFile
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_digitmaps/"+strings.ReplaceAll(dMap, ".", "%2E"),
		nil, &file)
	if err != nil {
		return nil, err
	}

	var digitM []TBDigitMap
	reader := *strings.NewReader(file.Content)
	err = gocsv.Unmarshal(&reader, &digitM)
	if err != nil {
		return nil, err
	}

	return digitM, nil
}

func (c TBFileDBs) UpdateDigitMap(config string, digitMapFile string, digitMap []TBDigitMap) error {
	file := TBFile{
		Name: digitMapFile,
	}

	marsh, err := gocsv.MarshalString(&digitMap)
	if err != nil {
		return err
	}

	formatted := strings.ReplaceAll(marsh, "\n", "\r\n")

	file.Content = formatted

	err = c.Client.Request("PUT", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_digitmaps/"+strings.ReplaceAll(digitMapFile, ".", "%2E"),
		file, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c TBFileDBs) GetRouteDef(config string, rDef string) ([]TBRouteDef, error) {
	var file *TBFile
	// File_DB is default?
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/"+strings.ReplaceAll(rDef, ".", "%2E"),
		nil, &file)
	if err != nil {
		return nil, err
	}

	var routeD []TBRouteDef
	reader := *strings.NewReader(file.Content)
	err = gocsv.Unmarshal(&reader, &routeD)
	if err != nil {
		return nil, err
	}

	return routeD, nil
}

func (c TBFileDBs) DeleteRouteDef(config string, rDef string) error {
	// File_DB is default?
	err := c.Client.Request("DELETE", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/"+strings.ReplaceAll(rDef, ".", "%2E"),
		nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c TBFileDBs) UpdateRouteDef(config string, routeDefFile string, routeDef []TBRouteDef) error {
	file := TBFile{
		Name: routeDefFile,
	}

	marsh, err := gocsv.MarshalString(&routeDef)
	if err != nil {
		return err
	}

	formatted := strings.ReplaceAll(marsh, "\n", "\r\n")

	file.Content = formatted

	err = c.Client.Request("PUT", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/"+strings.ReplaceAll(routeDefFile, ".", "%2E"),
		file, nil)

	if err != nil {
		return err
	}

	return nil
}

func (c TBFileDBs) CreateRouteDef(config string, file TBFile) error {
	// File_DB is default?
	err := c.Client.Request("POST", "/configurations/"+config+"/file_dbs/"+
		c.fileDbPath+"/routesets_definitions/",
		file, nil)
	if err != nil {
		return err
	}
	return nil
}

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
