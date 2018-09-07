// go run main.go > index.html : create html file in current directory
package main

import "fmt"

func main() {
	name := "Jack Sparrow"

	concat := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>`

	fmt.Println(concat)
}
