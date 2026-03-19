package call

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Transport string

func (t *Transport) isValid() bool {
	switch *t {
	case TransportTCP, TransportUDP:
		return true
	}

	return false
}

func (t *Transport) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v := Transport(s)
	if !v.isValid() {
		return errors.New("invalid transport")
	}

	*t = v

	return nil
}

const (
	TransportTCP Transport = "tcp"
	TransportUDP Transport = "udp"
)

type SIPAuth struct {
	username string
	password string
}

type SIPTrunk struct {
	Host      string
	Port      string
	Auth      *SIPAuth
	Transport Transport
}

type DialParams struct {
	Variables     map[string]string `json:"variables"`
	CustomHeaders map[string]string `json:"headers"`
	SofiaProfile  string
	Destination   string
	SIPTrunk      SIPTrunk
}

func (d *DialParams) String() string {
	var varsString string
	for varKey := range d.Variables {
		varsString = fmt.Sprintf("%s,%s=%s", varsString, varKey, d.Variables[varKey])
	}

	for headerKey := range d.CustomHeaders {
		varsString = fmt.Sprintf(
			"%s,sip_h_X-%s=%s",
			varsString,
			headerKey,
			d.CustomHeaders[headerKey],
		)
	}

	if d.SIPTrunk.Auth != nil {
		varsString = fmt.Sprintf(
			"%s,sip_auth_username=%s,sip_auth_password=%s",
			varsString,
			d.SIPTrunk.Auth.username,
			d.SIPTrunk.Auth.password,
		)
	}

	return fmt.Sprintf(
		"[%s]sofia/%s/%s@%s:%s;transport=%s",
		varsString,
		d.SofiaProfile,
		d.Destination,
		d.SIPTrunk.Host,
		d.SIPTrunk.Port,
		d.SIPTrunk.Transport,
	)
}

type Call struct {
	DialParams     DialParams `json:"dial_params"`
	TransferParams DialParams `json:"transfer_params"`
}
