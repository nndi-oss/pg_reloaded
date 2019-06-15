package pg_reloaded

import (
	"github.com/zikani03/pg_reloaded/cron"
)

// Config stores all the configuration information for pg_reloaded
type Config struct {
	Daemonize bool 
	PsqlDir string `mapstructure:"psql_path"`
	LogPath string `mapstructure:"log_path"`
	Servers []ServerConfig `mapstructure:"servers"` 
	Databases []DatabaseConfig `mapstructure:"databases"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type DatabaseConfig struct {
	Name string `mapstructure:"name"`
	Server string `mapstructure:"server"`
	Schedule string `mapstructure:"schedule"`
	Source SourceConfig `mapstructure:"source"`
}

type SourceConfig struct {
	Type string `mapstructure:"type"`
	File string `mapstructure:"file"`
	Files []string `mapstructure:"files"`
	Schema string `mapstructure:"schema"`
}

// GetServerByName Gets a server by name from the list of servers 
// returns nil if not available 
func (c Config) GetServerByName(name string) ServerConfig {
	for _, server := range config.Servers {
		if name == server.Name {
			return server
		}		
	}
	return nil
}

func Validate(cfg Config) error {
	if cfg.PsqlDir == "" {
		return errors.New("psql-path must be specified")
	}
	
	if cfg.Server == nil {
		return errors.New("Please specify atleast one server under 'server' must be specified")
	}

	if cfg.Databases == nil {
		return errors.New("Please specify atleast one database under 'databases' must be specified")
	}
	for idx, d := range cfg.Databases {
		if d.Name == "" {
			return errors.New("Please specify the name for database at index: %s", idx)
		}
		if d.Server == "" {
			return errors.New("Please specify the name for the server for database '%s'", d.Name)
		}
		if s = cfg.GetServerByName(d.Server); s == nil {
			return errors.New(fmt.Sprintf("Server for database '%s' does not exist in 'servers' list", d.Name))
		}
		if d.Schedule == "" {
			return errors.New("Please provide a 'schedule' for database '%s'", d.Name)
		}
		if err := cron.ParseSchedule(d.Schedule); err != nil {
			return errors.New("Invalid expression for 'schedule' for database '%s'. Error: %v", d.Name, err)
		}
	}
}
