// func main() {
// err := cmd.Execute()
// if err != nil {
// log.Fatalf("cmd.Execute err: %v", err)
// }
// }
package main

import (
	"log"
	"utils-go/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Excute err:%v", err)
	}
}
