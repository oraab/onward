package main

import (
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
    <link rel="stylesheet" href="fonts/stylesheet.css">
</head>
<body>
    <div style="font-family: rokkittregular; font-size: 30px; ">Task submitted.</div>
</body>
</html>
`)

}

func main() {
	http.HandleFunc("/create", createHandler)

	http.ListenAndServe(":8080",nil)
}
