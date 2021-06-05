package main

import (
	"os"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/pelletier/go-toml"
)

func main() {
	cfg := util.Config{}

	toml.Unmarshal([]byte{}, &cfg)

	bytes, _ := toml.Marshal(&cfg)

	os.WriteFile("config.toml.example", bytes, 0644)
}
