package service

import (
	. "feishu-chatbot/models"
	"fmt"
	"testing"
)

// TestGenMsg 测试生成返回值
func TestGenMsg(t *testing.T) {
	msg, _, err := GenMsg("Hello", "oc_151692536b0b2d4a0873872d980ae59f", "om_f6f147ebe36cf26d55e9afe47b9b6d55", "1710659469323")
	if err != nil {
		t.Error(err)
	}
	if msg != "Hi, What can I do for you?[Smile]" {
		t.Error("Hello should respond: Hi, What can I do for you?[Smile]")
	}

}

func TestTransactionSucceed(t *testing.T) {
	reply, _, err := TransactionSucceed(Reply{
		Content:    "Payment Failed",
		ChatId:     "oc_151692536b0b2d4a0873872d980ae59f",
		MessageId:  "om_b3ab347039fc47e6858abdbcab977886",
		CreateTime: "1710604968113",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("response is: %s", reply)
	if reply != "Sorry, Pay Failed[CrossMark]，Please try again!" {
		t.Error("Payment Failed should respond: Sorry, Pay Failed[CrossMark]，Please try again! ")
	}
}

func TestGenMsg2(t *testing.T) {
	msg, _, err := GenMsg("thanks", "oc_151692536b0b2d4a0873872d980ae59f", "om_f6f147ebe36cf26d55e9afe47b9b6d55", "1710659469323")
	if err != nil {
		t.Error(err)
	}
	if msg != "It's my pleasure [Smile]" {
		t.Error("Thanks should respond: It's my pleasure [Smile]")
	}
}

func TestGenMsg3(t *testing.T) {
	msg, _, err := GenMsg("23333", "oc_151692536b0b2d4a0873872d980ae59f", "om_f6f147ebe36cf26d55e9afe47b9b6d55", "1710659469323")
	if err != nil {
		t.Error(err)
	}
	if msg != "Sorry, I don't understand what you say![What?]" {
		t.Error("23333 should respond: Sorry, I don't understand what you say![What?]")
	}
}
