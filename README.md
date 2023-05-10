# pg-ping

`pg-ping` is a command line utility to continuously ping your postgres. This is useful for checking for downtime when making changes to your Postgres instance.

## Installation

```bash
# Go 1.18>=
go install github.com/luisnquin/pg-ping@latest 
```

## Usage

```bash
NAME:
   pg-ping - Ping your postgres continuously

USAGE:
   pg-ping [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --once                       Ping only once and quit [$PGPING_ONCE]
   --debug                      Print debug logs [$PGPING_DEBUG]
   --username value, -U value   Username to connect to postgres [$PGPING_USERNAME]
   --password value, -p value   Password to connect to postgres [$PGPING_PASSWORD]
   --host value, -h value       Host of postgres server [$PGPING_HOST]
   --dbname value, -d value     DBName to connect to [$PGPING_DBNAME]
   --frequency value, -f value  Frequency to ping (default: 0) [$PGPING_FREQUENCY_IN_MS]
   --query value, -c value      Query to run (default: "SELECT 1") [$PGPING_QUERY]
   --help                       
   --version, -v                print the version
```

## Example

```bash
$ pg-ping -U 'username' -p 'pwd' -d 'dbname' -h 'localhost:5432'
{"message":"dial tcp [::1]:5432: connect: connection refused","status":"failed","time_taken":"2.209ms","timestamp":"20:09:30"}
{"message":"5","status":"success","time_taken":"1082.709ms","timestamp":"20:09:31"}
{"message":"5","status":"success","time_taken":"84.863ms","timestamp":"20:09:32"}
{"message":"5","status":"success","time_taken":"0.891ms","timestamp":"20:09:33"}
{"message":"5","status":"success","time_taken":"0.893ms","timestamp":"20:09:34"}
{"message":"5","status":"success","time_taken":"1.185ms","timestamp":"20:09:35"}```
