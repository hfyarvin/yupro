package test_controller

import (
	"path/filepath"
	"os"
	"os/exec"
	"fmt"
	"strings"

	// "ybpro/models/test_model"
)

func (self *TestController) Test() {
	self.Data["json"] = ""
	self.ServeJSON()
}

func (self *TestController) GetTestNumber() {
	_id := self.Ctx.Input.Param(":id")
	ret := fmt.Sprintf("the number of test: %s", _id)
	self.Data["json"] = ret
	self.ServeJSON()
}

func t() string {
	m := make(map[string]string)
	m["babe"] = "baba"

	return m["babe"]
}

// 函数作为类型传递
type testInt func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
func fliter(slice []int, f testInt) []int {
	var ret []int
	for _, v := range slice {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func testIntFunc() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := fliter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := fliter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}

// ------------------------------------------------------------------------------------
// Comma-ok断言，确定变量类型value, ok = element.(T)/value := element.(type)
// ------------------------------------------------------------------------------------

// ------------------------------------------------------------------------------------
// http
// ------------------------------------------------------------------------------------

func (self *TestController) GetRequestInfo() {
	r := self.Ctx.Request
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	self.Data["json"] = r.Form
	self.ServeJSON()
}

// 访问路径
func (self *TestController) GetPath() {
	type Ret struct {
		Path string `json:"path"`
	}
	ret := new(Ret)
	ret.Path = GetAPPRootPath()
	self.Data["json"] = ret
	self.ServeJSON()
}

func GetAPPRootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return ""
	}
	p, err := filepath.Abs(file) 
	if err != nil {
		return ""
	}
	return filepath.Dir(p)
}
