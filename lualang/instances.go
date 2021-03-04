package lualang

import (
	"Go-Hill/classes"

	lua "github.com/yuin/gopher-lua"
)

const (
	LuaInstanceTypeName = "Instance"
)

var instanceClassTypeID uint8

type Instance struct {
	ClassName   string
	ClassType   interface{}
	ClassTypeID uint8
	ClassID     uint8
}

func getClassType(className string) (interface{}, uint8) {
	switch className {
	case "Player":
		{
			return &classes.Player{
				Username: "Player",
			}, 1
		}
	case "Brick":
		{
			return &classes.Brick{
				Color: 0xff00ff,
				Name:  "Brick",
			}, 2
		}
	default:
		return nil, 0
	}
}

// Constructor
func newInstance(L *lua.LState) int {
	i := &Instance{
		ClassName: L.CheckString(1),
		ClassID:   uint8(1),
	}
	cType, cID := getClassType(i.ClassName)
	i.ClassType = cType
	i.ClassID = cID
	instanceClassTypeID = cID

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

func getInstanceMethods() map[string]lua.LGFunction {
	switch instanceClassTypeID {
	case 1:
		return PlayerMethods
	case 2:
		return BrickMethods
	default:
		return PlayerMethods
	}

}

var InstanceMethods = map[string]lua.LGFunction{
	"id": instanceGetSetID,
}
