package test

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

}

type key struct{}

func TestWithValue(t *testing.T) {
	c := context.Background()
	//使用context.WithValue() 向context中 写入 key => value 数据,返回一个新的context
	//第一个参数 为 要写入的对象context
	//第二个参数 为 key interface{}
	//第三个参数 为 value interface{}
	c = context.WithValue(c, key{}, "val")

	//通过 context.Value() 获取 key 对应的值
	//第一个参数 要获取的key,存在返回key对应的值，不存在则返回nil
	fmt.Println(c.Value(key{}))
}
