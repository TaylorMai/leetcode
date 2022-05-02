package greedy

func RemoveDuplicateLetters(s string) string {
	res := make([]byte, 0)
	stack := make([]byte, 0, 26) //单调栈，栈顶为最大
	has := make(map[byte]struct{}, 26)
	m := make(map[byte]int) //当前剩余字符的出现频率
	for i := range s {
		m[s[i]]++
	}
	for i := range s {
		if _, ok := has[s[i]]; ok { //已在栈中的字符
			m[s[i]]-- //跳过了
			continue
		}

		l := len(stack)
		//栈顶的字符（最大），大于当前字符，说明可能有更优的排序
		for l != 0 && stack[l-1] >= s[i] {
			//如果pop至一个“当前仅有唯一一个”的字符时，说明该字符前已完成局部最优。而且只剩一个了，不能pop
			if m[stack[l-1]] == 1 {
				break
			}
			stack = stack[:l-1] //pop
			l--
			m[stack[l-1]]--         //pop出去了
			delete(has, stack[l-1]) //栈中没有该字符了
		}

		stack = append(stack, s[i])
		has[s[i]] = struct{}{}
	}
	for i := range stack {
		res = append(res, stack[i])
	}
	return string(res)
}
