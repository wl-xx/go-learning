package main

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// 变量声明
	var num int = 10
	var str = "test"
	arr := []int{1, 2, 3}

	fmt.Printf("num: %d, str: %s\n", num, str)

	// 遍历
	for index, value := range arr {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}

	// 数组操作
	fmt.Println("---------------array-----------------")
	arr[0] = 3
	fmt.Println("数组长度", len(arr))
	arr = append(arr, 4, 5)
	fmt.Println("数组长度", len(arr))
	fmt.Println("查找5的下标", slices.Index(arr, 5))
	fmt.Println(slices.Contains(arr, 3))
	fmt.Println(arr)
	slices.Reverse(arr)
	fmt.Println(arr)
	slices.Sort(arr)
	fmt.Println(arr)
	slices.SortFunc(arr, func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})
	fmt.Println(arr)
	slices.SortFunc(arr, func(a, b int) int {
		return cmp.Compare(a, b)
	})
	fmt.Println(arr)
	fmt.Println(slices.Clone(arr))
	fmt.Println(arr[2:4])
	fmt.Println(arr[:0])
	for chunk := range slices.Chunk(arr, 2) {
		fmt.Println(chunk)
	}
	fmt.Println(slices.Collect(slices.Chunk(arr, 2)))

	// map
	fmt.Println("---------------map-----------------")
	person := map[string]string{
		"name":  "Tom",
		"hobby": "game",
	}
	fmt.Println(person)
	person["gender"] = "male"
	for key, value := range person {
		fmt.Println(key, value)
	}
	name, exist := person["name"]
	fmt.Println(name, exist)
	delete(person, "hobby")
	fmt.Println(person)
	person2 := maps.Clone(person)
	fmt.Println(person2)
	clear(person2)
	fmt.Println(person2)

	// string
	fmt.Println("---------------string-----------------")
	s := "你好Go"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(strings.Split(s, ""))
	fmt.Println(strings.Join(strings.Split(s, ""), "-"))
	for i, c := range s {
		fmt.Println(i, string(c))
	}
	fmt.Println(strings.TrimSpace("  hello  "))
	fmt.Println(strings.TrimLeft("  hello  ", " "))
	fmt.Println(strings.TrimPrefix("abcabcHello", "abc"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strconv.Itoa(32))
	n, e := strconv.Atoi("45")
	fmt.Println(n, e, e == nil)
	s2 := fmt.Sprintf("name: %s, scores: %d", "Jerry", 89)
	fmt.Println(s2)

	// function
	fmt.Println("---------------string-----------------")
	defer fmt.Println("最后打印")
	fmt.Println(Add(1, 2, 3))
	v, err := Divide(3, 4)
	fmt.Println(v, err)
	next := Counter()
	fmt.Println(next(), next())

	// time
	fmt.Println("---------------time-----------------")
	now := time.Now()
	fmt.Println(int(now.Month()))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(time.DateTime)
	nextTime := now.Add(24 * time.Hour)
	fmt.Println(nextTime.Sub(now))
	t, err := time.Parse("2006-01-02", "2026-10-08")
	fmt.Println(t, err)

	// file
	os.WriteFile("data.json", []byte("{\"name\": \"张三\"}"), 0o655)
	// ReadFile会读取到内存中，不适合大文件读取
	data, err := os.ReadFile("data.json")
	fmt.Println(string(data), err)
}
