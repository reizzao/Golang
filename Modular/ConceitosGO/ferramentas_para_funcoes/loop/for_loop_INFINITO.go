package loop

import (
	"fmt"
	"time"
)

func Loop_Infinito() string {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("oi")
	}
}
