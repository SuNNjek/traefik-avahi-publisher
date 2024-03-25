package avahi

import (
	"github.com/godbus/dbus/v5"
	"github.com/holoplot/go-avahi"
	"github.com/miekg/dns"
)

const (
	DnsClassIn   = uint16(0x01)
	DnsTypeCname = uint16(0x05)
)

type AvahiClient struct {
	server *avahi.Server
}

func NewAvahiClient(dbusConn *dbus.Conn) (*AvahiClient, error) {
	server, err := avahi.ServerNew(dbusConn)
	if err != nil {
		return nil, err
	}

	return &AvahiClient{
		server,
	}, nil
}

func (c *AvahiClient) Close() {
	c.server.Close()
}

func (c *AvahiClient) GetHostNameFqdn() (string, error) {
	return c.server.GetHostNameFqdn()
}

func (c *AvahiClient) PublishCnames(cnames []string) error {
	avahiFqdn, err := c.GetHostNameFqdn()
	if err != nil {
		return err
	}

	eg, err := c.server.EntryGroupNew()
	if err != nil {
		return err
	}

	fqdn := dns.Fqdn(avahiFqdn)

	rdata := make([]byte, len(fqdn)+1)
	_, err = dns.PackDomainName(fqdn, rdata, 0, nil, false)
	if err != nil {
		return err
	}

	for _, cname := range cnames {
		err = eg.AddRecord(
			avahi.InterfaceUnspec,
			avahi.ProtoUnspec,
			uint32(0),
			cname,
			DnsClassIn,
			DnsTypeCname,
			60,
			rdata,
		)

		if err != nil {
			return err
		}
	}

	return eg.Commit()
}
