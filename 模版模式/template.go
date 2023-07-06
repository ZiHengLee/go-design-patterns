package 模版模式

import "fmt"

//模拟一个短信发送服务
//基本步骤，检测号码，发送短信
//对于不同的运营商和不同的国家，其检测号码和发送的具体的逻辑是不一样的

type ISMS interface {
	send(content string, phone int) error
	valid(phone int) error
}

// BaseSms SMS 短信发送基类
type BaseSms struct {
	ISMS
}

// Send 发送短信
func (s *BaseSms) Send(content string, phone int) error {
	if err := s.valid(phone); err != nil {
		return err
	}
	// 调用子类的方法发送短信
	return s.send(content, phone)
}

// TelecomSms 走电信通道
type TelecomSms struct {
	*BaseSms
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}

func (tel *TelecomSms) valid(phone int) error {
	fmt.Println("its telecom phone")
	return nil
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.BaseSms = &BaseSms{ISMS: tel}
	return tel
}
