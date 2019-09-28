package main

import "fmt"


func main() {
	msas := MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5})
	fmt.Printf("msas = %v\n", msas)
}


func MaximumSubarraySum(numbers []int) int {
	max := 0
	  
	  for x := len(numbers); x > 0; x-- {				// # of digits to pull
		  for y := 0; (y + x) <= len(numbers); y++ {
			  sum := 0
			  for z := y; z < y+x; z++ { sum += numbers[z] }
			  if (sum > max) { max = sum;}
		  }
	  }
  
	return max
  }

