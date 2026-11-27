package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

// how to run:
// go install example/user/hello OR go install
// export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
// hello
func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
