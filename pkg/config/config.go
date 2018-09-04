package config

import (
	"encoding/json"
	"fmt"

	"github.com/caicloud/simple-object-storage/pkg/constants"
)

type Config struct {
	// db
	DbString string
}

func NewDefaultConfig() *Config {
	return &Config{
		DbString: constants.DefaultDbString,
	}
}

func (c *Config) Validate() error {
	if len(c.DbString) == 0 {
		return fmt.Errorf("empty database string")
	}
	return nil
}
func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}
