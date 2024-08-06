package example

import "fmt"

func Hello(msg string, repeate int) {
        for repeate > 0 {
                fmt.Println("lib says: " + msg)
                repeate-=1
        }
}
