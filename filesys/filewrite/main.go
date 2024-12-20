package main

import (
	"fmt"

	"github.com/vladimirvivien/gexe"
)

func main() {
	gexe.FileWrite("./ca-config.json").String(`
{
	"signing": {
	  "default": { "expiry": "8760h" },
      "profiles": {"Kubernetes": { "usages": ["signing", "key encipherment", "server auth", "client auth"],"expiry": "8760h"}
	  }
	}
}`,
	)

	// format it with jq
	fmt.Println(gexe.Run(`jq '.[]' ca-config.json`))
}
