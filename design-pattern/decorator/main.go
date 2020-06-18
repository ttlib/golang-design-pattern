// Author: Tong Zicheng
// Date: 2020/6/17 15:33
// Describe:

package decorator

// 1. 定义抽象的被装饰者
type Noodles interface {
	Description() string
	Price() float32
}

// 2. 定义方便面（被装饰者）
type InstantNoodles struct {
	description string
	price       float32
}

func (i *InstantNoodles) Description() string {
	return i.description
}

func (i *InstantNoodles) Price() float32 {
	return i.price
}

// 3. 定义鸡蛋（装饰者A - 加了鸡蛋的方便面）
type Egg struct {
	Noodles
	description string
	price       float32
}

func (e *Egg) Description() string {
	return e.Noodles.Description() + "+" + e.description
}

func (e *Egg) Price() float32 {
	return e.Noodles.Price() + e.price
}

// 4. 定义火腿肠（装饰者B - 加了火腿肠的方便面）
type HamSausage struct {
	Noodles
	description string
	price       float32
}

func (h *HamSausage) Description() string {
	return h.Noodles.Description() + "+" + h.description
}

func (h *HamSausage) Price() float32 {
	return h.Noodles.Price() + h.price
}
