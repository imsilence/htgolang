package main

import "fmt"

// 定义Sender接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

type SingleSender interface {
	Send(to, msg string) error
}

// 定义EmailSender结构体
type EmailSender struct {
	addr           string
	port           int
	user, password string
}

// 接收者为值对象
func (sender EmailSender) Send(to, msg string) error {
	fmt.Printf("发送邮件给: %s, 内容: %s\n", to, msg)
	return nil
}

// 接收者为值对象
func (sender EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送邮件给: %s, 内容: %s\n", to, msg)
	}
	return nil
}

// 定义SMSSender结构体
type SMSSender struct {
	url, key string
}

func NewSMSSender(url, key string) *SMSSender {
	return &SMSSender{url, key}
}

// 接收者为指针对象
func (sender *SMSSender) Send(to, msg string) error {
	fmt.Printf("发送短信给: %s, 内容: %s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *SMSSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送短信给: %s, 内容: %s\n", to, msg)
	}
	return nil
}

// 定义WeChatSender结构体
type WeChatSender struct {
	sessionFile string
}

// 接收者为值对象
func (sender WeChatSender) Send(to, msg string) error {
	fmt.Printf("发送短信给: %s, 内容: %s\n", to, msg)
	return nil
}

// 接收者为指针对象
func (sender *WeChatSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		fmt.Printf("发送微信给: %s, 内容: %s\n", to, msg)
	}
	return nil
}

func main() {

	// 声明Sender接口对象
	var sender Sender
	fmt.Printf("%T: %v\n", sender, sender)

	// 赋值sender对象为NewEmailSender值对象
	var sender01 Sender = EmailSender{"smtp.qq.com", 465, "kk", "12356"}
	fmt.Printf("%T: %v\n", sender01, sender01)
	sender01.Send("imsilence@outlook.com", "Hi Silence")
	sender01.SendAll([]string{"imsilence@outlook.com", "iamkk@outlook.com"}, "早上好")

	// 赋值sender对象为NewSMSSender指针对象
	var sender02 Sender = &SMSSender{"http://v.juhe.cn/sms/send", "123"}
	fmt.Printf("%T: %v\n", sender02, sender02)
	sender02.Send("15200000000", "Hi KK")
	sender02.SendAll([]string{"15200000000", "15800000000"}, "早上好")

	// 赋值sender对象为NewSMSSender指针对象
	//var sender03 Sender = WeChatSender{"/tmp/wx.session"}  // ???
	//var sender03 Sender = &WeChatSender{"/tmp/wx.session"} // ???
	var sender03 Sender = &WeChatSender{"/tmp/wx.session"}
	fmt.Printf("%T: %v\n", sender03, sender03)
	sender03.Send("786725806", "Hi KK")
	sender03.SendAll([]string{"786725806", "786725807"}, "早上好")

	//var sender04 Sender = &EmailSender{"smtp.qq.com", 465, "kk", "12356"}  // ???
	//var sender05 Sender = SMSSender{"http://v.juhe.cn/sms/send", "123"} // ???

	var ssender SingleSender = sender03
	ssender.Send("786725806", "Hi KK")

	// 使用类型断言进行转换
	sender04, ok := ssender.(Sender)
	fmt.Printf("%T, %#v, %v\n", sender04, sender04, ok)
	sender03.SendAll([]string{"786725806", "786725807"}, "早上好")

	emailSender, ok := sender01.(EmailSender)
	fmt.Printf("%T, %#v, %v\n", emailSender, emailSender, ok)
	if ok {
		fmt.Println(emailSender.addr)
	}

	if smsSender, ok := sender02.(*SMSSender); ok {
		fmt.Printf("%T, %#v, %v\n", smsSender, smsSender, ok)
		fmt.Println(smsSender.url)
	}

	if smsSender02, ok := sender03.(*SMSSender); !ok {
		fmt.Printf("%T, %#v, %v\n", smsSender02, smsSender02, ok)
	}

	// 使用类型查询
	senders := []SingleSender{sender, ssender, sender01, sender02, sender03, sender04}
	for _, s := range senders {
		switch v := s.(type) {
		case EmailSender:
			fmt.Printf("EmailSender: %#v, %v\n", v, v.addr)
		case *SMSSender:
			fmt.Printf("SMSSender: %#v, %v\n", v, v.url)
		// case *WeChatSender:
		// 	fmt.Printf("WeChatSender: %#v, %v\n", v, v.sessionFile)
		case Sender:
			fmt.Printf("Sender: %#v, %v\n", v, v)
			v.SendAll([]string{"one", "two"}, "sender")
		default:
			fmt.Printf("error, %#v\n", v)
		}
	}
}
