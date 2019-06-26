package pg_reloaded

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/hashicorp/go-hclog"
)

func RunRestoreDatabase(psqlDir, username, database, host string, port int, file, password string) error {
	if "postgres" == database {
		return errors.New("Nope, I cannot CREATE the 'postgres' database.")
	}
	args := createDatabaseArgs(username, database, host, port)
	fmt.Println("Running", command(psqlDir, "psql"), args)
	cmd := exec.Command(command(psqlDir, "psql"), args...)
	// cmd.Dir = db.Source.GetDir()
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	hclog.Default().Debug("Restoring database.",
		"username", username,
		"database", database)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil || !cmd.ProcessState.Success() {
		hclog.Default().Error("Failed to run 'psql'.",
			"error", err,
			"output", string(output))
		return err
	}
	return RunPsql(psqlDir, username, database, host, port, file, password)
}

// RunDropDatabase Executes a DROP DATABASE via psql
func RunDropDatabase(psqlDir, username, database, host string, port int, password string) error {
	if "postgres" == database {
		return errors.New("Nope, I cannot DROP the 'postgres' database.")
	}
	args := dropDatabaseArgs(username, database, host, port)
	fmt.Println("Running", command(psqlDir, "psql"), args)
	cmd := exec.Command(command(psqlDir, "psql"), args...)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	hclog.Default().Debug("Dropping database.",
		"username", username,
		"database", database)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil || !cmd.ProcessState.Success() {
		hclog.Default().Error("Failed to run 'psql'.",
			"error", err,
			"output", string(output))
		return err
	}

	return nil
}

// RunPgRestore Executes a database restore using pg_restore
func RunPgRestore(psqlDir, username, database, host string, port int, file, password string) error {
	args := append(psqlArgs(username, database, host, port), file)
	fmt.Println("Running", command(psqlDir, "pg_restore"), args)
	cmd := exec.Command(command(psqlDir, "pg_restore"), args...)

	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	hclog.Default().Debug("Running restore via pg_restore.",
		"database", database,
		"file", file,
		"username", username)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil || !cmd.ProcessState.Success() {
		hclog.Default().Error("Failed to run 'pg_restore'.",
			"error", err, "output", string(output))
		return err
	}

	return nil
}

// RunPsql Executes a command using psql
func RunPsql(psqlDir, username, database, host string, port int, file, password string) error {
	args := append(psqlArgs(username, database, host, port), "-f", file)
	fmt.Println("Running", command(psqlDir, "psql"), args)
	cmd := exec.Command(command(psqlDir, "psql"), args...)
	// cmd.Dir = db.Source.GetDir()
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	hclog.Default().Debug("Running restore via psql.",
		"database", database,
		"file", file,
		"username", username)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil || !cmd.ProcessState.Success() {
		hclog.Default().Error("Failed to run 'psql'.",
			"error", err, "output", string(output))
		return err
	}

	return nil
}

// createDatabaseArgs Creates an argument string for passing to psql to CREATE a database
func createDatabaseArgs(username, database, host string, port int) []string {
	return append(
		psqlArgs(username, "postgres", host, port),
		"-c", fmt.Sprintf("CREATE DATABASE %s OWNER %s", database, username))
}

// dropDatabaseArgs Creates an argument string for passing to psql to DROP a database
func dropDatabaseArgs(username, database, host string, port int) []string {
	return append(
		psqlArgs(username, "postgres", host, port),
		"-c", fmt.Sprintf("DROP DATABASE %s", database))
}

// psqlArgs Creates an argument string for passing to Postgresql clients
func psqlArgs(username, database, host string, port int) []string {
	args := []string{
		"-U", username,
		"-h", host,
		"-p", fmt.Sprintf("%d", port),
		"-d", database,
	}

	return args
}

// command Returns command with base directory if provided or just the command name
func command(dir, commandName string) string {
	if dir == "" {
		return commandName
	}
	return path.Join(dir, commandName)
}

// DropAndRestoreUsingPsql Creates a command-line to drop a database and restore via Psql
func DropAndRestoreUsingPsql(psqlDir, username, database, host string, port int, file, password string) string {
	return fmt.Sprintf("psql X && psql Y %s", "yellow")
	//return fmt.Sprintf("psql %s && psql %s < %s",
	//	dropDatabaseArgs(username, database, host, port),
	//	psqlArgs(username, database, host, port),
	//	file,
	//)
}
