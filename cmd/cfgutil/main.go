package main

import (
	"os"

	"github.com/cebarks/TinGrizzly/internal/util"
	"github.com/pelletier/go-toml"
)

func main() {
	cfg := util.Config{}

	toml.Unmarshal([]byte{}, &cfg)

	file, _ := os.OpenFile("config.toml.example", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	enc := toml.NewEncoder(file).Order(toml.OrderPreserve)

	err := enc.Encode(&cfg)
	if err != nil {
		panic(err)
	}
	file.Sync()
}
