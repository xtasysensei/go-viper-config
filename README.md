## This is a simple package that helps load envronmental variables from a toml file using viper


## usage
```go
import (
	"fmt"
	"log"

	loadcfg "github.com/xtasysensei/go-viper-config/loadCfg"
)

func main() {
	config, err := loadcfg.LoadTomlCfg()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Postgres Host: %s\n", config.Db.Psql.Host)
	fmt.Printf("Postgres Port: %s\n", config.Db.Psql.Port)
	fmt.Printf("Server Port: %s\n", config.SRV.Port)
}

```

And a toml file should be present in the root directory or you can modify the `loadcfg.go` to specify another directory

```toml
[database]
    [database.postgresql]
        hostname = "localhost"
        port = 5432
        username = "postgres"
        password = "1234"

[server]
    port = 8080
```

## License
Use as you please
