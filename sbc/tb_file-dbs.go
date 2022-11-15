package sbc

// file_dbs

// TBFileDBs client
type TBFileDBs struct {
	Client
}

// TBFileDBs constructor (from Client)
func (c Client) TBFileDBs() TBFileDBs {
	return TBFileDBs{
		Client: c,
	}
}

func (c TBFileDBs) Get(config string) (*Nap, error) {
	// why do they do /file_dbs/File_DB... can you even have multiple?? wot? or is it config based?
	err := c.Client.Request("GET", "/configurations/"+config+"/file_dbs/File_DB", nil, nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
