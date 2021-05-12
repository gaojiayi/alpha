package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"encoding/json"
)

func modify(num int) {
	num = 1000
}
func modify_p(num *int) {
	*num = 1000
}

type s struct {
	aa string
	D  *d
}
type d struct {
	dd string
}

func modifyS(s *s) {
	s.aa = "word"
}
func main() {
	x := 1
	//&x 获取的是变量x的地址，并赋值给p,这个时候p就是一个指针
	//p是指针，所以*p获取的就是变量的值，指针指向的是变量x的值，即*p为1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	//这里*p 进行赋值，也就是更改了变量x的值，即实现不知道变量的名字更改变量的值
	*p = 2
	fmt.Println(x)

	a := 1
	modify(a)
	fmt.Println(a)

	modify_p(&a)
	fmt.Println(a)

	var s s
	modifyS(&s)
	fmt.Printf("----------")
	fmt.Println(s.aa)

	structD := d{
		"dddddddd",
	}

	s.D = &structD
	fmt.Println("----------")
	fmt.Println((*s.D).dd)
	fmt.Println("----等价于------")
	fmt.Println(s.D.dd)
	//out, err := exec.Command("pwd").Output()
	var outInfo bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/Users/gaojiayi/PycharmProjects/alpha-script/venv/bin/python",
		"/Users/gaojiayi/PycharmProjects/alpha-script/scripts/image_detect_executor.py","--imageName", "1.jpg" )
	cmd.Stdout = &outInfo
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error(), stderr.String())
	}
	// 执行
	//err = cmd.Wait()
	//out, err := cmd.Output()
	//out, err := exec.Command("python",
	//	"../../scripts/image_detect_executor.py --imageName 1.jpg").Output()
	//if err != nil {
	//	print(err)
	//}
	result := strings.Index(outInfo.String(),"data:")
	if result >= 0 {
		fmt.Println(string(outInfo.String()[result+5:]))
	}
	var mapArray  []map[string]string
	json.Unmarshal([]byte(string(outInfo.String()[result+5:])), &mapArray)
	fmt.Println(mapArray)
}
