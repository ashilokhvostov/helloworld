package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "JBOX is online now")
        files, _ := ioutil.ReadDir("/java")
            for _, f := range files {
                    fmt.Fprintln(w, f.Name())
            }
}

func main() {
        http.HandleFunc("/", helloHandler)

        fmt.Println("Started, serving at 80")
        err := http.ListenAndServe(":80", nil)
        if err != nil {
                panic("ListenAndServe: " + err.Error())
        }
}
