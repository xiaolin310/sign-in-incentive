package main

import (
    "fmt"
)


func main() {
    signInIndexList := make([]bool, 7)
    index := 3
    for i := 0; i <= index; i++ {
	signInIndexList[i] = true
    }
   fmt.Println(signInIndexList)
}

