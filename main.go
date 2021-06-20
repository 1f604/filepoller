package main

import (
        "fmt"

        "log"
        "os"

        "time"
)

func main() {

        log.SetFlags(log.LstdFlags | log.Lmicroseconds)
        const N = 1000
        ticker := time.NewTicker(time.Second / N)
        var lastmodtime int64 = 0
        for range ticker.C {
                fileinfo, err := os.Stat("/tmp/foo")
                if err != nil {
                        panic(err)
                }
                modtime := fileinfo.ModTime().UnixNano()
                if modtime != lastmodtime {
                        lastmodtime = modtime
                        fmt.Println(modtime)
                }
        }

}
