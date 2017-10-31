package main

import (
    "fmt"
    "io"
    "net/http"
)

func serveMe(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/html")

    io.WriteString(
        res,
        `<DOCTYPE html>
        <html>
            <head>
                <title>Hello World</title>
            </head>
            <body>
                <h1>Hello World!</h1>
            </body>
        </html>
        `,
    )
}

func main() {
    http.HandleFunc("/", serveMe)
    fmt.Println("Go to http://localhost:9000/")
    http.ListenAndServe(":9000", nil)
}
