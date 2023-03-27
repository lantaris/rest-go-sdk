package rest

import (
	"errors"
	"fmt"
	"github.com/lantaris/rest-go-sdk/fmtlog"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Call struct {
	Retry int
}

// ======================================================================
func (SELF *Call) Exec(Method string, Url string, Header http.Header, Body []byte) (ResCode int, ResHeader http.Header, ResData []byte, err error) {
	var (
		req    *http.Request
		client *http.Client
		res    *http.Response
		RdBody io.Reader
	)
	if SELF.Retry == 0 {
		SELF.Retry = 3
	}
	client = &http.Client{}

	fmtlog.Traceln(fmt.Sprintf("REST call to: [%s]", Url ))

	if Body != nil {
		RdBody = strings.NewReader(string(Body))
	} else {
		RdBody = nil
	}

	req, err = http.NewRequest(Method, Url, RdBody)
	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error create rest request to: [%s], Error: [%s]", Url, err.Error()))
		return 0, nil, nil, err
	}

	req.Header = Header

	// Call with retry
	for RetryCnt := 1; RetryCnt <= SELF.Retry; RetryCnt++ {
		res, err = client.Do(req)
		if err != nil {
			fmtlog.Errorln(fmt.Sprintf("Error rest call to: [%s], Error: [%s]", Url, err.Error()))
			continue
		} else  {
			break
		}
	}

	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error rest call to: [%s], Error: [%s]", Url, err.Error()))
		return 0, nil, nil, err
	}

	//defer res.Body.Close()
	if res.Body == nil {
		err = errors.New(fmt.Sprintf("Response body is nil"))
		fmtlog.Errorln(err.Error())
		return 0, nil, nil, err
	}
	ResData, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmtlog.Errorln(fmt.Sprintf("Error unpack data rest call to: [%s], Error: [%s]", Url, err.Error()))
		return 0, nil, nil, err
	}
	ResHeader = res.Header
	ResCode = res.StatusCode

	return ResCode, ResHeader, ResData, err
}
