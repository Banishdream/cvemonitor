package tool

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Request struct {
	Method string
	Url    string
	Value  string
	Ps     *Params
}

type Params struct {
	Timeout     int //超时时间
	Retry       int //重试次数
	Headers     map[string]string
	ContentType string
}

func (req *Request) Do(result interface{}) error {
	// 获取 请求的 body
	res, err := asyncCall(doRequest, req)
	if err != nil {
		return err
	}

	// 传入的结构体对象为空,返回nil
	if result == nil {
		return nil
	}

	// 将body数据解析成 result结构体
	if err = Decode(*res, &result); err != nil {
		return err
	}

	return nil
}

//type timeout struct {
//	body *io.ReadCloser
//	err  error
//}

func doRequest(request *Request) (*io.ReadCloser, error) {
	var (
		req    *http.Request
		errReq error
	)
	if request.Value != "null" {
		buf := strings.NewReader(request.Value)
		req, errReq = http.NewRequest(request.Method, request.Url, buf)
		if errReq != nil {
			return nil, errReq
		}
	} else {
		req, errReq = http.NewRequest(request.Method, request.Url, nil)
		if errReq != nil {
			return nil, errReq
		}
	}

	client := http.Client{
		Timeout: time.Duration(request.Ps.Timeout) * time.Millisecond,
	}
	res, err := client.Do(req)
	fmt.Println(1)
	fmt.Println(res, err)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(res.Body)

	return &res.Body, nil
}

/*
asyncCall
req中的 retry 重试次数, 和 timeout超时时间设置
当超时的时候发起一次新的请求
*/
func asyncCall(f func(request *Request) (*io.ReadCloser, error), req *Request) (*io.ReadCloser, error) {
	p := req.Ps
	// 重试的时候只有上一个http请求真的超时了，之后才会发起一次新的请求
	for i := 0; i < p.Retry; i++ {
		// 发送HTTP请求
		res, err := f(req)
		// 判断超时
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			continue
		}
		return res, err
	}
	return nil, errors.New("超时啦~")
}
