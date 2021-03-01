package lualang

import (
	"../classes"
	lua "github.com/yuin/gopher-lua"
)

const (
	LuaInstanceTypeName = "Instance"
)

type Instance struct {
	ClassName string
	ClassType interface{}
	ClassID   uint8
}

func getClassType(className string) interface{} {
	switch className {
	case "Player":
		{
			return &classes.Player{}
		}
	default:
		return nil
	}
}

// Constructor
func newInstance(L *lua.LState) int {
	i := &Instance{
		ClassName: L.CheckString(1),
		ClassID:   uint8(1),
	}
	i.ClassType = getClassType(i.ClassName)

	ud := L.NewUserData()
	ud.Value = i
	L.SetMetatable(ud, L.GetTypeMetatable(LuaInstanceTypeName))
	L.Push(ud)
	return 1
}

func checkData(L *lua.LState) interface{} {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Instance); ok {
		return v
	}
	return nil
}

func instanceGetSetID(L *lua.LState) int {
	p := checkData(L)

	if L.GetTop() == 2 {
		p.(*Instance).ClassID = uint8(L.CheckInt(2))
		return 0
	}
	L.Push(lua.LNumber(p.(*Instance).ClassID))
	return 1
}

var InstanceMethods = map[string]lua.LGFunction{
	"id": instanceGetSetID,
}
