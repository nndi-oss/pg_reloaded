package pg_reloaded

import (
	"testing"
)

// TestValidatePsqlPath test validation of psqlDir
func TestValidatePsqlPath(t *testing.T) {
	config := Config{PsqlDir: "non-existent-dir"}

	want := "The path for Postgresql clients (psql_path) 'non-existent-dir' does not exist or cannot be read"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateServersNil(t *testing.T) {
	config := Config{Servers: nil}

	want := "Please specify at least one server under 'servers'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateServersEmpty(t *testing.T) {
	config := Config{Servers: []ServerConfig{}}

	want := "Please specify at least one server under 'servers'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesNil(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: nil,
	}

	want := "Please specify at least one database under 'databases'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesEmpty(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{},
	}

	want := "Please specify at least one database under 'databases'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesEmptyName(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name: "",
			},
		},
	}

	want := "Please specify the name for database at index: 0"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesEmptyServer(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Name:     "local",
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name:   "dev",
				Server: "",
			},
		},
	}

	want := "Please specify the name for the server for database 'dev'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesInvalidServer(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Name:     "local",
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name:   "dev",
				Server: "localhost",
			},
		},
	}

	want := "Server for database 'dev' does not exist in 'servers' list"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesEmptySchedule(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Name:     "local",
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name:     "dev",
				Server:   "local",
				Schedule: "",
			},
		},
	}

	want := "Please provide a 'schedule' for database 'dev'"
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesInvalidSourceType(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Name:     "local",
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name:     "dev",
				Server:   "local",
				Schedule: "@every 4h",
				Source: SourceConfig{
					Type: "invalid",
				},
			},
		},
	}

	want := "Provided source type 'invalid' is not supported for database 'dev'."
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}

func TestValidateDatabasesInvalidSourceFile(t *testing.T) {
	config := Config{
		Servers: []ServerConfig{
			ServerConfig{
				Name:     "local",
				Host:     "localhost",
				Port:     5432,
				Username: "user",
				Password: "password",
			},
		},
		Databases: []DatabaseConfig{
			DatabaseConfig{
				Name:     "dev",
				Server:   "local",
				Schedule: "@every 24h",
				Source: SourceConfig{
					Type: "sql",
					File: "non-existent-file",
				},
			},
		},
	}

	want := "File 'non-existent-file' does not exist. File must be provided for source type 'sql' for database 'dev'."
	err := Validate(config)
	if err == nil {
		t.Error("Expected an error from Validate()")
		return
	}
	have := err.Error()
	if want != have {
		t.Errorf("Error\n\twant:%s\n\thave:%s", want, have)
	}
}
