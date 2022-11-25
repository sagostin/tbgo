package sbc

type TBDigitMap struct {
	Called       string `json:"called"csv:"called"`
	Calling      string `json:"calling"csv:"calling"`
	RouteSetName string `json:"routeset_name"csv:"routeset_name"`
}

type TBRouteDef struct {
	RouteSetName string `json:"routeset_name"csv:"routeset_name"`
	Priority     int    `json:"priority"csv:"priority"`
	Weight       int    `json:"weight"csv:"weight"`
	RouteGroup   string `json:"route_group"csv:"route_group"`
}
