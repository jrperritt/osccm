package openstack

import (
	"net/http"
	"net/http/httputil"

	"github.com/golang/glog"
)

type loggerClient struct {
	rt http.RoundTripper
}

func newHTTPLoggerClient() http.Client {
	return http.Client{
		Transport: &loggerClient{
			rt: http.DefaultTransport,
		},
	}
}

func (lc *loggerClient) RoundTrip(req *http.Request) (*http.Response, error) {
	reqbytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	glog.V(5).Infof("req: %s\n", string(reqbytes))

	resp, err := lc.rt.RoundTrip(req)
	if resp == nil {
		return nil, err
	}

	respbytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	glog.V(5).Infof("resp: %s\n", string(respbytes))

	return resp, nil
}
