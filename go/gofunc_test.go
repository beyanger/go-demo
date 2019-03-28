package gofunc

// go test .
import (
	"math"
	"testing"
)

func add(i, j int) int {
	return i + j
}
// TestMain会首先执行,可以执行一些初始化操作
func TestMain(m *testing.M) {
	// 如果不调用 m.Run()，那么所有的TestCase不会执行
	m.Run()
}

// 子test，所有的TestCase不保证顺序执行，如果需要顺序执行，可以作为子test

func testSub1(t *testing.T) {

}

func testSub2(t *testing.T) {
}


func TestSet(t *testing.T) {
	t.Run("sub1", testSub1)
	t.Run("sub2", testSub2)
}


func TestAdd(t *testing.T) {
	// 会跳过当前testing
	t.SkipNow()


	testCase := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{4, 5, 10},
		{math.MaxInt64, 1, math.MinInt64},
	}

	for _, tc := range testCase {
		r := add(tc.a, tc.b)
		if r != tc.c {
			t.Errorf("test add(%d, %d) expected: %d, but actual: %d\n",
				tc.a, tc.b, tc.c, r)
		}

	}
}


// 这个test在bench之中永远跑不完，无法达到稳态，注意要避免
func test(n int) bool {
	for ; n > 0; n-- {
	}
	return false
}

// go test -bench=.

func BenchmarkAll(b *testing.B) {
	// 测试达到一个稳定状态后就会停止，如果不稳定，就一直不会退出
	for i := 0; i < b.N; i++ {
		if test(i) {
			b.Error("error")
		}
	}
}

