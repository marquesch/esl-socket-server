package call

import (
	"encoding/json"
	"errors"
)

type SIPTransport string

func (t *SIPTransport) isValid() bool {
	switch *t {
	case SIPTransportTCP, SIPTransportUDP:
		return true
	}

	return false
}

func (t *SIPTransport) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v := SIPTransport(s)
	if !v.isValid() {
		return errors.New("invalid sip transport")
	}

	*t = v

	return nil
}

const (
	SIPTransportTCP SIPTransport = "tcp"
	SIPTransportUDP SIPTransport = "udp"
)

func NewSIPAuth(username, password string) *SIPAuth {
	return &SIPAuth{Username: username, Password: password}
}

type SIPAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewSIPTrunk(host, port string, transport SIPTransport, auth *SIPAuth) (*SIPTrunk, error) {
	if !transport.isValid() {
		return nil, errors.New("invalid sip transport")
	}

	return &SIPTrunk{
		Host:      host,
		Port:      port,
		Auth:      auth,
		Transport: transport,
	}, nil
}

type SIPTrunk struct {
	Host      string       `json:"host"`
	Port      string       `json:"port"`
	Auth      *SIPAuth     `json:"auth"`
	Transport SIPTransport `json:"transport"`
}
