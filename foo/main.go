package main

import "github.com/mariuszs/url62-go/url62"

func main() {
	id, _ := url62.FromUUID("c3587ec5-0976-497f-8374-61e0c2ea3da5")

	println(id)

	uuid, _ := url62.ToUUID("5wbwf6yUxVBcr48AMbz9cb")

	println(uuid)

}
