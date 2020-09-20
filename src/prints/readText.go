package prints

import (
	"bufio"
	"fmt"
	"os"
)

func Prints() {
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)
	s1, s2 := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1, s2)

}
