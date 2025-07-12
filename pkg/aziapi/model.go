package aziapi

type CountsResponse struct {
	TotalPlayers int                           `json:"total_players"`
	Games        map[string]CountsResponseGame `json:"games"`
}

type CountsResponseGame struct {
	Players int            `json:"players"`
	Modes   map[string]int `json:"modes,omitempty"`
}
