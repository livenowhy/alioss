package utils

import (
	"net/url"
	"fmt"
)


type CallbackActionType struct {
	Uuid string    `json:"uuid"`   // 用户uuid
	ActionType     string   `json:"actionType"`           // add lzp 上传动作(用户头像,镜像头像)
	ActionResourceId string   `json:"actionResourceId"`  // add lzp 资源id
    Filename string `json:"filename"`  // 文件名
    Size string `json:"size"`  // 文件名
    MimeType string `json:"mimeType"`  // 文件名
    Height string `json:"height"`  // 文件名
    Width string `json:"width"`  // 文件名
}

func NewCallbackActionType(urlstr string) (*CallbackActionType, error) {

	u, err := url.Parse(urlstr)

	if err != nil {
		return nil, err
	}

	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。
	// 你也可以将查询参数解析为一个map。
	// 已解析的查询参数 map 以查询字符串为键，对应值字符串切片为值，
	// 所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
	mp, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	var actionT CallbackActionType
	for k, v := range mp{

		if "actionType" == k {
			actionT.ActionType = v[0]
		} else if "actionResourceId" == k {
			actionT.ActionResourceId = v[0]
		} else if "uuid" == k {
			actionT.Uuid = v[0]
		} else if "filename" == k {
			actionT.Filename = v[0]
		} else if "size" == k {
			actionT.Size = v[0]
		} else if "mimeType" == k {
			actionT.MimeType = v[0]
		} else if "height" == k {
			actionT.Height = v[0]
		} else if "width" == k {
			actionT.Width = v[0]
		}
		fmt.Println(k, v)
	}

	fmt.Println("%+v", actionT)

	return &actionT, nil
}
