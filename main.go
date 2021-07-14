//go:generate goversioninfo
package main

import "C"
import (
	"crypto/tls"
	"encoding/json"
	"github.com/axgle/mahonia"
	"gopkg.in/resty.v1"
	"time"
)

//export HttpGet
func HttpGet(reqUrl *C.char, resp **C.char, headers *C.char, proxyUrl *C.char) C.int {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(15 * time.Second)
	// proxy
	if proxyUrl != nil {
		client.SetProxy(C.GoString(proxyUrl))
	}
	if headers != nil {
		m := make(map[string]string)
		err := json.Unmarshal([]byte(C.GoString(headers)), &m)
		if err == nil {
			return -1
		}
		client.SetHeaders(m)
	}
	response, err := client.R().Get(C.GoString(reqUrl))
	if err != nil {
		return -3
	}
	encoder := mahonia.NewEncoder("GBK")
	output := encoder.ConvertString(string(response.Body()))
	*resp = C.CString(output)
	return 0
}

//export HttpPost
func HttpPost(reqUrl *C.char, data *C.char, resp **C.char, headers *C.char, proxyUrl *C.char) C.int {
	var dataMap map[string]string
	err := json.Unmarshal([]byte(C.GoString(data)), &dataMap)
	if err != nil {
		return -2
	}
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(15 * time.Second)
	// proxy
	if proxyUrl != nil {
		client.SetProxy(C.GoString(proxyUrl))
	}
	client.SetFormData(dataMap)
	if headers != nil {
		m := make(map[string]string)
		err := json.Unmarshal([]byte(C.GoString(headers)), &m)
		if err == nil {
			return -1
		}
		client.SetHeaders(m)
	}
	response, err := client.R().Post(C.GoString(reqUrl))
	if err != nil {
		return -3
	}
	encoder := mahonia.NewEncoder("GBK")
	output := encoder.ConvertString(string(response.Body()))
	*resp = C.CString(output)
	return 0
}

//export HttpPostJson
func HttpPostJson(reqUrl *C.char, data *C.char, resp **C.char, headers *C.char, proxyUrl *C.char) C.int {
	var dataMap map[string]string
	err := json.Unmarshal([]byte(C.GoString(data)), &dataMap)
	if err != nil {
		return -2
	}
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(15 * time.Second)
	// proxy
	if proxyUrl != nil {
		client.SetProxy(C.GoString(proxyUrl))
	}
	if headers != nil {
		m := make(map[string]string)
		err := json.Unmarshal([]byte(C.GoString(headers)), &m)
		if err == nil {
			return -1
		}
		client.SetHeaders(m)
	}
	response, err := client.R().SetBody(dataMap).Post(C.GoString(reqUrl))
	if err != nil {
		return -3
	}
	encoder := mahonia.NewEncoder("GBK")
	output := encoder.ConvertString(string(response.Body()))
	*resp = C.CString(output)
	return 0
}

//export UploadFile
func UploadFile(reqUrl *C.char, filedName *C.char, filePath *C.char, resp **C.char, data *C.char, headers *C.char, proxyUrl *C.char) C.int {
	// proxy
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// client.SetTimeout(15 * time.Second)
	// proxy
	if proxyUrl != nil {
		client.SetProxy(C.GoString(proxyUrl))
	}
	if headers != nil {
		m := make(map[string]string)
		err := json.Unmarshal([]byte(C.GoString(headers)), &m)
		if err == nil {
			return -1
		}
		client.SetHeaders(m)
	}
	if data != nil {
		var dataMap map[string]string
		err := json.Unmarshal([]byte(C.GoString(data)), &dataMap)
		if err != nil {
			return -2
		}
		client.SetFormData(dataMap)
	}
	response, err := client.R().
		SetFile(C.GoString(filedName), C.GoString(filePath)).
		SetContentLength(true).
		Post(C.GoString(reqUrl))
	if err != nil {
		return -3
	}
	encoder := mahonia.NewEncoder("GBK")
	output := encoder.ConvertString(string(response.Body()))
	*resp = C.CString(output)
	return 0
}

//export DownloadFile
func DownloadFile(reqUrl *C.char, filename *C.char, headers *C.char, proxyUrl *C.char) C.int {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// proxy
	if proxyUrl != nil {
		client.SetProxy(C.GoString(proxyUrl))
	}
	if headers != nil {
		m := make(map[string]string)
		err := json.Unmarshal([]byte(C.GoString(headers)), &m)
		if err == nil {
			return -1
		}
		client.SetHeaders(m)
	}
	_, err := client.R().
		SetOutput(C.GoString(filename)).
		Get(C.GoString(reqUrl))
	if err != nil {
		return -3
	}
	return 0
}

func main() {
	// export dll
	// test
	// client := resty.New()
	// client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// client.R().
	// 	SetOutput("./test.fr3").
	// 	Get("https://39.100.244.169/sip-api//baseReport/baseReportDownload/784")
}
