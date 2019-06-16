package pg_reloaded

import (
	"testing"
)

func TestCreateDatabaseArgs(t *testing.T) {
	want := []string{
		"-U", "user",
		"-h", "my-host",
		"-p", "5432",
		"-d", "postgres",
		"-c", "CREATE DATABASE test_database OWNER user",
	}
	have := createDatabaseArgs("user", "test_database", "my-host", 5432)

	for idx, val := range want {
		if have[idx] != val {
			t.Errorf("TestCreateDatabaseArgs want: %s have:%s", val, have[idx])
		}
	}
}

func TestDropDatabaseArgs(t *testing.T) {
	want := []string{
		"-U", "user",
		"-h", "my-host",
		"-p", "5432",
		"-d", "postgres",
		"-c", "DROP DATABASE test_database",
	}
	have := dropDatabaseArgs("user", "test_database", "my-host", 5432)

	for idx, val := range want {
		if have[idx] != val {
			t.Errorf("TestDropDatabaseArgs want: %s have:%s", val, have[idx])
		}
	}
}

func TestPsqlArgs(t *testing.T) {
	want := []string{
		"-U", "user",
		"-h", "my-host",
		"-p", "5432",
		"-d", "test_database",
	}
	have := psqlArgs("user", "test_database", "my-host", 5432)
	
	for idx, val := range want {
		if have[idx] != val {
			t.Errorf("TestPsqlArgs want: %s have:%s", val, have[idx])
		}
	}
}
