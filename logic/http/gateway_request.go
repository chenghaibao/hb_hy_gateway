package http

import (
	"fmt"
	"hb_hy_gateway/utils"
	"io/ioutil"
	"math/rand"
	"net/http"
)

var family map[string][]string

func (g *Gateway) sendRequest(w http.ResponseWriter, r *http.Request) bool {
	//流量分发 （模拟map）
	family =  make(map[string][]string)
	family["family"] = []string{"http://127.0.0.1:9088", "http://127.0.0.1:9099"}
	fmt.Println(family["family"])
	requestSuccess := g.addressSelect(w, r)
	//
	////请求失败
	if requestSuccess == true {
		g.jsonErrorWithTraceId(w, "22222", 403, "request failed")
		return false
	}
	return true

}

func (g *Gateway) request(w http.ResponseWriter, r *http.Request, sendGateway string) bool {
	count := 3
	//端口流量转发  判断请求类型  模拟的我默认就都是get了
	resp, err := http.Get(fmt.Sprintf("%s%s", sendGateway, r.URL.Path))

	//复制request
	if err != nil {
		isRequest, secondRequest := isRequestRetry(count, sendGateway, r.URL.Path)
		if isRequest == true {
			resp = secondRequest
		}else{
			return true
		}
	}

	utils.CheckError(err)

	//关闭请求
	defer close(resp)
	body, err := ioutil.ReadAll(resp.Body)

	utils.CheckError(err)

	//返回参数
	g.jsonSimpleTextSuccess(w,string(body))
	return false
}

func close(resp *http.Response) {
	resp.Body.Close()
}

func (g *Gateway) addressSelect(w http.ResponseWriter, r *http.Request) bool {
	// 获取随机地址  没有写权重和ip_hash策略
	count := len(family["family"])
	randNum := rand.Intn(count)
	sendGateway := family["family"][randNum]
	return g.request(w, r, sendGateway)
}

//是否多次请求
func isRequestRetry(count int, address string, url string) (bool, *http.Response) {
OuterLoop:
	for i := 0; i < count; i++ {
		if i < 1 {
			//删除模拟map  正常是删除没用的ip
			family["family"] = append(family["family"][:0], family["family"][1])
			break OuterLoop
		} else {
			//多次请求
			secondRequest, err := http.Get(fmt.Sprintf("%s/%s", address, url))
			if err == nil {
				return true, secondRequest
			}
		}
	}
	//返回false 没有成功返回一个空的http
	return false, &http.Response{}
}
