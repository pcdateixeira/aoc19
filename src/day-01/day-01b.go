package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func readInts(fname string) (nums []int, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.FieldsFunc(string(b), split)
    // Assign cap to avoid resize on every append.
    nums = make([]int, 0, len(lines))

    for _, l := range lines {
        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }
        // Atoi better suits the job when we know exactly what we're dealing
        // with. Scanf is the more general option.
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
      } else {
        fuel[i] = 0
      }
    }

    for _, v := range fuel {
      totalFuel = totalFuel + v
    }

    mass = fuel // Treats the list of fuel as the list of mass for the next iteration
  }

  fmt.Println("Total fuel required:", totalFuel)

}
