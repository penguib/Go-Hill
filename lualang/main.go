package lualang

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

const (
	coreScriptsPath = "./lualang/corescripts/"
	testScriptsPath = "./lualang/tests/"
)

// L is the main Lua state that the server will be using
var L *lua.LState = lua.NewState()

// Init initializes all the Lua globals
// Should be called when the server starts
func Init() {

	// Instances
	mt := L.NewTypeMetatable(LuaInstanceTypeName)
	L.SetGlobal("Instance", mt)
	L.SetField(mt, "new", L.NewFunction(newInstance))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), getInstanceMethods()))

	// Try to define all Instance methods
	// First check the certain Instance
	// Then only allow the methods that pertain to that class

	err := L.DoFile(testScriptsPath + "script.lua")

	if err != nil {
		panic(err)
	}

	fmt.Println("Lua successfully loaded")
}
