package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("3秒後に開始します")
	time.Sleep(2 * time.Second)

	for i := 3; i >= 1; i--{
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	
	var input string
	fmt.Println("[ golang ] と入力してください")
	fmt.Scan(&input)
	if input == "golang"{
		fmt.Println("正解！")
	} else {
		fmt.Println("不正解...")
		fmt.Printf("あなたが入力したのは%sです\n",input)
		time.Sleep(1 * time.Second)
	}
}