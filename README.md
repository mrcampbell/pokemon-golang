# Pokemon Golang

This is a portfolio project for my applications for Go jobs. It highlights some of my go-to's for building out projects, and follows some best practices.

## Installation and Setup

The project uses `Taskfiles` instead of `make`, as it was a personal preference that accomplishes a similar task. Commands can be copied from the `Taskfile.yml` and run in the terminal, or [you can follow the instructions to install `task` here](https://taskfile.dev/installation/)

This project uses `golang-migrate` for database schema migrations, and `sqlc` for query generation.

[Instructions to install the `golang-migrate` cli can be found here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)


### DB Migration Instructions

Ensure Docker Compose is running by starting it with 

```bash
docker compose up -d
```

(note the `-d` flag is for daemon mode, and the Postgres server is running on port `:6543` to avoid conflicts)

Migrations can be created using the following command:

```bash
migrate create -dir db/migrations -ext .sql <migration_name>
```

## Points of Interest

1. The whole project is loosely following the Ben Johnson project structure for Go.

