package koanflua

import (
	"errors"

	"github.com/Shopify/go-lua"
	"github.com/Shopify/goluago"
	"github.com/Shopify/goluago/util"
)

// LUA implements a LUA parser for koanf.
type LUA struct {
	state *lua.State
}

// Parser returns a LUA Parser.
func Parser() *LUA {
	state := lua.NewState()
	lua.OpenLibraries(state)
	goluago.Open(state)

	return &LUA{
		state: state,
	}
}

// Parse parses the given LUA bytes.
func (p *LUA) Parse(b []byte) (map[string]interface{}, error) {
	var (
		raw interface{}
		err error
		ok  bool
		cfg map[string]interface{}
	)

	//

	p.state.Register("luaconfig", func(l *lua.State) int {
		raw, err = util.PullTable(l, 1)
		return 0
	})

	//

	if errl := lua.DoString(p.state, string(b)); errl != nil {
		return nil, errl
	}

	if err != nil {
		return nil, err
	}

	cfg, ok = raw.(map[string]interface{})
	if !ok {
		return nil, errors.New("config is not a table with only string keys at first level")
	}

	return cfg, nil
}
