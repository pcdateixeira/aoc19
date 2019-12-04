package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func readInts(fname string) (nums []int, err error) {
    b, err := ioutil.ReadFile(fname) // Reads the entire input file
    if err != nil { return nil, err }

    lines := strings.FieldsFunc(string(b), split) // Splits the numbers in each line of the file

    nums = make([]int, 0, len(lines)) // Creates a slice of capacity len(lines) but size 0, so it can be appended at every iteration

    for _, l := range lines {

        if len(l) == 0 { continue } // An empty line occurs at the end of the file afer being split

        n, err := strconv.Atoi(l)
        if err != nil { return nil, err }

        nums = append(nums, n)
    }

    return nums, nil
}

func split(r rune) bool {
    return r == '\n' || r == '\r'
}

func main() {

  mass, err := readInts("input.txt")
  if err != nil { panic(err) }

  var fuel = make([]int,len(mass))
  var totalFuel = 0
  var positiveFuel = true

  for positiveFuel == true { // As long as any fuel requires positive fuel to work
    positiveFuel = false

    for i, _ := range fuel {
      fuel[i] = (mass[i]/3) - 2

      if fuel[i] >= 0 {
        positiveFuel = true
        totalFuel += fuel[i]
      } else {
        fuel[i] = 0
      }
    }

    mass = fuel // Treats the list of fuel as the list of mass for the next iteration
  }

  fmt.Println("Total fuel required:", totalFuel)

}
