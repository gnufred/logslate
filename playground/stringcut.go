package main

import (
    "fmt"
    "strings"
)

func main() {
    l := "208.115.199.20 - - [05/Jan/2023:07:07:54 +0000] \"HEAD /?monitoring=uptimerobot HTTP/1.1\" 200 0 \"https://moovadapt.fr/?monitoring=uptimerobot\" \"Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)\""

    // RemoteAddr
    i := strings.Index(l, "-")
    fmt.Println(i)
    c := l[0:i-1]
    fmt.Println(c)
    l = l[i+2:]
    fmt.Println(l)
    fmt.Println("")

    // RemoteUser
    i = strings.Index(l, "[")
    fmt.Println(i)
    c = l[0:i-1]
    fmt.Println(c)
    l = l[i+1:]
    fmt.Println(l)
    fmt.Println("")

    // TimeLocal
    i = strings.Index(l, "]")
    fmt.Println(i)
    c = l[0:i-1]
    fmt.Println(c)
    l = l[i+2:]
    fmt.Println(l)
    fmt.Println("")
}
