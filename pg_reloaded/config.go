package pg_reloaded

import (
	"errors"
	"fmt"
	"github.com/zikani03/pg_reloaded/cron"
	"os"
	"regexp"
	"strings"
)

// Config stores all the configuration information for pg_reloaded
type Config struct {
	Daemonize bool
	PsqlDir   string           `mapstructure:"psql_path"`
	LogPath   string           `mapstructure:"log_path"`
	Servers   []ServerConfig   `mapstructure:"servers"`
	Databases []DatabaseConfig `mapstructure:"databases"`
}

type ServerConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type DatabaseConfig struct {
	Name     string       `mapstructure:"name"`
	Server   string       `mapstructure:"server"`
	Schedule string       `mapstructure:"schedule"`
	Source   SourceConfig `mapstructure:"source"`
}

type SourceConfig struct {
	Type   string   `mapstructure:"type"`
	File   string   `mapstructure:"file"`
	Files  []string `mapstructure:"files"`
	Schema string   `mapstructure:"schema"`
}

// GetServerByName Gets a server by name from the list of servers
// returns nil if not available
func (c Config) GetServerByName(name string) *ServerConfig {
	for _, server := range c.Servers {
		if name == server.Name {
			return &server
		}
	}
	return nil
}

func Validate(cfg Config) error {
	if cfg.PsqlDir != "" {
		if _, err := os.Stat(cfg.PsqlDir); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("The path for Postgresql clients (psql_path) '%s' does not exist or cannot be read", cfg.PsqlDir))
		}
	}

	if cfg.Servers == nil || len(cfg.Servers) < 1 {
		return errors.New("Please specify at least one server under 'servers'")
	}

	if cfg.Databases == nil || len(cfg.Databases) < 1 {
		return errors.New("Please specify at least one database under 'databases'")
	}
	for idx, d := range cfg.Databases {
		if d.Name == "" {
			return errors.New(fmt.Sprintf("Please specify the name for database at index: %d", idx))
		}
		if d.Server == "" {
			return errors.New(fmt.Sprintf("Please specify the name for the server for database '%s'", d.Name))
		}
		if s := cfg.GetServerByName(d.Server); s == nil {
			return errors.New(fmt.Sprintf("Server for database '%s' does not exist in 'servers' list", d.Name))
		}
		if d.Schedule == "" {
			return errors.New(fmt.Sprintf("Please provide a 'schedule' for database '%s'", d.Name))
		}
		if _, err := cron.Parse(d.Schedule); err != nil {
			return errors.New(fmt.Sprintf("Invalid expression for 'schedule' for database '%s'. Error: %v", d.Name, err))
		}
		stype := strings.ToLower(d.Source.Type)

		matched, err := regexp.MatchString("^(sql|tar)$", stype)
		if err != nil {
			return err
		}
		// TODO: Add conditions when CSV and JSON are supported || stype != "csv" || stype != "json"
		if !matched {
			return errors.New(fmt.Sprintf("Provided source type '%s' is not supported for database '%s'.", d.Source.Type, d.Name))
		} else {
			if _, err := os.Stat(d.Source.File); os.IsNotExist(err) {
				return errors.New(fmt.Sprintf("File '%s' does not exist. File must be provided for source type '%s' for database '%s'.", d.Source.File, d.Source.Type, d.Name))
			}
		}

		// TODO: Check configuration for CSV and JSON source types i.e. d.Source.Schema, d.Source.Files
		// if stype == "csv" || stype == "json"  {
		// if _, err := os.Stat(d.Source.Schema); os.IsNotExist(err) {
		//		return errors.New(fmt.Sprintf("Schema File '%s' does not exist - a schema must be provided for source type '%s' for database '%s'.", d.Source.Schema, d.Source.Type, d.Name))
		//	}
		//  for _, file := range d.Source.Files {
		// 	    if _, err := os.Stat(d.Source.File); os.IsNotExist(err) {
		// 		    return errors.New(fmt.Sprintf("File '%s' does not exist. File must be provided for source type '%s' for database '%s'.", d.Source.File, d.Source.Type, d.Name))
		// 	    }
		//   }
		// }
	}
	return nil
}
