package main

import "fmt"

type Number interface{
	int64 | float64
}

func SumInts(m map[string]int64) int64{
	
	var s int64
	for _, v:= range m{
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64{
	var s float64
	for _, v:= range m{
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V{

	var s V
	for _, v:= range m{
		s += v
	}
	return s

}

func main(){
	ints := map[string]int64{
		"first"	: 20,
		"second": 30,
	}
	floats := map[string]float64{
		"first"  : 20.5,
		"second" : 30.2,
	}
	fmt.Println("Sumints: ", SumInts(ints), " Sumfloats ", SumFloats(floats))
	fmt.Println("SumIntsOrFloats ", SumIntsOrFloats(ints))
}