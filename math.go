package comm

import (
	"fmt"
	"strconv"
)

//四舍五入 int 为精度
func MathRounding(f float64, n int) float64 {
	s := fmt.Sprintf("%0."+strconv.Itoa(n)+"f", f)
	res, _ := strconv.ParseFloat(s, 64)
	return res
}
