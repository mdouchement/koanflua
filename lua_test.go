package koanflua_test

import (
	"os"
	"testing"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/mdouchement/koanflua"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cfg := []byte(`
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
		
		luaconfig(config)`)

	//

	os.Setenv("MAGIC_FEATURE", "enabled")

	konf := koanf.New(".")
	err := konf.Load(rawbytes.Provider(cfg), koanflua.Parser())
	assert.NoError(t, err)

	assert.Equal(t, []string{"localhost:42", "localhost:4242"}, konf.Strings("brokers"))
	assert.Equal(t, "localhost", konf.String("listen"))

	assert.Equal(t, "localhost:6379", konf.String("redis.addr"))
	assert.Equal(t, "trololo", konf.String("redis.password"))
	assert.Equal(t, 1, konf.Int("redis.db"))

	assert.Equal(t, "testouille", konf.String("feature"))
}
