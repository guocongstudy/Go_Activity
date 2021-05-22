package main

import "fmt"

func main() {
	var scores map[string]int //nil映射

	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)

	//字面量
	//scores =map[string]int{}
	scores = map[string]int{"小郭": 90, "小李": 80, "小刘": 70}
	fmt.Println(scores)
	//增删改查
	//key
	fmt.Println(scores["小郭"])
	v, ok := scores["宝成"] //验证存不存在，不存在v返回0值，
	if ok {
		fmt.Println(v)
	}
	fmt.Println(ok)
	scores["烟灰"] = 8
	fmt.Println(scores)
	scores["烟灰"] = 9
	fmt.Println(scores)

	scores["保定"] = 6
	fmt.Println(scores)

	//删除保定
	delete(scores, "保定")
	fmt.Println(scores)

	//fmt.Printf 说白了就是需要%d，%v占位时这种才会用到。
	fmt.Println(len(scores))
	for k, v := range scores {
		fmt.Println(k, v)
	}
	//scores = map[string]int{"小郭": 90, "小李": 80, "小刘": 70}
	//范例：映射[字符串]字符串{"地方","联系方式","成绩"}
	//错误写法：name =map[string]string{"小郭":"汉中","电话","成绩"}
    //正确示范
    var users map[string]map[string]string
	users = map[string]map[string]string{"小郭":{"地方":"陕西","成绩":"49","联系方式":"222323"}}
	fmt.Println(users)
	fmt.Printf("%T,%#v\n",users,users)

	//_,ok :=users["吴鹏飞"]
	//fmt.Println(ok)
	//users["吴鹏飞"] = map[string]string{"地方":"北京"}
	//fmt.Println(users)
	//users["吴鹏飞"]["成绩"] = "9"
	//fmt.Println(users)
	//delete(users["烟灰"],"联系方式")
	//fmt.Println(users)

}
