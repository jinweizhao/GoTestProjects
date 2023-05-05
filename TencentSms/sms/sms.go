package sms

import "TencentSms/sms/tencent"

type smsInterface interface {
	Send(phoneNumbers []string, params []string) (string, error)
}

func GetSmsSender() smsInterface {
	return &tencent.Sms{}
}
