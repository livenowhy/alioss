package main
import "github.com/liuzhangpei/alioss/utils"

func main() {
//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。

	s := "?&actionType=ActionType-test&actionResourceId=actionResourceId-test&filename=user-dir%2Fsy_62331342149.jpg&size=94121&mimeType=image%2Fjpeg&height=691&width=1024"

	utils.NewCallbackActionType(s)
}