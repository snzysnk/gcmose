package text

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

type A struct {
	Name string
	Age  int64
}

func TestCmp(t *testing.T) {
	a := A{
		Name: "xie_rui_xiang",
		Age:  28,
	}
	b := A{
		Name: "xie_rui_xiang",
		Age:  28,
	}
	c := A{
		Name: "xie_rui_xiang",
		Age:  29,
	}

	fmt.Println(cmp.Diff(a, b))
	fmt.Println(cmp.Diff(b, c))

	//如果比较的数据包含proto生成的struct，则需要使用 protocmp.Transform()
	//下面是假设a，b包含proto生成的struct
	fmt.Println(cmp.Diff(a, b, protocmp.Transform()))
}
