package comm

import (
	"crypto/tls"
	"net/url"
	"time"

	"log"

	"github.com/astaxie/beego/httplib"
)

func GetJsonFromUrl(Url string) ([]byte, int, error) {
	var code int
	u, err := url.Parse(Url)
	if err != nil {
		log.Println("err1=", err)
		return nil, 501, err
	}

	switch u.Scheme {
	case "http":

		req := httplib.Get(Url)
		req.SetTimeout(30*time.Second, 30*time.Second)
		resp, err := req.Response()
		if err != nil {
			log.Println("err2=", err)
			code = 501
		}
		code = resp.StatusCode

		body, err := req.Bytes()

		return body, code, err

	case "https":
		req := httplib.Get(Url)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		resp, err := req.Response()
		if err != nil {
			log.Println("err3=", err)
			code = 501
			return nil, code, err
		}

		code = resp.StatusCode
		body, err := req.Bytes()

		return body, code, err
	default:
		log.Println("hit default")
		return nil, 501, nil
	}

}
