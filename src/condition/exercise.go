package condition

import "fmt"

//需要大写才能被看到
func Run()  {

	i := '2'

	if i == '3' {
		fmt.Println("条件通过")
	}else {
		fmt.Println("条件不成立")
	}
}