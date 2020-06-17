// Author: Tong Zicheng
// Date: 2020/6/11 14:37
// Describe: 观察者模式-推模式

package push

import "fmt"

// 1. 定义抽象观察者（Observer）：抽象的粉丝
type Observer interface {
	Update(message string)
}

// 2. 具体的观察者对象：具体的粉丝
type Fan struct {
	fanName string
}

func (f *Fan) Update(message string) {
	fmt.Println(f.fanName + "收到了：" + message + "通知")
}

// 3. 定义抽象主题（Subject）：抽象明星
type Subject interface {
	// 增加粉丝
	AddFan(fan Observer)
	// 拉黑粉丝
	DelFan(fan Observer)
	// 告诉粉丝我的状态
	Notify(msg string)
}

// 4. 具体的主题对象：微博大V
type BigV struct {
	fanList map[Observer]string
}

func (b *BigV) AddFan(fan Observer) {
	b.fanList[fan] = "0"
}

func (b *BigV) DelFan(fan Observer) {
	delete(b.fanList, fan)
}

func (b *BigV) Notify(msg string) {
	for k, _ := range b.fanList {
		k.Update(msg) //通知每一个粉丝消息
	}
}
