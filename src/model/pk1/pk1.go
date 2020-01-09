package pk1

import (
	"fmt"
	"model/pk2"
)

func ExexPk1() {
	pk2.Execpk2()
	fmt.Println("ExecPk1")
}

func init()  {
	fmt.Println("pk1 init")
}
