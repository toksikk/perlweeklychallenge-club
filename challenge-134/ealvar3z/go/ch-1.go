package main

import (
	"fmt"

	"github.com/ealvar3z/pwc"
)

/*
Reference: http://oeis.org/A050278

ten digit pandigitals from 0,9
1023456789,1023456798,1023456879,1023456897,
1023456978,1023456987,1023457689,1023457698,
1023457869,1023457896,1023457968,1023457986,
1023458679,1023458697,1023458769,1023458796,
1023458967,1023458976,1023459678,1023459687,
1023459768

zeroless pandigitals. i put them here for reference
123456789, 123456798, 123456879, 123456897, 123456978,
123456987, 123457689, 123457698, 123457869, 123457896,
123457968, 123457986, 123458679, 123458697, 123458769,
123458796, 123458967, 123458976, 123459678, 123459687,
123459768, 123459786, 123459867, 123459876, 123465789

*/

func main() {
	digits := []int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}
	// spin up a goroutine to calc the permutations
	panDigits := pwc.Permutator(digits)
	for i := 0; i < 5; i++ {
		fmt.Println(<-panDigits)
	}
	fmt.Println("done")
}