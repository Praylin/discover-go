package main

  import (
    "fmt"
    "math/rand"
  )

  func main() {
    randomNumber := rand.Intn(100)
    if randomNumber > 50 {
      fmt.Println("my random number is" , randomNumber, "and is greater than 50")
    } else {
      fmt.Println("my random number is" , randomNumber, "and is lesser than 50")
    }
    school := "Holberton School"
    if school == "Holberton School" {
      fmt.Println("I am a student of the", school)
    } else {
      fmt.Println("I am not a student of the", school)
    }
    beautifulWeather := true
    if beautifulWeather == true {
      fmt.Println("It's a beautiful weather!")
    }
    holbertonFounders := [3]string{"Rudy Rigot","Sylvain Kalache","Julien Barbier"}
    for i := 0; i < 3; i++ {
      fmt.Println(holbertonFounders[i],"is a founder")
    }
  }
