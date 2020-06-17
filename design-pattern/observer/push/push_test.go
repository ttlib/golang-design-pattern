// Author: Tong Zicheng
// Date: 2020/6/11 14:47
// Describe:

package push

import (
	"testing"
)

func Test(t *testing.T) {

	// 1. 制造一个明星 有一个粉丝列表
	idol := BigV{map[Observer]string{}}

	// 2. 制造很多粉丝
	fanA := Fan{"老张"}
	fanB := Fan{"老何"}
	fanC := Fan{"老吴"}

	// 3. 粉丝关注明星（反过来，明星把他们拉进来）
	idol.AddFan(&fanA)
	idol.AddFan(&fanB)
	idol.AddFan(&fanC)

	// 4. 明星发动态，粉丝获取动态
	idol.Notify("618演唱会，关注免费送门票")
	idol.DelFan(&fanA) // 拉黑一个粉丝
	idol.Notify("还有90张，家人们，快关注啊！")

}
