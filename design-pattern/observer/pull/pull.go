// Author: Tong Zicheng
// Date: 2020/6/13 16:56
// Describe: 观察者模型-拉模型

package pull

import (
	"log"
)

// 1. 定义抽象观察者（Observer）：抽象的粉丝
type Observer interface {
	Update(sb Subject)
}

// 2. 具体的观察者对象：具体的粉丝
type Fan struct {
	// 微博详细信息
	blog    Blog
	fanName string
}

func (f *Fan) Update(sb Subject) {
	f.blog = sb.GetMessage()
}

// 3. 定义抽象主题（Subject）：抽象明星
type Subject interface {
	// 增加粉丝
	AddFan(fan Observer)
	// 拉黑粉丝
	DelFan(fan Observer)
	// 告诉粉丝我的状态 （通知每个观察者，非公开）
	notify()
	// 添加微博
	AddMessage(blog Blog)
	// 获取详细的微博内容
	GetMessage() Blog
}

// 4. 具体的主题对象：微博大V
type BigV struct {
	// 微博内容
	blog    Blog
	fanList map[Observer]string
}

func (b *BigV) AddFan(fan Observer) {
	b.fanList[fan] = "0"
}

func (b *BigV) DelFan(fan Observer) {
	delete(b.fanList, fan)
}

func (b *BigV) notify() {
	for k, _ := range b.fanList {
		k.Update(b) //通知每一个粉丝消息
	}
}

func (b *BigV) AddMessage(blog Blog) {
	b.blog = blog
	log.Println("[pull]", "发布一条新微博:", b.blog.Title)
	b.notify()
}

func (b *BigV) GetMessage() Blog {
	return b.blog
}

// 5. 定义微博对象：一个详细的微博对象
type Blog struct {
	Id      int32
	Title   string
	Content string
}
