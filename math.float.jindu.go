package comm

import (
	"strconv"
	"strings"
)

//get jd fro float type
func FloatPoint(f float64, jd int) float64 {
	f1 := strconv.FormatFloat(f, 'f', jd, 64)
	str := strings.Split(f1, ".")
	var s string
	if len(str) == 2 {
		s = str[0] + "." + str[1][0:1]
	}
	f2, _ := strconv.ParseFloat(s, 64)
	return f2
}
