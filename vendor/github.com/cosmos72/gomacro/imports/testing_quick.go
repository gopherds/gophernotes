// this file was generated by gomacro command: import _b "testing/quick"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"math/rand"
	"reflect"
	"testing/quick"
)

// reflection: allow interpreted code to import "testing/quick"
func init() {
	Packages["testing/quick"] = Package{
	Binds: map[string]Value{
		"Check":	ValueOf(quick.Check),
		"CheckEqual":	ValueOf(quick.CheckEqual),
		"Value":	ValueOf(quick.Value),
	}, Types: map[string]Type{
		"CheckEqualError":	TypeOf((*quick.CheckEqualError)(nil)).Elem(),
		"CheckError":	TypeOf((*quick.CheckError)(nil)).Elem(),
		"Config":	TypeOf((*quick.Config)(nil)).Elem(),
		"Generator":	TypeOf((*quick.Generator)(nil)).Elem(),
		"SetupError":	TypeOf((*quick.SetupError)(nil)).Elem(),
	}, Proxies: map[string]Type{
		"Generator":	TypeOf((*P_testing_quick_Generator)(nil)).Elem(),
	}, 
	}
}

// --------------- proxy for testing/quick.Generator ---------------
type P_testing_quick_Generator struct {
	Object	interface{}
	Generate_	func(_proxy_obj_ interface{}, rand *rand.Rand, size int) reflect.Value
}
func (P *P_testing_quick_Generator) Generate(rand *rand.Rand, size int) reflect.Value {
	return P.Generate_(P.Object, rand, size)
}
