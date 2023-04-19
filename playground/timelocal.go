package main

import (
    "fmt"
    "time"
    "reflect"
)

func main() {
    fmt.Println("time to roll!")
    timeStr := "05/Jan/2023:07:07:54 +0000"
    timePrs, _ := time.Parse("02/Jan/2006:15:04:05 -0700", timeStr)
    fmt.Println(reflect.TypeOf(timePrs))
    fmt.Println(timePrs)
    fmt.Println(timePrs.Unix())
}
