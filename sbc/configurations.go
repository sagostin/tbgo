package sbc

// Configurations  client
type Configurations struct {
	Client
}

// Configurations constructor (from Client)
func (c Client) Configurations() Configurations {
	return Configurations{
		Client: c,
	}
}
