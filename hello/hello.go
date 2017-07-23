package main

import (
  "C"
  "log"
)

//export Hello
func Hello(name string) string {
  return "Hello World! " + name
}

func init() {
  log.Println("Loaded!!")
}

func main() {
}
