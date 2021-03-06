package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type StringElement struct {
	key   string
	value int
}

func parse(line string) (string, string, int, error) {
	//空白符分割
	elements := strings.Fields(line)
	if len(elements) < 8 {
		return "", "", 0, fmt.Errorf("format error")
	}
	//字符串转换成int
	statusCode, err := strconv.Atoi(elements[8])
	if err != nil {
		return "", "", 0, err
	}
	return elements[0], elements[6], statusCode, nil
}

func main() {
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	//如果成功延迟关闭
	defer file.Close()
	//每次读取文件一行
	reader := bufio.NewReader(file)

	ipStats := make(map[string]int)
	urlStats := make(map[string]int)
	//状态码
	statusCodeStats := make(map[int]int)
	for {
		ctx, _, err := reader.ReadLine()
		fmt.Println(err, string(ctx))
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if ip, url, statusCode, err := parse(string(ctx)); err == nil {
			ipStats[ip]++
			urlStats[url]++
			statusCodeStats[statusCode]++
		}
	}
	fmt.Println(ipStats)
	fmt.Println(statusCodeStats)

	ipStatsSlice := make([]StringElement, 0, len(ipStats))
	for k, v := range ipStats {
		ipStatsSlice = append(ipStatsSlice, StringElement{k, v})
	}
	fmt.Println(ipStatsSlice)
	//冒泡
	sort.Slice(ipStatsSlice, func(i, j int) bool {
		count1 := ipStatsSlice[i].value
		count2 := ipStatsSlice[j].value
		return count1 < count2

	})

	topn := 10
	for i := len(ipStatsSlice) - 1; i >= 0 && i >= len(ipStatsSlice)-1-topn; i-- {
		fmt.Println(ipStatsSlice[i])

	}
}

//2:22:46
