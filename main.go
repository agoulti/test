package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	str := `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a test message!");
}
`

	if filename == "" {
		fmt.Println(str)
	} else {
		err := ioutil.WriteFile(filename, []byte(str), 0644)
		if err != nil {
			panic(err)
		}
	}
}
