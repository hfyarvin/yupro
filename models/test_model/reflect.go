package test_model
import (
	"fmt"
	"reflect"
)
/*
	reflect.TypeOf
	reflect.ValueOf
*/

 func Ref() {
	 var x float64 = 3.4
	 v := reflect.ValueOf(x)
	 fmt.Println("type:", v.Type())
	 fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	 fmt.Println("value:", v.Float())
 }