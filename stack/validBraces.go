package main


/**
有效括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 */
func isValid(s string) bool {
	arr := []string{}
	for i:=0;i<len(s) ; i++ {
		l := s[i]
		l1 := string(l)
		if l1 == "(" || l1 == "[" || l1 == "{"{
			arr = append(arr, l1)
		}else {
			if len(arr) == 0 {
				return false
			}
			p := arr[len(arr)-1]

			if l1 == ")" && p == "(" {
				arr = arr[:len(arr) -1]
			}else if l1 == "]" && p == "["{
				arr = arr[:len(arr) -1]
			}else if l1 == "}" && p == "{" {
				arr = arr[:len(arr) -1]
			}else {
				return false
			}
		}
	}

	if len(arr) == 0 {
		return true
	}
	return false
}

func main() {

}