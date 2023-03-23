package rest

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/lantaris/rest-go-sdk/fmtlog"
	"io"
	"io/ioutil"
	"net/http"
)

type Call struct {
	Retry int
}

// ********************************************************
func (SELF *Call) POST(URL string, data []byte) (rbody []byte, err error) {

	fmtlog.Debugln("Rest call POST:[" + URL + "] ")
	for RetryCnt := 1; RetryCnt <= SELF.Retry; RetryCnt++ {
		resp, err := http.Post(URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmtlog.Errorln("Error call to [" + URL + "]. Retry...  " + err.Error())
			continue
		}

		rbody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmtlog.Errorln("Error read data from [" + URL + "]. Retry...  " + err.Error())
			continue
		}

		return rbody, nil
	}
	err = errors.New("Rest call to [" + URL + "] retry exceeded")
	fmtlog.Errorln(err.Error())
	return nil, err
}

// ********************************************************
func (SELF *Call) GET(URL string) (rbody []byte, err error) {

	fmtlog.Debugln("Rest call to GET:[" + URL + "] ")

	for RetryCnt := 1; RetryCnt <= SELF.Retry; RetryCnt++ {
		resp, err := http.Get(URL)
		if err != nil {
			fmtlog.Errorln("Error call to [" + URL + "]. Retry...  " + err.Error())
			continue
		}

		rbody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmtlog.Errorln("Error read data from [" + URL + "]. Retry...  " + err.Error())
			continue
		}

		return rbody, nil
	}
	err = errors.New("Rest call to [" + URL + "] retry exceeded")
	fmtlog.Errorln(err.Error())
	return nil, err
}

// ======================================================================
func (SELF *Call) post(url string, header http.Header, body io.Reader) (ResHeader http.Header, ResData []byte, err error) {
	var (
		req    *http.Request
		client *http.Client
		res    *http.Response
	)
	method := "POST"
	client = &http.Client{}
	req.Header = header

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error create rest request to: [%s], Error: [%s]", url, err.Error()))
		return nil, nil, err
	}

	// Call with retry
	for RetryCnt := 1; RetryCnt <= SELF.Retry; RetryCnt++ {
		res, err = client.Do(req)
		if err != nil {
			fmtlog.Errorln(fmt.Sprintf("Error rest call to: [%s], Error: [%s]", url, err.Error()))
			continue
		}
	}

	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error rest call to: [%s], Error: [%s]", url, err.Error()))
		return nil, nil, err
	}

	defer res.Body.Close()

	ResData, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error unpack data rest call to: [%s], Error: [%s]", url, err.Error()))
		return nil, nil, err
	}
	ResHeader = res.Header

	return ResHeader, ResData, err
}
