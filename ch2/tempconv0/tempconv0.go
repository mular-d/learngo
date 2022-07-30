package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func main() {
	fmt.Printf("Boiling Point in Celsius : %s\n", BoilingC-FreezingC)
	fmt.Printf("Boiling Point in Fahrenheit : %s\n", CToF(BoilingC))
	fmt.Printf("Boiling Point in Kelvin : %s\n", CToK(BoilingC))
	// fmt.Printf("%g\n", boilingF - BoilingC) // Compile time error

	// c := FToC(212.0)
	// fmt.Println(c.String())
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%s\n", c)
	// fmt.Println(c)
	// fmt.Printf("%g\n", c)
	// fmt.Println(float64(c))
}
