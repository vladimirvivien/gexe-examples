package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// func main() {
// 	if gexe.ProgAvail("go") == ""{
// 		fmt.Printf("Go program not found")
// 		os.Exit(1)
// 	}

// 	// build matrix
// 	for _, arch := range []string{"arm64", "amd64"} {

// 		for _, opsys := range []string{"darwin", "linux"} {
// 			gexe.SetVar("arch", arch).SetVar("os", opsys)
// 			gexe.SetVar("binpath", fmt.Sprintf("build/%s/%s/bin", arch, opsys))

// 			result := gexe.Envs("CGO_ENABLED=0", "GOOS=$os", "GOARCH=$arch").Run("go build -o $binpath .")
// 			if result != "" {
// 				fmt.Printf("Build %s/%s failed: %s\n", arch, opsys, result)
// 				os.Exit(1)
// 			}

// 			fmt.Printf("Build %s/%s: %s OK\n", arch, opsys, gexe.Eval("$binpath"))
// 		}
		
// 	}
// }

func main() {
	if gexe.ProgAvail("go") == ""{
		fmt.Printf("Go program not found")
		os.Exit(1)
	}

	// build matrix
	for _, arch := range []string{"arm64", "amd64"} {

		for _, opsys := range []string{"darwin", "linux"} {
			gexe.SetVar("arch", arch).SetVar("os", opsys)
			gexe.SetVar("binpath", fmt.Sprintf("build/%s/%s/bin", arch, opsys))

			proc := gexe.Envs("CGO_ENABLED=0", "GOOS=$os", "GOARCH=$arch").RunProc("go build -o $binpath .")

			if proc.Err() != nil {
				fmt.Printf("Build %s/%s failed: %s\n", arch, opsys, proc.Result())
				os.Exit(1)
			}

			fmt.Printf("Build %s/%s: %s OK\n", arch, opsys, gexe.Eval("$binpath"))
		}
		
	}
}
