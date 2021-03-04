package lualang

import (
	"Go-Hill/classes"

	lua "github.com/yuin/gopher-lua"
)

func validateBrickType(L *lua.LState) interface{} {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Instance).ClassType.(*classes.Brick); ok {
		return v
	}
	return nil
}

func getSetColor(L *lua.LState) int {
	b := checkData(L)
	c := b.(*Instance).ClassType.(*classes.Brick).Color
	if L.GetTop() == 2 {
		c = L.CheckInt(2)
		return 0
	}
	L.Push(lua.LNumber(c))

	return 1
}

var BrickMethods = map[string]lua.LGFunction{
	"color": getSetColor,
}
