package server




import (
	"net/http"
	"github.com/liuzhangpei/alioss/utils"
)


// 获取
func (cg *Config)ServerTest(w http.ResponseWriter, r *http.Request) {
	utils.ResponseError(w, "OK", "test is o00s")
	return
}



