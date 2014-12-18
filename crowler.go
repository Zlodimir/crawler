package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    response, err := http.Get("http://bash.im/")
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
        }
        fmt.Printf("%s\n", string(contents))
    }
}