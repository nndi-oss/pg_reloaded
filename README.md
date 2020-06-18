<p align="center">
  <img src="logo.svg" alt="PG Reloaded logo" height="128px" >
</p>

PG Reloaded
===

`pg_reloaded` is a program that's useful for restoring databases. You can use it to refresh databases for online demos, development databases and anywhere where you may want to reset the data after use. You schedule your databases to be restored from a backup. 

## Installation

Currently, you will have to build it from source, binary Releases will be made available soon.

## Usage

Get the usage information by running `pg_reloaded --help`

```sh
$ pg_reloaded check --config="pg_reloaded.yml"

# Run with the default configuration
$ pg_reloaded start

# Or with a path to a configuration file
$ pg_reloaded start --config="development.yml" --log-file="./path/to/log"
```

If you would like to restore a database immediately, run the following:

```sh
$ pg_reloaded run "database"
```

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

## Example configuration

PG Reloaded is configured via YAML configuration file which records the details
about how and when to restore your databases.

By default `pg_reloaded` reads configuration from a file named `pg_reloaded.yml`
in the home directory if present i.e. `$HOME/pg_reloaded.yml`
(on Windows in `%UserProfile%\pg_reloaded.yml`)

You can specify a path for the configuration file via the `--config` option
on the command-line.

The configuration basically looks like the following:

```yaml
# Absolute path to the directory containing postgresql client programs
# The following client programs are searched for specifically:
# psql, pg_restore, pg_dump
psql_path: "/path/to/psql-dir"
# Absolute path to the logfile, will be created if it does not exist
log_file: "/path/to/logfile"
servers:
  # name - A name to identify the server in the "databases" section of the configuration
  - name: "my-development-server"
    # host - The host for the database
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

The real value of pg_reloaded is in its ability to restore databases according
to a schedule. 

The following syntax is supported for scheduling

* Intervals: Simple interval notation is supported. Only for seconds(`s`), 
minutes (`m`) and hours (`h`). 

e.g. `@every 10m`, `@every 2h`, `@weekly`, `@monthly`

* CRON Expression: Most CRON expressions valid [here](http://crontab.guru) should be valid


## Prerequisites

* The postgresql client programs must be present on your path or configured in 
the config file or command-line for `pg_reloaded` to work. In particular the 
program may need to execute `psql`, `pg_restore` or `pg_dump` during it's operation.

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

## Docker

You can build Docker images/containers using the [Dockerfile](./Dockerfile).
At this time, you can pull images from [@gkawamoto's](https://github.com/gkawamoto) Dockerhub:

```
$ docker pull gkawamoto/pg_reloaded
```

## Building from Source

I encourage you to build pg_reloaded from source, if only to get you to try out
Go ;). So the first step is to Install Go.

The next step is to clone the repo and then after that, building the binary _should_
be as simple as running `go build`

```sh
$ git clone https://github.com/zikani03/pg_reloaded.git
$ cd pg_reloaded
$ go build -o dist/pg_reloaded
```

### Dependencies

This project could not be made possible without these great Open-Source 
tools and their authors/contributors whose shoulders are steady enough to stand on:

* [Go 1.12.x](https://golang.org)
* [Postgresql (Ofcourse!)](https://postgresql.org)
* [cobra](https://github.com/spf13/cobra)
* [viper](https://github.com/spf13/viper)
* The [cron](./cron) package is copied from [dkron](https://github.com/victorcoder/dkron/tree/master/cron)

## CHENJEZO ( *Notice* )

- **Running as a Service on Windows**

Since you have to run it in the background for the scheduling functionality to be
of any value a Service wrapper would be ideal on Windows - but until then you will have 
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
* Add support for the following sources:

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

Copyright (c) 2019 - 2020, NNDI 
