package tencent

import (
	"TencentSms/config"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Sms struct {
}

func getCredential() *common.Credential {
	return common.NewCredential(
		config.Secrets.GetString("TencentCloudSectet.SecretId"),
		config.Secrets.GetString("TencentCloudSectet.SecretKey"),
	)
}

func getClient() *sms.Client {
	credential := getCredential()
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.Confs.GetString("SmsConf.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, config.Confs.GetString("SmsConf.Region"), cpf)
	return client
}

func (*Sms) Send(phoneNumbers []string, params []string) (string, error) {
	client := getClient()
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phoneNumbers)
	request.SmsSdkAppId = common.StringPtr("1400818063")
	request.SignName = common.StringPtr("mybackgar公众号")
	request.TemplateId = common.StringPtr("1785595")
	request.TemplateParamSet = common.StringPtrs(params)

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", err
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	return response.ToJsonString(), nil
}
