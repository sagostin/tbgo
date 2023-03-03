package sbc

type TBRoutes struct {
	Client
}

// TBRoutes constructor (from Client)
func (c Client) TBRoutes() TBSystem {
	return TBSystem{
		Client: c,
	}
}

type Route struct {
	RemappedCalled                string `json:"remapped_called"`
	RemappedDestinationLegProfile string `json:"remapped_destination_leg_profile"`
	Called                        string `json:"called"`
	RemappedSourceLegProfile      string `json:"remapped_source_leg_profile"`
	RemappedNap                   string `json:"remapped_nap"`
	Nap                           string `json:"nap"`
	RemappedCalling               string `json:"remapped_calling"`
	Calling                       string `json:"calling"`
	CustomAttributes              struct {
		RouteGroup             string `json:"route_group"`
		PrivateAddress         string `json:"private_address"`
		Weight                 string `json:"weight"`
		ForwardSipDomain       string `json:"forward_sip_domain"`
		ForwardSipParameters   string `json:"forward_sip_parameters"`
		RemappedPrivateAddress string `json:"remapped_private_address"`
		Priority               string `json:"priority"`
	} `json:"custom_attributes"`
	Name         string `json:"name"`
	RoutesetName string `json:"routeset_name"`
}

// we can't use the auto generation for the routes, so we need to manually create the routes
// the format for the names will be to_pbx_CUSTOMER, this is just to distinguish the correct name

func (c TBRoutes) GetRoute(routeStr string) (*Route, error) {
	var route *Route

	err := c.Client.Request("GET", "/routes/"+routeStr, nil, &route)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (c TBRoutes) Create(route Route) error {
	err := c.Client.Request("POST", "/routes/"+route.Name, route, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c TBRoutes) Update(route Route) error {
	err := c.Client.Request("PUT", "/routes/"+route.Name, route, nil)
	if err != nil {
		return err
	}
	return nil
}
