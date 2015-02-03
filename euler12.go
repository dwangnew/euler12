package main

import (
  "fmt"
)

// returns number of divisors of a number num
func divisors(num int) (total int) {
  total = 0
  largest_possible_divisor := (num/2)+1
  for i := 1; i < largest_possible_divisor; i++ {
    // fmt.Println(i, num, num%i == 0)
    if num % i == 0 {
      total += 2
      largest_possible_divisor = num/i
    }
  }
  // fmt.Println("total of divisors is ", total)
  return total
}

// loops through all the triangle nums and checks the num of divisors
func triangle_nums(n int) (num int) {
  num = 1
  for i := 2; i < n; i++ {
    num += i
    // fmt.Println(i, num)
    if d := divisors(num); d > 500 {
      fmt.Println("number of divisors is: ", d)
      return num
      break
    }
  }
  return num
}

func main() {
  fmt.Println(triangle_nums(100000))
}