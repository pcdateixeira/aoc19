package main

import (
  "fmt"
  "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

  mass := [5]int{10,20,30,40,50}
  var module [5]int
  var totalFuel = 0


  for i := 0; i < 5; i++ {
    module[i] = (mass[i]/3) - 2
    fmt.Println(module[i])
  }

  for i := 0; i < len(module); i++ {
		totalFuel = totalFuel + module[i]
	}

  fmt.Println(totalFuel)

  dat, err := ioutil.ReadFile("input.txt")
    check(err)
    fmt.Print(dat[0])

}
