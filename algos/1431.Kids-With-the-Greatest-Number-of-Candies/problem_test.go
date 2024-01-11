package algos

import (
	"testing"

	"github.com/xinxin001/goleet/common/utils"
)

func TestBase(t *testing.T) {
	got := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	want := []bool{true, true, true, false, true}
	utils.TestError(t, got, want)
}
