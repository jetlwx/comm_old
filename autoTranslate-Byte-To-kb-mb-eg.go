package comm

import (
	"strconv"
)

//计算机单位自动匹配

const (
	_        = iota // iota = 0
	KB int64 = 1 << (10 * iota)
	MB       //iota=1
	GB       // 与 KB 表达式相同,但 iota = 2
	TB
	PB
	YB
)

//将byte自动转换为小于1024的单位，如1024Byte自动返回Kb
func TransferByte(size int64) (val string) {
	aB := size
	aKB := (size / KB)
	aMB := (size / MB)
	aGB := (size / GB)
	aTB := (size / TB)
	aPB := (size / PB)
	aYB := (size / YB)
	f := int64(1024)
	var res int64
	var dw string

	if aB > f {
		if aKB > f {
			if aMB > f {
				if aGB > f {
					if aTB > f {
						if aPB > f {
							res = aYB
							dw = "YB"
						} else {
							res = aPB
							dw = "PB"
						}
					} else {
						res = aTB
						dw = "TB"
					}
				} else {
					res = aGB
					dw = "GB"
				}
			} else {
				res = aMB
				dw = "MB"
			}
		} else {
			res = aKB
			dw = "KB"
		}
	} else {
		res = aB
		dw = "Byte"
	}

	return strconv.FormatInt(res, 10) + dw
}
