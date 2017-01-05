# sysutils
This package implements some useful system utility functions in the way of which Delphi(2007) RTL has done.</br>
At present, it is still under development.</br></br>

HOW TO USE:</br>
package main

import (
	"fmt"
	
	"github.com/ecofast/sysutils"
)

func main() {
	fmt.Println(sysutils.IncludeTrailingBackslash(`E:\admin\Software\`))
	fmt.Println(sysutils.GetApplicationPath())
}
