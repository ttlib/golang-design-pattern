// Author: Tong Zicheng
// Date: 2020/6/17 20:57
// Describe:

package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// 1. 我们下一碗方便面
	instantNoodles := InstantNoodles{"老坛酸菜牛肉面", 4}
	fmt.Println(instantNoodles.Description())
	fmt.Println(instantNoodles.Price(), "元")
	fmt.Println("--------------------------------")

	// 2. 面条里面加个蛋（鸡蛋方便面）
	eggInstantNoodles := Egg{&instantNoodles, "鸡蛋", 0.5}
	fmt.Println(eggInstantNoodles.Description())
	fmt.Println(eggInstantNoodles.Price(), "元")
	fmt.Println("--------------------------------")

	// 3. 在加一个火腿肠（鸡蛋火腿肠方便面）
	HamAndEggInstantNoodles := HamSausage{&eggInstantNoodles, "火腿肠", 2.0}
	fmt.Println(HamAndEggInstantNoodles.Description())
	fmt.Println(HamAndEggInstantNoodles.Price(), "元")
	fmt.Println("--------------------------------")

	// 4. 不够吃再来一个蛋（双蛋火腿肠老坛酸菜牛肉面）
	HamAnd2EggInstantNoodles := Egg{&HamAndEggInstantNoodles, "鸡蛋2", 0.5}
	fmt.Println(HamAnd2EggInstantNoodles.Description())
	fmt.Println(HamAnd2EggInstantNoodles.Price(), "元")
	fmt.Println("---------------这酸爽！-----------------")
}
