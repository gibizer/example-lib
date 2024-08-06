package example

import "fmt"

func Hello(msg string, repeate int) {
        if repeate < 0 {
                panic("repate cannot be negative")
        }
        for repeate > 0 {
                fmt.Println("lib says: " + msg)
                repeate-=1
        }
}
