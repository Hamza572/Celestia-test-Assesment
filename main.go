package main

import ("fmt")

func evenOdd(n int) string{
    if(n%2==0) {
        return "Even"
    } else {
        return "Odd"
    }
}

func main() {
    fmt.Println(evenOdd(2))
}