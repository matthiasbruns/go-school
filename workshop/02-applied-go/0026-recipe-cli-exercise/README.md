This project has two entry points: a CLI and a server.
Both offer the same functionality: searching for recipes by name but offer different interfaces.

## Starting the CLI

```bash
go run cmd/cli/main.go search bananabread
```

### Commands

- `search <name>`: Search for recipes by name
- `id <id>`: Get a recipe by id
- `random`: Get a random recipe

All commands also have an optional flat `-save`, which saves the recipe to a file.
> The file is saved in the current directory and is named after the recipe's id.

## Starting the Server

```bash 
go run cmd/server/main.go
```

### Endpoints

#### `GET /search?name=bananabread`

```bash
curl http://localhost:8090/search?name=bananabread
```

#### `GET /recipes/{id}`

```bash
curl http://localhost:8090/recipes/1
```

#### `GET /random`

```bash
curl http://localhost:8090/random
```


## Hints

This time, there are no hints. You should be able to complete this exercise on your own.
