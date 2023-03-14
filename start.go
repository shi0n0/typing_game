package main

import (
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	for i := 3; i >= 1; i--{
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	score := 0
	for e := 1; e <= 3; e++ {
		var input string
		word := randomdata.Adjective()
		fmt.Printf("[問題%d] %s と入力してください\n",e,word)
		fmt.Scan(&input)
		if input == word{
			fmt.Println("正解！")
			score++
		} else {
			fmt.Println("不正解...")
			fmt.Printf("あなたが入力したのは%sです\n",input)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Printf("あなたのスコアは%dでした\n",score)

}
