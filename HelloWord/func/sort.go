package main

import (
	"fmt"
	"sort"
)

type U struct {
	ID    int
	Name  string
	Score int
}

func main() {
	list := [][2]int{{1, 3}, {5, 1}, {5, 7}, {7, 1}, {5, 8}}
	//对切片进行排序，使用数组的第二个（索引为1）元素比较大小进行排序

	sort.Slice(list, func(i, j int) bool {
		//前面小的往前移
		return list[i][1] < list[j][1]
	})
	fmt.Println(list)

	Us := []U{{1, "kk", 5}, {2, "gg", 6}, {3, "hh", 10}, {4, "烟灰", 16}}
	sort.Slice(Us, func(i, j int) bool {
		return Us[i].Score < Us[j].Score
	})
	fmt.Println(Us)
}

/*作业
1.复习整理课程内容
2.用户管理
  用户改成struct，并且针对属性使用不同的数据类型
  ID int
  Name string
  birthday time.Time
  tel string
  addr string
  desc string
  注意需要进行类型转换

3.在添加和修改数据是对用户进行检查，用户名不能为空，且用户名必须唯一
a.添加kk
b.再添加kk 不能添加（用户名已经存在）
c.有kk
d.修改一个非kk 不同添加（用户名已经存在）
e.修改kk = >kk(修改成功)
4.在查询时，让用户输入排序属性，按照属性值从小到大顺序显示









*/
