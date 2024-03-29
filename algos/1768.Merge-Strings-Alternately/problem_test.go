package algos

import (
	"testing"

	"github.com/xinxin001/goleet/common/utils"
)

func TestBase(t *testing.T) {
	res := mergeAlternately("abc", "pqr")
	want := "apbqcr"
	utils.TestError(t, res, want)
}
