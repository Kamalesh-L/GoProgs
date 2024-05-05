package main
import (
    "fmt"
)
func main() {
    var arr=[]int{21,22,23,24,25,26,27,28,29,30}
    chOdd:=make(chan int)
    chEven:=make(chan int)
    go Odd(chOdd)
    go Even(chEven)
    for _,i:=range arr {
        if i%2==0 {
            chEven <- i
        } else {
            chOdd <- i
        }
    }
}
func Odd(ch <- chan int) {
    for i:= range ch {
        fmt.Println("Odd: ",i,"hlelo")
    }
}
func Even(ch <- chan int) {
    for i:= range ch {
        fmt.Println("Even: ",i)
    }
}
