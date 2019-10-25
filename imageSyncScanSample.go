package main

import (
	"fmt"
	"github.com/Jurassic-Park/aliyun-green-go-sdk/greensdksample"
	"github.com/Jurassic-Park/aliyun-green-go-sdk/uuid"
)

const accessKeyId string = "<your access key id>"
const accessKeySecret string = "<your access key secret>"

func main() {
	profile := greensdksample.Profile{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret}

	path := "/green/image/scan"

	clientInfo := greensdksample.ClinetInfo{Ip: "127.0.0.1"}

	// 构造请求数据
	bizType := "Green"
	scenes := []string{"porn"}

	task := greensdksample.Task{DataId: uuid.Rand().Hex(), Url: "https://xxx.png"}
	tasks := []greensdksample.Task{task}

	extras := map[string]string{
		// 验证卡的类型
	}

	bizData := greensdksample.BizData{bizType, scenes, tasks, extras}

	var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile: profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, bizData))

}
