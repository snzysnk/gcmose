package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

}

type key struct{}
type key2 struct{}

func TestWithCancel(t *testing.T) {
	context.WithCancel(context.Background())
}

//  设子context生存时间为Y父context生存时间为X
//  子context实际生存时间为W
//  if Y > X {
//  W = X
//  }else {
//  W = Y
//  }
func TestTimeOutWithCoroutine(t *testing.T) {
	c := context.Background()
	c1, cancelFunc := context.WithTimeout(c, 5*time.Second)
	c2, cancelFunc2 := context.WithTimeout(c1, 8*time.Second)
	defer cancelFunc2()
	defer cancelFunc()

	//并行下子context的生存时间小于等于父context
	go task(c1, 2*time.Second, "context 2 秒超时")
	go task(c2, 4*time.Second, "context 4 秒done")

	time.Sleep(10 * time.Second)
}

func TestWithValue(t *testing.T) {
	c := context.Background()
	//使用context.WithValue(parent Context,key interface{},value interface{}) 向context中 写入 key => value 数据,返回一个新的context
	//parent 为 要写入的对象context
	//key 键值
	//value 值
	c = context.WithValue(c, key{}, "key1's value")

	//通过 context.Value(key interface{}) 获取 key 对应的值
	//key 键值,存在返回key对应的值，不存在则返回nil
	fmt.Println(c.Value(key{}), c.Value(key2{}))

	c2 := context.WithValue(c, key2{}, "key2's value")

	//c2继承了c的key=>value
	//证明 content可以继承(子承父业)
	fmt.Println(c2.Value(key{}), c2.Value(key2{}))
}

func TestContextTimeOut(t *testing.T) {
	background := context.Background()
	//withTimeout(parent context,timeout time.duration) 返回一个有生存时间的context
	//parent context
	//time 超时时间
	c, cancelFunc := context.WithTimeout(background, time.Second*10)
	c2, cancelFunc2 := context.WithTimeout(c, time.Second*11)
	defer cancelFunc()
	defer cancelFunc2()
	task(c, 7*time.Second, "测试过期")

	//证明子context的存货时间不能大于父context的存货时间
	//且其真实存货时间为 10 - 7 = 3 秒
	//通过下面三组函数依次对比 证明 子context的生存计算公式正确
	task(c2, 7*time.Second, "子context生存时间不能大于父context的剩余存活时间")
	//task(c2, 2*time.Second,"子context只能存货3秒") timeout
	//task(c2, 4*time.Second,"子context只能存货3秒") done
}

func task(c context.Context, sec time.Duration, taskName string) {
	//用来显示context的存活时间
	//context 存活时间小于 sec 输出 done
	//context 存活时间大于 sec 输出 timeout
	select {
	case <-c.Done():
		fmt.Printf("taskName %s done \n", taskName)
	case <-time.After(sec):
		fmt.Printf("timeout by taskNmae %s \n", taskName)
	}
}

func TestCancelContext(t *testing.T) {
	c := context.Background()
	//context.WithCancel(c context) 获取新的context和其终止方法cancelFunc
	//cancelFunc 能主动终止context
	c, cancelFunc := context.WithCancel(c)

	//开个协程2秒钟主动调用cancelFunc,终止
	//如果没有主动调用cancelFunc task 方法中因该是超时timeout 打印
	go func() {
		select {
		case <-time.After(time.Second * 2):
			cancelFunc()
		}
	}()

	task(c, time.Second*6, "test cancelFunc")
}

//设置主context生存时间为五秒。主程序休眠7秒后调用主context的cancelFunc
//会将子程序一起cancel掉
//得出主context的cancelFunc会将其后辈context一同cancel
func TestCancelWithCoroutineOne(t *testing.T) {
	background := context.Background()
	c1, cancelFunc := context.WithCancel(background)
	defer cancelFunc()
	c2, _ := context.WithCancel(c1)

	go task(c1, 5*time.Second, "父context")
	go task(c2, 10*time.Second, "子context")

	time.Sleep(7 * time.Second)
}

//设置主context 5 秒超时
//子context 3 秒超时
//当2秒时主动调用子context的cancelFunc
//主context正常输出超时结果
//得出子context的cancelFunc不会影响主context
func TestCancelWithCoroutineTwo(t *testing.T) {
	background := context.Background()
	c1, _ := context.WithCancel(background)
	c2, cancelFunc := context.WithCancel(c1)

	go task(c1, 5*time.Second, "父context")
	go task(c2, 3*time.Second, "子context")

	go func() {
		select {
		case <-time.After(2 * time.Second):
			cancelFunc()
		}
	}()

	time.Sleep(10 * time.Second)
}
