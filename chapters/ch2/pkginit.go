package ch2

import (
	"fmt"
	"os"
)

var PkgVarToInitForInternalOrExternalUse int

func init() {
	var err error
	PkgVarToInitForInternalOrExternalUse, err = funcMightReturnErr()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch2:init: chapters variable init job failed\n")
		os.Exit(1)
	}
	fmt.Printf(
		"ch2:init: chapters ch2 variable PkgVarToInitForInternalOrExternalUse init finished: %d\n",
		PkgVarToInitForInternalOrExternalUse)

}

func funcMightReturnErr() (int, error) {
	return 42, nil
}
