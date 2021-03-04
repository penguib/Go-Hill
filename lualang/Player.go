package lualang

import (
	"Go-Hill/classes"

	lua "github.com/yuin/gopher-lua"
)

func getSetName(L *lua.LState) int {
	b := checkData(L)
	n := b.(*Instance).ClassType.(*classes.Player).Username

	if L.GetTop() == 2 {
		n = L.CheckString(2)
		return 0
	}

	L.Push(lua.LString(n))
	return 1
}

var PlayerMethods = map[string]lua.LGFunction{
	"username": getSetName,
}
