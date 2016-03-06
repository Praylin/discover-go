package main

import "fmt"
import "time"

func main() {
  fmt.Printf("%s\n", "Hello, we are Holberton School")
  fmt.Printf("%s", "the date is ")
  t := time.Now()
  fmt.Println(t);
  fmt.Println("the year is", t.Year())
  fmt.Println("the month is", t.Month())
  fmt.Println("the day is", t.Day())
}
