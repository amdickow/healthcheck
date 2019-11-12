package main

import (
        "flag"
        "net/http"
        "os"
)

func main() {
        port := flag.String("port", "80", "port on localhost to check")
        host := flag.String("host", "127.0.0.1", "host to run the check against")
        flag.Parse()

        resp, err := http.Get("http://" + *host + ":" + *port + "/health")
        if err != nil || resp.StatusCode != 200 {
                os.Exit(1)
        }
        os.Exit(0)
}
