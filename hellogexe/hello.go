package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// func main() {
// 	exe := gexe.New().SetVar("MYUSERNAME", "$USER")
// 	if exe.Eval("$MYUSERNAME") == "" {
// 		fmt.Println("You do not exist")
// 		os.Exit(1)
// 	}
// 	fmt.Printf("Hello %s !!! \n", exe.Eval("$MYUSERNAME"))
// }

func main() {
	if gexe.SetVar("MYUSERNAME", "$USER").Eval("$MYUSERNAME") == "" {
		fmt.Println("You do not exist")
		os.Exit(1)
	}
	fmt.Printf("Hello %s !!! \n", gexe.Eval("$MYUSERNAME"))
}