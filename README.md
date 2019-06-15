PG Reloaded
===

PG Reloaded is a simple tool to help developers restore PostgreSQL databases 
periodically. Useful for databases used for online demos where you
want to reset the demo data after users have played with your system and also
for local development where you can schedule your databases to be restored from a dump. 

## Installation

Download a binary from the Releases page.
If you would like to build from source, see below.

## Usage

In order to be effective, `pg_reloaded` needs to run in the background as a daemon.
There's a sample [Systemd Unit File here](./config/pg_reloaded.service).

You can also use supervisor to run the `pg_reloaded` daemon, below is an example
configuration for [Supervisor](https://github.com/supervisor/supervisor)

```ini
[program:pg_reloaded]
command=/usr/bin/pg_reloaded start --config=/etc/pg_reloaded/pg_reloaded.yml
```
Please note that these process management systems must be configured to 
start pg_reloaded on boot for best effective use of the scheduling capabilities.

You can also run the pg_reloaded binary in the background by adding the `--daemonize` flag.

**Example**

```sh
$ pg_reloaded check --config="pg_reloaded.yml"
$ pg_reloaded start

# Or with a path to a configuration file
$ pg_reloaded start --daemonize --config="development.yml" --log-file="./path/to/log"
```

If you would like to restore a database immediately, run the following:

```sh
$ pg_reloaded run "database"
```

You can also override the user, host and port from the configuration by passing
the arguments on the command-line.

```sh 
$ pg_reloaded run "database" --username="postgres" --host="remote-server" --port=15432 
```

**Usage**

Get the usage information by running `pg_reloaded --help`

## Example configuration

PG Reloaded is configured via YAML configuration file which records the details
about how and when to restore your databases.

By default `pg_reloaded` reads configuration from a file named `pg_reloaded.yml`
in the current directory or from file `$HOME/pg_reloaded.yml`
(on Windows in `%UserProfile%\pg_reloaded.yml`) if present

You can specify a path for the configuration file via the `--config` option
on the command-line.

The configuration basically looks like the following, add databases:

```yaml
# Absolute path to the directory containing postgresql client programs
# The following client programs are searched for specifically:
# psql, pg_restore, pg_dump
psql_path: "/path/to/psql-dir"
# Absolute path to the logfile, will be created if it does not exist
log_file: "/path/to/logfile"
servers:
  # name - A name to identify the server in the "databases" section
  #        of the configuration
  - name: "my-development-server"
    # port - The host for the database
    host: "localhost"
    # port - The port for the database
    port: 5432
    # Username for the user must have CREATE DATABASE & DROP DATABASE privileges
    username: "appuser" 
    # Password for the user role on the database 
    password: "password"
databases:
  # name - The database name 
  - name: "my_database_name"
    # server - The name of a server in the "servers" list
    server: "my-development-server"
    # schedule - Specifies the schedule for running the database restores in
    # daemon mode. Supports simple interval notation and CRON expressions
    schedule: "@every 24h"
    # Source specifies where to get the schema and data to restore  
    source:
      # The type of file(s) to restore the database from.
      # The following types are (will be) supported:
      #
      # * sql - load schema & data from SQL files using psql
      # * tar - load schema & data from SQL files using pg_restore
      # * csv - load data from CSV files using pgfutter
      # * json - load data from JSON files using pgfutter
      type: "sql"
      
      # The absolute path to the file to restore the database from 
      file: "/path/to/file"
      
      # The absolute path to the schema file to be used to create tables, functions etc..
      # Schema MUST be specified if source type is one of: csv, json 
      # or if the SQL file only contains data
      schema: "/path/to/schema/file.sql"
```

### Supported notation for Scheduling

The real value of pg_reloaded is in it's ability to restore databases according
to a schedule. 

The following syntax is supported for scheduling

* Intervals: Simple interval notation is supported. Only for minutes (`m`),
hours (`h`) and days (`d`).

e.g. `@every 10m`, `@every 2h`, `@every 7d`, `@every 30d`

* CRON Expression: Most CRON expressions valid [here](http://crontab.guru) are valid


## Prerequisites

* The postgresql client programs must be present on your path or configured in 
the config file or command-line for `pg_reloaded` to work. In particular the 
program may to execute `psql`, `pg_restore`, `pg_dump` during it's operation.

In the YAML file:

```yaml
psql_path: "/path/to/psql-dir/"
```

On the command-line

```sh
$ pg_reloaded --psql-path="/path/to/psql-dir" 
```

### Supported Sources

Currently, the only supported sources are:

* *SQL* via dumped *SQL file* - default source, load dumps/data files from the filesystem

## Building from Source

I encourage you to build pg_reloaded from source, if only to get you to try out
Go ;). So the first step is to Install Go.

The next step is to clone the repo and then after that, building the binary _should_
be as simple as running `go build`

```sh
$ git clone https://github.com/zikani03/pg_reloaded.git
$ cd pg_reloaded
$ go build
```

### Dependencies

This project could not be made possible without these great Open-Source 
tools and their authors/contributors whose shoulders are steady enough to stand on:

* [Go 1.11.x](https://golang.org)
* [Postgresql (Ofcourse!)](https://postgresql.org)
* [cobra](https://github.com/spf13/cobra)
* [viper](https://github.com/spf13/viper)
* The [cron](./cron) package is copied from [dkron](https://github.com/victorcoder/dkron/tree/master/cron)

## CHENJEZO ( *Notice* )

- **Running as a daemon on Windows**

Since you have to run it in the background for the scheduling functionality to be
of any value a service would be ideal on Windows - but until then you will have 
to run it in the foreground.

- **BE CAREFULL, IT MAY EAT YOUR LUNCH!**

This is not meant to be run on *production* databases which house critical data
that you can't afford to lose. It's meant for demo and development databases that
can be dropped and restored without losing a dime. Use good judgment.


## ROADMAP

Some rough ideas on how to take this thing further:

* Add preload and post-load conditions for databases

```yaml
preload:
  dump: true
  checkActivity: true
postload:
  notify: true
```

* Backup the current database before restoring using pg_dump or [pgclimb](https://github.com/lukasmartinelli/pgclimb)
* A Windows Service wrapper
* Add suport for the following sources:

  * *csv* : loads data from CSV files, just like you would with [pgfutter](https://github.com/lukasmartinelli/pgfutter)
  * *json*: loads data from JSON files, just like you would with pgfutter
  * *postgres* - load database from a remote postgresql database
  * *http* - load a database dump from an HTTP server
  * *s3* - load database dump from AWS S3

## CONTRIBUTING

Issues and Pull Requests welcome.

## License

MIT License

---

Copyright (c) 2019, Zikani Nyirenda Mwase