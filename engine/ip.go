package engine

import (
	"io"
	"net"
	"net/http"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

const (
	PUBLIC_IP_PROVIDER = "https://icanhazip.com"
)

// PublicAddress will return the public IP address for this host.
func PublicAddress() (string, error) {
	resp, err := http.Get(PUBLIC_IP_PROVIDER)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := strings.TrimFunc(string(body), func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsPunct(r)
	})

	log.Debugf("public address: %s", ip)
	return ip, nil
}

// FQDNAddress will return the IP address for the given FQDN.
func FQDNAddress(fqdn string) (string, error) {
	ips, err := net.LookupIP(fqdn)
	if err != nil {
		return "", err
	}

	ip := ""
	if len(ips) > 0 {
		ip = ips[0].String()
	}

	log.Debugf("FQDN address: %s", ip)
	return ip, nil
}
