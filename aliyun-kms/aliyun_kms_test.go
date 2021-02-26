package aliyun_kms

import (
	"encoding/json"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var client *kms.Client
var AccessKeyId = tea.String("testKeyId")
var AccessKeySecret = tea.String("testKeySecret")

func init() {
	var err error
	if client == nil {
		client, err = kms.NewClientWithAccessKey("cn-hangzhou", *AccessKeyId, *AccessKeySecret)
		if err != nil {
			panic(fmt.Errorf("creat client error:%s", err))
		}
	}
}

func Test_GetRegions(t *testing.T) {
	ecsClient, err := ecs.NewClientWithAccessKey("cn-hangzhou", *AccessKeyId, *AccessKeySecret)
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"

	response, err := ecsClient.DescribeRegions(request)
	assert.NoError(t, err)
	fmt.Printf("response is %#v\n", response)
}

func ReadAsJSON(byt []byte) (result interface{}, err error) {
	err = json.Unmarshal(byt, &result)
	log.Println("the json umarshal error", err)
	return
}

func Test_creatKey(t *testing.T) {
	client, err := kms.NewClientWithAccessKey("cn-hangzhou", *AccessKeyId, *AccessKeySecret)

	request := kms.CreateCreateKeyRequest()
	request.Scheme = "https"

	response, err := client.CreateKey(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func Test_keyList(t *testing.T) {
	request := kms.CreateListKeysRequest()
	request.Scheme = "https"

	response, err := client.ListKeys(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func Test_encrypt(t *testing.T){
	request := kms.CreateEncryptRequest()
	request.Scheme = "https"
	request.Plaintext = "testKms"
	request.KeyId = "a699a5a9-482b-463b-8614-05c2f930b00d"
	response, err := client.Encrypt(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
