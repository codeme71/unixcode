package main

import (
 "fmt"
 "time"
)

func main() {
 now := time.Now()
 unixSeconds := now.Unix()
 fmt.Println(" Число секунд, прошедших с полуночи 1 января 1970 года:", unixSeconds)
}
