package main

import (
  "fmt"
  "os"
)

func main() {
  // access inputs as environment vars
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")
  
  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)
  
  // Sometimes inputs are not "required"
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
  }
}
