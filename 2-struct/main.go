package main

import (
	"fmt"
	"strconv"
	"time"
)

type user struct {
	Name string
	DOB  string
	City string
}

func (u user) PrintName() string {
	fmt.Println("Hello", u.Name)
	return u.Name
}

func (u user) PrintInfo() string {
	year := u.DOB[9:len(u.DOB)]
	i, err := strconv.Atoi(year)
	fmt.Println(i)
	fmt.Println(err)
	t := time.Now()
  fmt.Sprintf(t)
	Curyear := t[0:4]
	i1, err1 := strconv.Atoi(Curyear)
	age := i1 - i
	//fmt.Println(t1)
	//age := t - i
	fmt.Println(age)
	//fmt.Println(u.Name, "who was born in", u.City, "would be" )
	return u.Name
}

func main() {
	u := user{"Betty Holberton", "March 7, 1917", "Philadelphia"}
	u.PrintName()
	u.PrintInfo()
}
