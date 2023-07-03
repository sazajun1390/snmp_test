package snmptest

import (
	"fmt"

	gsnm "github.com/gosnmp/gosnmp"
)

func main() {
	gsnm.Default.Target = "192.168.1.10"
	fmt.Println("Hello World")
}
