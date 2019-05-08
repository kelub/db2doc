package handle

import (
	"fmt"
)

func Main(mapArgs map[string]*string) {
	for k, v := range mapArgs {
		fmt.Printf("%s=%s\n", k, *v)
	}
	//NewMyEngine(mapArgs)
}
