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

  var totalFuel = 0

  for i, _ := range mass {
    totalFuel += (mass[i]/3) - 2
  }

  fmt.Println("Total fuel required:", totalFuel)

}
