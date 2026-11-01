package conf

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/transport/internet/headers/http"
	"github.com/xtls/xray-core/transport/internet/headers/noop"
	"github.com/xtls/xray-core/transport/internet/tcp"
	"github.com/xtls/xray-core/transport/internet/websocket"
	"google.golang.org/protobuf/proto"
)

type NoOpConnectionAuthenticator struct{}

func (NoOpConnectionAuthenticator) Build() (proto.Message, error) {
	return new(noop.Config), nil
}

type AuthenticatorRequest struct {
	Version string                 `json:"version"`
	Method  string                 `json:"method"`
	Path    StringList             `json:"path"`
	Headers map[string]*StringList `json:"headers"`
}

func (v *AuthenticatorRequest) Build() (*http.RequestConfig, error) {
	config := &http.RequestConfig{
		Uri: []string{"/"},
	}
	if len(v.Version) > 0 {
		config.Version = &http.Version{Value: v.Version}
	} else {
		config.Version = &http.Version{Value: "1.1"}
	}
	if v.Method != "" {
		config.Method = &http.Method{Value: v.Method}
	} else {
		config.Method = &http.Method{Value: "GET"}
	}
	if len(v.Path) > 0 {
		config.Uri = append([]string(nil), v.Path...)
	}
	if len(v.Headers) > 0 {
		config.Header = make([]*http.Header, 0, len(v.Headers))
		for key, value := range v.Headers {
			if value == nil {
				return nil, errors.New("empty HTTP header value: ", key).AtError()
			}
			config.Header = append(config.Header, &http.Header{
				Name:  key,
				Value: *value,
			})
		}
	}
	return config, nil
}

type AuthenticatorResponse struct {
	Version string                 `json:"version"`
	Status  string                 `json:"status"`
	Reason  string                 `json:"reason"`
	Headers map[string]*StringList `json:"headers"`
}

func (v *AuthenticatorResponse) Build() (*http.ResponseConfig, error) {
	config := &http.ResponseConfig{}
	if len(v.Version) > 0 {
		config.Version = &http.Version{Value: v.Version}
	}
	if len(v.Status) > 0 || len(v.Reason) > 0 {
		code, reason := "200", "OK"
		if len(v.Status) > 0 {
			code = v.Status
		}
		if len(v.Reason) > 0 {
			reason = v.Reason
		}
		config.Status = &http.Status{Code: code, Reason: reason}
	}
	if len(v.Headers) > 0 {
		config.Header = make([]*http.Header, 0, len(v.Headers))
		for key, value := range v.Headers {
			if value == nil {
				return nil, errors.New("empty HTTP header value: ", key).AtError()
			}
			config.Header = append(config.Header, &http.Header{
				Name:  key,
				Value: *value,
			})
		}
	}
	return config, nil
}

type TCPHeaderConfig struct {
	Type     string           `json:"type"`
	Request  *json.RawMessage `json:"request"`
	Response *json.RawMessage `json:"response"`
}

func (c *TCPHeaderConfig) Build() (proto.Message, error) {
	switch strings.ToLower(c.Type) {
	case "", "none":
		return new(noop.Config), nil
	case "http":
		config := new(http.Config)
		if c.Request != nil {
			r := new(AuthenticatorRequest)
			if err := json.Unmarshal(*c.Request, r); err != nil {
				return nil, errors.New("invalid HTTP request header").Base(err)
			}
			req, err := r.Build()
			if err != nil {
				return nil, err
			}
			config.Request = req
		}
		if c.Response != nil {
			r := new(AuthenticatorResponse)
			if err := json.Unmarshal(*c.Response, r); err != nil {
				return nil, errors.New("invalid HTTP response header").Base(err)
			}
			resp, err := r.Build()
			if err != nil {
				return nil, err
			}
			config.Response = resp
		}
		return config, nil
	default:
		return nil, errors.New("unknown TCP header type: ", c.Type)
	}
}

type TCPConfig struct {
	HeaderConfig        *TCPHeaderConfig `json:"header"`
	AcceptProxyProtocol bool             `json:"acceptProxyProtocol"`
}

func (c *TCPConfig) Build() (proto.Message, error) {
	config := new(tcp.Config)
	if c.HeaderConfig != nil {
		header, err := c.HeaderConfig.Build()
		if err != nil {
			return nil, err
		}
		config.HeaderSettings = serial.ToTypedMessage(header)
	}
	config.AcceptProxyProtocol = c.AcceptProxyProtocol
	return config, nil
}

type SplitHTTPConfig struct{}

func (c *SplitHTTPConfig) Build() (proto.Message, error) {
	return nil, errors.New("splithttp/xhttp transport removed in minimal build")
}

type KCPConfig struct{}

func (c *KCPConfig) Build() (proto.Message, error) {
	return nil, errors.New("kcp/mkcp transport removed in minimal build")
}

type GRPCConfig struct{}

func (c *GRPCConfig) Build() (proto.Message, error) {
	return nil, errors.New("grpc transport removed in minimal build")
}

type WebSocketConfig struct {
	Path                 string            `json:"path"`
	Headers              map[string]string `json:"headers"`
	AcceptProxyProtocol  bool              `json:"acceptProxyProtocol"`
	MaxEarlyData         int32             `json:"maxEarlyData"`
	EarlyDataHeaderName  string            `json:"earlyDataHeaderName"`
	UseBrowserForwarding bool              `json:"useBrowserForwarding"`
	Host                 string            `json:"host"`
}

func (c *WebSocketConfig) Build() (proto.Message, error) {
	path := c.Path
	if u, err := url.Parse(path); err == nil {
		if q := u.Query(); q.Get("ed") != "" {
			q.Del("ed")
			u.RawQuery = q.Encode()
			path = u.String()
		}
	}
	config := &websocket.Config{
		Path:                path,
		Header:              c.Headers,
		AcceptProxyProtocol: c.AcceptProxyProtocol,
		Host:                c.Host,
	}
	return config, nil
}

type HttpUpgradeConfig struct{}

func (c *HttpUpgradeConfig) Build() (proto.Message, error) {
	return nil, errors.New("httpupgrade transport removed in minimal build")
}

type HysteriaConfig struct{}

func (c *HysteriaConfig) Build() (proto.Message, error) {
	return nil, errors.New("hysteria transport removed in minimal build")
}

