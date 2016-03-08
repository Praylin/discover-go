package main

import (
	"fmt"
	//"strconv"
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

func (u user) PrintInfo() int  {
	t := time.Now()
	//fmt.Println(t.Year())
	value := "03/07/1917"
	layout := "01/02/2006"
	d, _ := time.Parse(layout, value)
	//fmt.Println(d.Year())
	age := (t.Year() - d.Year())
	//fmt.Println(age)
	fmt.Println("Betty Holberton who was born in", u.City, "would be", age, "years old today.")
	return age
}

func main() {
	u := user{"Betty Holberton", "1917-Mar-07", "Philadelphia"}
	u.PrintName()
	u.PrintInfo()
}
