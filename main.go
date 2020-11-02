package main

import (
    "fmt"
    "os"
    "io"
    "log"
    "io/ioutil"
    "strings"
    "regexp"
)

func cat(r io.Reader) ([]byte, error) {
    data, err := ioutil.ReadAll(r)

    if (err != nil) {
        return nil, err
    }

    return data, nil
}

func main() {
    var output []byte
    var err error

    args := os.Args[1:]

    if len(args) > 0 {
    	concat := strings.Join(os.Args[1:], " ")

        data := strings.NewReader(concat)

        output, err = cat(data)

        if err != nil {
            log.Panic(err)
        }

        fmt.Print(regexp.QuoteMeta(string(output)))
    } else {
        output, err = cat(os.Stdin)

        if err != nil {
            log.Panic(err)
        }

        fmt.Print(regexp.QuoteMeta(string(output)))
    }
}