package tiles

type TileType struct {
	Id     string            `json:"id"`
	Sprite string            `json:"sprite"`
	Data   map[string]string `json:"data"`
}

var TileTypes map[string]*TileType = make(map[string]*TileType)
