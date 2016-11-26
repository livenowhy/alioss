package main
import "fmt"
import "net/url"
func main() {
//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
    s := "path?k=v&sdsd=sdsd&dffdd=Effr3ef"
	s = "action={\"actionType\":\"ActionType-test\",\"actionResourceId\":\"actionResourceId-test\"}&filename=user-dir%2Fsy_62331342149.jpg&size=94121&mimeType=image%2Fjpeg&height=691&width=1024"
//解析这个 URL 并确保解析没有出错。
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。
	// 你也可以将查询参数解析为一个map。
	// 已解析的查询参数 map 以查询字符串为键，对应值字符串切片为值，
	// 所以如何只想得到一个键对应的第一个值，将索引位置设置为 [0] 就行了。
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)

	for k, v := range m{
		fmt.Println(k, v)
	}
    fmt.Println(m)
}