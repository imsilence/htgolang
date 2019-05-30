package main

import (
	"fmt"
	"sort"
)

func main() {
	//对int类型切片排序
	nums := []int{3, 2, 9, 8, 6, 7, 1}
	sort.Ints(nums)
	fmt.Println(nums)

	// 查找int类型切片中元素插入的位置
	fmt.Println(sort.SearchInts(nums, 5))
	fmt.Println(sort.SearchInts(nums, 6))

	// 定义结构体切片
	users := []struct {
		id   int
		name string
	}{
		{3, "kk"},
		{2, "silence"},
		{1, "woniu"},
		{6, "ada"},
		{5, "wd"},
	}

	fmt.Println(users)

	// 对结构体切片排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].id < users[j].id
	})

	fmt.Println(users)

	// 查找id 是否存在于切片中
	for _, value := range([]int{4, 5}) {

		idx := sort.Search(len(users), func(i int) bool {
			return users[i].id >= value
		})
		if idx < len(users) && users[idx].id == value {
			fmt.Printf("find %d index: %d\n", value, idx)
		} else {
			fmt.Printf("not found %d\n", value)
		}
	}



}
