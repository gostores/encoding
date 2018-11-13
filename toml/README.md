# Toml

## Features

Toml provides the following features for using data parsed from TOML documents:

* Load TOML documents from files and string data
* Easily navigate TOML structure using Tree
* Mashaling and unmarshaling to and from data structures
* Line & column position data for all parsed elements
* [Query support similar to JSON-Path](query/)
* Syntax errors contain line and column numbers

## Import

```go
import "github.com/govenue/encoding/toml"
```

## Usage example

Read a TOML document:

```go
config, _ := toml.Load(`
[postgres]
user = "pelletier"
password = "mypassword"`)
// retrieve data directly
user := config.Get("postgres.user").(string)

// or using an intermediate object
postgresConfig := config.Get("postgres").(*toml.Tree)
password := postgresConfig.Get("password").(string)
```

Or use Unmarshal:

```go
type Postgres struct {
    User     string
    Password string
}
type Config struct {
    Postgres Postgres
}

doc := []byte(`
[Postgres]
User = "pelletier"
Password = "mypassword"`)

config := Config{}
toml.Unmarshal(doc, &config)
fmt.Println("user=", config.Postgres.User)
```

Or use a query:

```go
// use a query to gather elements without walking the tree
q, _ := query.Compile("$..[user,password]")
results := q.Execute(config)
for ii, item := range results.Values() {
    fmt.Println("Query result %d: %v", ii, item)
}
```

## Documentation

## Tools

* `tomll`: Reads TOML files and lint them.

    ```
    go install github.com/govenue/encoding/toml/cmd/tomll
    tomll --help
    ```
* `tomljson`: Reads a TOML file and outputs its JSON representation.

    ```
    go install github.com/govenue/encoding/toml/cmd/tomljson
    tomljson --help
    ```

## Contribute

Feel free to report bugs and patches using GitHub's pull requests system on
[pelletier/go-toml](https://github.com/pelletier/go-toml). Any feedback would be
much appreciated!

## License

The MIT License (MIT). Read [LICENSE](LICENSE).
