package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

// how to run:
// go mod init example/user/hello (to init project)
// go mod tidy (to add google dependency)
// go install example/user/hello OR go install (to build)
// export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .)) (add the install directory to our PATH to make running binaries easy)
// hello
func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
