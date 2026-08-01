package basics

// basic types

// bool

// string

// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr

// byte - alias for uint8
// rune - alias for int32 - represent the Unicode code point

// float32 float64

// complex64 complex128

// The int, uint, and uintptr types are usually
// 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
// When you need an integer value you should use int
// unless you have a specific reason to use a sized or unsigned integer type.

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14

// Numeric constants
// These are high-precision values.
// An untyped constant takes the type needed by its context.
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)

	// Zero values:
	// Variables declared without an explicit initial value
	// are given their zero value.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversion:
	// The expression T(v) converts the value v to the type T.
	var in int = 42
	var fl float64 = float64(in)
	var ui uint = uint(fl)
	fmt.Println(in, fl, ui)

	// Type inference:
	// When declaring a variable without specifying an explicit type
	// (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.

	// But when the right hand side contains an untyped numeric constant,
	// the new variable may be an int, float64, or complex128
	// depending on the precision of the constant:

	v := 5 + 3i // change me!
	fmt.Printf("v is of type %T\n", v)

	// Constants:
	// Constants cannot be declared using the := syntax.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
