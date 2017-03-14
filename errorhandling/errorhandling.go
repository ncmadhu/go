package main

import "fmt"

type negativeSqrt struct {
	e float64
}

func (n *negativeSqrt) Error() string {
    return fmt.Sprintf("%v - %s", n.e, "is a negative number")
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//n := fmt.Sprint(float64(e))

	return "nil" 
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, &negativeSqrt{x}
		//return x, ErrNegativeSqrt(x).Error()
	} else {
		var sqrt float64
		for z := 1; z <= 2; z++ {
			z := float64(z)
			sqrt := z - (z * z - x) / 2 * z
			fmt.Println(sqrt)
		}
		return sqrt, nil
	}

}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


