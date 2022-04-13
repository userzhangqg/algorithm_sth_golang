package codeTop

import (
	"strconv"
	"strings"
)

/*
165. 比较版本号

给你两个版本号version1和version2，请你比较它们

版本号由一个或多个修订号组成，各修订号由一个'.'连接。每个修订号由多位数字 组成，可能包含前导0。
每个版本号至少包含一个字符。
修订号从左到右编号，下标从0开始，最左边的修订号下标为0，下一个修订号下标为1，以此类推。
例如，2.5.33 和 0.1 都是有效的版本号。
比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较忽略任何前导零的整数值。
也就是说，修订号1和修订号001相等。如果版本号没有指定某个下标处的修订号，则该修订号视为0。
例如，版本1.0 小于 版本1.1，因为他们下标为0的修订号相同，而下标为1的版本号分别为0和1，0 < 1.
返回规则如下：
- 如果 version1 > version2 返回1
- 如果 version1 < version2 返回-1
- 除此以外返回0

$ Done
思路：
1. 分隔比较（注意细节），单个0，1比较。1.001和1.01比较
2. 双指针
*/

func compareVersion(version1 string, version2 string) int {
	v1Array := strings.Split(version1, ".")
	v2Array := strings.Split(version2, ".")

	for i := len(v1Array) - 1; i > 0; i-- {
		if v1Array[i] != "0" {
			break
		} else {
			v1Array = v1Array[:len(v1Array)-1]
		}
	}
	for i := len(v2Array) - 1; i > 0; i-- {
		if v2Array[i] != "0" {
			break
		} else {
			v2Array = v2Array[:len(v2Array)-1]
		}
	}

	for i, j := 0, 0; i < len(v1Array) || j < len(v2Array); i, j = i+1, j+1 {
		x, y := 0, 0
		if i < len(v1Array) {
			x, _ = strconv.Atoi(v1Array[i])
		}
		if j < len(v2Array) {
			y, _ = strconv.Atoi(v2Array[i])
		}
		if x < y {
			return -1
		} else if x > y {
			return 1
		}
	}
	return 0
}
