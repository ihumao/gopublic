package main

import (
	".."
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(public.StrToUInt64(`123`))
}
