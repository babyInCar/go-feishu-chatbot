package calc

import (
	"fmt"
	"strings"
)

func SendMsg(str string) (string, error) {
	fmt.Println(str)

	// 把字母统一转程小写
	expression := strings.ToLower(str)
	var response string
	//out, _ := expression.Evaluate(nil)
	if expression == "hello" {
		response = "Hi, What can I do for you!"
	} else if strings.Contains(expression, "thanks") {
		response = "It's my pleasure"
	} else if strings.Contains(expression, "success") && strings.Contains(expression, "pay") {
		response = "Congratulations, Pay success!"
	} else {
		response = "Sorry, I don't understand what you say!"
	}
	return response, nil
}

func FormatMathOut(out string) string {
	//if is int
	//if out == float64(int(out)) {
	//	return fmt.Sprintf("%d", int(out))
	//}
	return fmt.Sprintf("%f", out)
}
