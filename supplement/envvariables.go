package supplement

import (
	"fmt"
	"os"
	"strings"
)

//设置当前进程的环境变量

func EnvVariables() {

	for _, e := range os.Environ() {
		fmt.Printf("key value: %#v\n", strings.SplitN(e, "=", 2))
	}

	fmt.Printf("%s\n", os.Getenv("Path"))
	os.Setenv("Path", "new path value")
	fmt.Printf("%s\n", os.Getenv("Path"))
	os.Getenv("Path")
	fmt.Printf("%s\n", os.Getenv("Path"))

}
