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

// ======================================================================
func (SELF *Call) Exec(url string, header http.Header, body io.Reader) (ResHeader http.Header, ResData []byte, err error) {
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
