# Koanf Lua

Koanf Lua allows to write dynamic configuration files for [koanf](https://github.com/knadh/koanf) configuration manager. Sometimes dynamic configuration is required.

## Usage


```lua
-- config.lua
local os = require "os"

local config = {
  brokers = array({"localhost:42", "localhost:4242"}), -- Hack for Golang to detect the table as an array
  listen = "localhost",
  redis = {
    addr = "localhost:6379",
    password = "trololo",
    db = 1
  }
}

if os.getenv("MAGIC_FEATURE") == "enabled" then
  config["feature"] = "testouille"
end

luaconfig(config)
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/file"
	"github.com/mdouchement/koanflua"
)

func main() {
	konf := koanf.New(".")
	err := konf.Load(file.Provider("config.lua"), koanflua.Parser())
	if err != nil {
		log.Fatal(err)
	}

	//

	redisAddr := konf.String("redis.addr")
	fmt.Println(redisAddr)
}
```

## Resources

- Default Lua library (same as Lua's VM): https://github.com/Shopify/go-lua
- Extra Lua library: https://github.com/Shopify/goluago

## License

**MIT**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request
