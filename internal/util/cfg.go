package util

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Core struct {
		LogLevel string `toml:"LogLevel" default:"info"`

		Tunables struct {
			Ups int `toml:"UpdatesPerSecondTarget" default:"60"`
			Fps int `toml:"FramesPerSecondTarget" default:"60"`

			MaxWorldPoolWorkers int `toml:"MaxWorldPoolWorkers" default:"5000"`
		} `toml:"tunables" comment:"low level tuning engine settings"`
	} `toml:"core" comment:"Core Engine settings"`
}

var cfg *Config

//Cfg returns the active cfg.Config holding the config information
func Cfg() *Config {
	if cfg == nil {
		cfg = ReadConfig()
	}
	return cfg
}

//ReloadCfgFromDisk reloads and reparses the config file from disk
func ReloadCfgFromDisk() {
	cfg = ReadConfig()
}

//ReadConfig parses the config file into a Config struct
func ReadConfig() *Config {
	config := &Config{}
	pwd, _ := os.Getwd()
	// configFile := "config.toml"
	// configFile := pwd + "\\config.toml"
	configFile := path.Join(pwd, "config.toml")

	// if !util.FileExists(configFile) {
	// 	config.configDir = ""

	// 	var bytes bytes.Buffer
	// 	err := toml.NewEncoder(&bytes).Order(toml.OrderPreserve).Encode(config)

	// 	if err != nil {
	// 		log.Panic().Err(err).Msg("Unable to save sample config file.")
	// 	}

	// 	ioutil.WriteFile(configFile, bytes.Bytes(), 0644)
	// 	log.Fatal().Msgf("Config file doesn't exist. An example has been saved in its place.")
	// }

	// Read config from the file
	bytes, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to read config file at: '%s'", configFile)
	}

	// Unmarshal the config file bytes into a Config struct
	err = toml.Unmarshal(bytes, config)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse config file.")
	}

	log.Debug().Msgf("Read config file: %s", configFile)
	log.Trace().Msgf("Config Struct: %+v", *config)

	logLevel, err := zerolog.ParseLevel(config.Core.LogLevel)
	if err != nil {
		log.Info().Msgf("Supplied config file log level (%s) is invalid. Defaulting to info.", config.Core.LogLevel)
		logLevel = zerolog.InfoLevel
	}

	log.Info().Msgf("Log Level set to: %s", logLevel.String())

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

	return config
}
