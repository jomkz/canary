package engine

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HoverDomain struct {
	Id string
}

const (
	HOVER_API_ENDPOINT = "https://www.hover.com/api/control_panel"
)

// HoverDomainList will return the list of domains from Hover.
func HoverDomainList() ([]HoverDomain, error) {
	log.Trace("list hover domains")
	domains := make([]HoverDomain, 0)
	resp, err := http.Get(fmt.Sprintf("%s/dns", HOVER_API_ENDPOINT))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Tracef("domains: %o", body)

	return domains, nil
}
