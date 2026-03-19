package call

import (
	"fmt"
	"strings"
)

func NewDialParams(
	variables, customHeaders map[string]string,
	destination string,
	sipTrunk SIPTrunk,
) *DialParams {
	return &DialParams{
		Variables:     variables,
		CustomHeaders: customHeaders,
		Destination:   destination,
		SIPTrunk:      sipTrunk,
	}
}

type DialParams struct {
	Variables     map[string]string `json:"variables"`
	CustomHeaders map[string]string `json:"headers"`
	Destination   string            `json:"destination"`
	SIPTrunk      SIPTrunk          `json:"sip_trunk"`
}

func (d *DialParams) String() string {
	vars := make([]string, 0, len(d.Variables))
	for key, value := range d.Variables {
		vars = append(vars, fmt.Sprintf("%s=%s", key, value))
	}

	for key, value := range d.CustomHeaders {
		vars = append(vars, fmt.Sprintf("sip_h_X-%s=%s", key, value))
	}

	if d.SIPTrunk.Auth != nil {
		vars = append(vars, fmt.Sprintf("sip_auth_username=%s", d.SIPTrunk.Auth.Username))
		vars = append(vars, fmt.Sprintf("sip_auth_password=%s", d.SIPTrunk.Auth.Password))
	}

	varsString := strings.Join(vars, ",")

	return fmt.Sprintf(
		"[%s]sofia/external/%s@%s:%s;transport=%s",
		varsString,
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
