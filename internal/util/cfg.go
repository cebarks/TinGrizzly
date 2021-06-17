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
			Ups int64 `toml:"UpdatesPerSecondTarget" default:"60"`
			Fps int64 `toml:"FramesPerSecondTarget" default:"60"`

			MaxWorldPoolWorkers int `toml:"MaxWorldPoolWorkers" default:"32"`
		} `toml:"tunables" comment:"low level engine settings"`
	} `toml:"core" comment:"Core Engine settings"`

	Graphics struct {
		Resolution string `toml:"resolution" default:"1920x1080"`
		Vsync      bool   `toml:"vsync" default:"false"`
		Fullscreen bool   `toml:"fullscreen" default:"false"`
		Resources  struct {
			Embedded bool `toml:"embedded" comment:"Should resources be loaded from the filesystem or the binary"`
		} `toml:"resources" comment:"resources related settings"`
	} `toml:"graphics" comment:"graphics related settings"`
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

	configFile := path.Join(pwd, "config.toml")

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
