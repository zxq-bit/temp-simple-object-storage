package config

import (
	"encoding/json"
	"fmt"

	"github.com/caicloud/simple-object-storage/pkg/constants"
)

type Config struct {
	// db
	DatabaseString string `desc:"database connect string, dir path for local db"`
	// storage
	StorageString string `desc:"storage connect string, dir path for local storage or mounted volume"`
	// all
	RootAllowed bool `desc:"local db/storage check option, is allow to use path under root"`
}

func NewDefaultConfig() *Config {
	return &Config{
		DatabaseString: constants.DefaultDatabaseString,
		StorageString:  constants.DefaultStorageString,
		RootAllowed:    false,
	}
}

func (c *Config) Validate() error {
	if len(c.DatabaseString) == 0 {
		return fmt.Errorf("empty database string")
	}
	if len(c.StorageString) == 0 {
		return fmt.Errorf("empty storage string")
	}
	return nil
}
func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}
