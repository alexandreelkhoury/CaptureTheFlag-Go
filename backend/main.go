package main

import "fmt"

func main() {
	secretkey := callFirstAPI()
	fmt.Println("Secret key is:", secretkey)

	secondAPI(secretkey)

	thirdAPI()
}
