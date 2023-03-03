package sbc

type TBSystem struct {
	Client
}

// TBSystem constructor (from Client)
func (c Client) TBSystem() TBSystem {
	return TBSystem{
		Client: c,
	}
}

func (c TBSystem) GetSystem(system string) (*Nap, error) {
	var nap *Nap

	err := c.Client.Request("GET", "/system/"+system, nil, nil)
	if err != nil {
		return nil, err
	}

	return nap, nil
}
