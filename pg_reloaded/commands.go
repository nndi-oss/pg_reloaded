package pg_reloaded

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunDropDatabase Executes a DROP DATABASE via psql
func RunDropDatabase(username, database, host string, port int, password string) error {
	if "postgres" == database {
		return errors.New("Nope, I cannot DROP the 'postgres' database.")
	}
	cmd := exec.Command("psql",
		dropDatabaseArgs(username, database, host, port))
	// cmd.Dir = db.Source.GetDir()
	cmd.Env = append(os.Environ(), fmt.Sprintf("PG_PASS=%s", password))
	// appLogger.Debug("Dropping database.",
	// 	"username", username,
	// 	"database", database)
	output, err := cmd.CombinedOutput()
	fmt.Println(output)
	if err != nil || !cmd.ProcessState.Success() {
		// appLogger.Error("Failed to run 'pg_restore'.",
		// 	"error", err,
		// 	"output", string(output))
		return err
	}

	return nil
}

// RunPgRestore Executes a database restore using pg_restore
func RunPgRestore(username, database, host string, port int, file, password string) error {
	cmd := exec.Command("pg_restore",
		fmt.Sprintf("%s %s", psqlArgs(username, database, host, port), file))
	// cmd.Dir = db.Source.GetDir()
	cmd.Env = append(os.Environ(), fmt.Sprintf("PG_PASS=%s", password))
	// appLogger.Info("Running restore via pg_restore.",
	// 	"database", database,
	// 	"file", file,
	// 	"username", username)
	output, err := cmd.CombinedOutput()
	fmt.Println(output)
	if err != nil || !cmd.ProcessState.Success() {
		// appLogger.Error("Failed to run 'pg_restore'.",
		// 	"error", err, "output", string(output))
		return err
	}

	return nil
}

// RunPsql Executes a command using psql
func RunPsql(username, database, host string, port int, file, password string) error {
	cmd := exec.Command("psql",
		fmt.Sprintf("%s < %s", psqlArgs(username, database, host, port), file))
	// cmd.Dir = db.Source.GetDir()
	cmd.Env = append(os.Environ(), fmt.Sprintf("PG_PASS=%s", password))
	// appLogger.Info("Running restore via psql.",
	// 	"database", database,
	// 	"file", file,
	// 	"username", username)
	output, err := cmd.CombinedOutput()
	fmt.Println(output)
	if err != nil || !cmd.ProcessState.Success() {
		// appLogger.Error("Failed to run 'psql'.",
		// 	"error", err, "output", string(output))
		return err
	}

	return nil
}

// dropDatabaseArgs Creates an argument string for passing to psql to drop a database
func dropDatabaseArgs(username, database, host string, port int) string {
	return fmt.Sprintf("%s -c \"DROP DATABASE %s\"",
		psqlArgs(username, "postgres", host, port), database)
}

// psqlArgs Creates an argument string for passing to Postgresql clients
func psqlArgs(username, database, host string, port int) string {
	return fmt.Sprintf("-U \"%s\" -d \"%s\" -h \"%s\" -p %d",
		username, database, host, port)
}

// DropAndRestoreUsingPsql Creates a command-line to drop a database and restore via Psql
func DropAndRestoreUsingPsql(username, database, host string, port int, file, password string) string {
	return fmt.Sprintf("psql %s && psql %s < %s",
		dropDatabaseArgs(username, database, host, port),
		psqlArgs(username, database, host, port),
		file,
	)
}
