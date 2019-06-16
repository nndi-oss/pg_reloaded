package pg_reloaded

import (
	"testing"
)

func TestDropDatabaseArgs(t *testing.T) {
	want := "-U \"user\" -d \"postgres\" -h \"my-host\" -p 5432 -c \"DROP DATABASE test_database\""
	have := dropDatabaseArgs("user", "test_database", "my-host", 5432)

	if want != have {
		t.Errorf("TestDropDatabaseArgs want: %s have:%s", want, have)
	}
}

func TestPsqlArgs(t *testing.T) {
	want := "-U \"user\" -d \"test_database\" -h \"my-host\" -p 5432"
	have := psqlArgs("user", "test_database", "my-host", 5432)

	if want != have {
		t.Errorf("TestPsqlArgs want: %s have:%s", want, have)
	}
}
