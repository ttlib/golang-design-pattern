// Author: Tong Zicheng
// Date: 2020/6/16 22:46
// Describe:

package pull

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	// 1. 制造一个明星 有一个粉丝列表
	idol := BigV{fanList: map[Observer]string{}}

	// 2. 制造很多粉丝
	fanA := Fan{fanName: "老张"}
	fanB := Fan{fanName: "老何"}
	fanC := Fan{fanName: "老吴"}

	// 3. 粉丝关注明星（反过来，明星把他们拉进来）
	idol.AddFan(&fanA)
	idol.AddFan(&fanB)
	idol.AddFan(&fanC)

	// 4. 明星发布一条新微博
	blog := Blog{Id: 1, Title: "618演唱会", Content: "618演唱会，关注免费送门票"}
	idol.AddMessage(blog)

	// 5. 粉丝获取博客部分信息
	fmt.Println(fanA.blog.Title)
	fmt.Println(fanB.blog.Content)
	fmt.Println(fanA.blog.Id)
}
