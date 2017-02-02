package xmpp

import "encoding/xml"

type ServiceDiscovery struct {
	XMLName  xml.Name `xml:"http://jabber.org/protocol/disco#info query"`
	InnerXML string   `xml:",innerxml"`
}

func (*ServiceDiscovery) IsIQPayload() {
}

type Blocklist struct {
	XMLName  xml.Name `xml:"urn:xmpp:blocking blocklist"`
	InnerXML string   `xml:",innerxml"`
}

func (*Blocklist) IsIQPayload() {
}

type Block struct {
	XMLName  xml.Name `xml:"urn:xmpp:blocking block"`
	InnerXML string   `xml:",innerxml"`
}

func (*Block) IsIQPayload() {
}

type Unblock struct {
	XMLName  xml.Name `xml:"urn:xmpp:blocking unblock"`
	InnerXML string   `xml:",innerxml"`
}

func (*Unblock) IsIQPayload() {
}
