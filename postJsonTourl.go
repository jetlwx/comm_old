package comm

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"log"
	"time"
)

type PostRes struct {
	Code int
	Data string
	URL  string
	Msg  error
}

func PostAndGetResponseFromAPI(url string, body []byte) ([]byte, error) {
	req := httplib.Post(url)
	req.SetTimeout(60*time.Second, 60*time.Second)
	req.Header("Content-Type", "application/json")
	req.Header("charset", "utf-8")
	req.Body(body)
	st, err := req.Bytes()
	return st, err
}


func PostJsonBody(url string, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[EE]  Get an error", err)
		}
	}()
	req := httplib.Post(url)
	req.SetTimeout(5*time.Second, 5*time.Second)
	req.Header("Content-type", "application/json")
	req.Body(data)
	res, err := req.Response()
	if err != nil {
		log.Println("[E24] an error at req.response:", err)
	}
	r := PostRes{}
	r.Code = res.StatusCode
	r.Data, _ = req.String()
	r.URL = url
	r.Msg = err
	d, err := json.Marshal(r)
	if err != nil {
		log.Println("[E] Get an error at json.Marshal:", err)
	}
	log.Println("[D] response from server:", string(d))

}

// has return value
func PostJsonBody2(url string, data []byte) []byte {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[EE]  Get an error", err)
		}
	}()
	log.Println("[D] post url=", url)
	req := httplib.Post(url)
	req.SetTimeout(5*time.Second, 5*time.Second)
	req.Header("Content-type", "application/json")
	req.Body(data)
	res, err := req.Response()
	r := PostRes{}
	r.Code = res.StatusCode
	r.Data = string(data)
	r.URL = url
	r.Msg = err
	d, _ := json.Marshal(r)
	return d
}
