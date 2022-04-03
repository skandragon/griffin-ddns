/*
 * Copyright 2022 Michael Graff.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dns

// Header defines the DNS header.
type Header struct {
	ID      uint16
	QR      bool
	Opcode  Opcode
	AA      bool
	TC      bool
	RD      bool
	RA      bool
	RCode   ResponseCode
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

// Opcode is the requested operation for the server to perform.
type Opcode byte

const (
	// Query is a normal query [RFC1035]
	Query Opcode = iota
	// IQuery is depricated [RFC3425]
	IQuery
	// Status isn't really used either [RFC1035]
	Status
	_
	// Notify indicates zone changes to secondaries [RFC1996]
	Notify
	// Update DDNS update request [RFC2136]
	Update
	// DNSStatefulOperations [RFC8490]
	DNSStatefulOperations
)

// ResponseCode holds the various types of respose codes.
type ResponseCode byte

const (
	// NoError indicates no error occurred.
	NoError ResponseCode = iota
	// FormError means the server rejected the contents of our query.
	FormError
	// ServFail indicates the server failed in some way.
	ServFail
	// NXDomain means the name we are asking for does not exist.
	NXDomain
	// NotImplemented The name server does not support the specified Opcode.
	NotImplemented
	// Refused means the server will not answer our query.
	Refused
	// YXDomain Some name that ought not to exist, does exist.
	YXDomain
	// YXRRSet Some RRset that ought not to exist, does exist.
	YXRRSet
	// NSRRSet means the name exists, but the specific record type doesn't.
	NSRRSet
	// NotAuthorized The server is not authoritative for the zone named in the Zone Section
	NotAuthorized
	// NotZone A name used in the Prerequisite or Update Section is
	// not within the zone denoted by the Zone Section
	NotZone
)

// Class if the zone class type.
type Class uint16

const (
	// Internet zones [RFC1035]
	Internet Class = 1
	// Chaos is used mostly for other purposes now
	Chaos = 2
	// None [RFC2136]
	None = 254
	// Any [RFC1035]
	Any = 255
)

// RRType holds the resource record types
type RRType uint16

const (
	// A holds IPv6 host addresses [RFC1035]
	A RRType = 1
	// NS - name service [RFC1035]
	NS = 2
	// MD - Mail destination (obsolete, use MX) [RFC1035]
	MD = 3
	// MF - Mail forwarder (obsolete, use MX) [RFC1035]
	MF         = 4
	CNAME      = 5
	SOA        = 6
	MB         = 7
	MG         = 8
	MR         = 9
	NULL       = 10
	WKS        = 11
	PTR        = 12
	HINFO      = 13
	MINFO      = 14
	MX         = 15
	TX         = 16
	RP         = 17
	AFSDB      = 18
	X25        = 19
	ISDN       = 20
	RT         = 21
	NSAP       = 22
	NSAPPTR    = 23
	SIG        = 24
	KEY        = 25
	PX         = 26
	GPOS       = 27
	AAAA       = 28
	LOC        = 29
	NXT        = 30
	EID        = 31
	NIMLOC     = 32
	SRV        = 33
	ATMA       = 34
	NAPTR      = 35
	KX         = 36
	CERT       = 37
	A6         = 38
	DNAME      = 39
	SINK       = 40
	OPT        = 41
	APL        = 42
	DS         = 43
	SSHFP      = 44
	IPSECKEY   = 45
	RRSIG      = 46
	NSEC       = 47
	DNSKEY     = 48
	DHCID      = 49
	NSEC3      = 50
	NSEC3PARAM = 51
	TLSA       = 52
	SMIMEA     = 53
	HIP        = 55
	NINFO      = 56
	RKEY       = 57
	TALINK     = 58
	CDS        = 59
	CDNSKEY    = 60
	OPENPGPKEY = 61
	CSYNC      = 62
	ZONEMD     = 63
	SVCB       = 64
	HTTPS      = 65
	SPF        = 99
	UINFO      = 100
	UID        = 101
	GID        = 102
	UNSPEC     = 103
	NID        = 104
	L32        = 105
	L64        = 106
	LP         = 107
	EUI48      = 108
	EUI64      = 109
	TKEY       = 249
	TSIG       = 250
	IXFR       = 251
	AXFR       = 252
	MAILB      = 253
	MAILA      = 254
	ANY        = 255
	URI        = 256
	CCA        = 257
	AVC        = 258
	DOA        = 259
	AMTRELAY   = 260
	TA         = 32768
	DLV        = 32769
)

// HeaderMarshal will return a byte slide containing the binary representation of the header.
func HeaderMarshal(h *Header) ([]byte, error) {
	ret := make([]byte, 12)
	ret[0] = byte(h.ID & 0xff00 >> 8)
	ret[1] = byte(h.ID & 0x00ff)
	if h.QR {
		ret[2] |= 0x80
	}
	ret[2] |= (byte(h.Opcode) & 0x0f << 3)
	if h.AA {
		ret[2] |= 0x04
	}
	if h.TC {
		ret[2] |= 0x02
	}
	if h.RD {
		ret[2] |= 0x01
	}
	if h.RA {
		ret[3] |= 0x80
	}
	ret[3] |= byte(h.RCode & 0x0f)

	ret[4] = byte(h.QDCount & 0xff00 >> 8)
	ret[5] = byte(h.QDCount & 0x00ff)

	ret[6] = byte(h.ANCount & 0xff00 >> 8)
	ret[7] = byte(h.ANCount & 0x00ff)

	ret[8] = byte(h.NSCount & 0xff00 >> 8)
	ret[9] = byte(h.NSCount & 0x00ff)

	ret[10] = byte(h.ARCount & 0xff00 >> 8)
	ret[11] = byte(h.ARCount & 0x00ff)

	return ret, nil
}

// HeaderUnmarshal will unmarshal the DNS header, and return the number of bytes consumed.
func HeaderUnmarshal(h *Header, data []byte) (int, error) {
	h.ID = uint16(data[0])<<8 + uint16(data[1])
	h.QR = data[2]&0x80 != 0
	h.Opcode = Opcode(data[2] & 0x78 >> 3)
	h.AA = data[2]&0x04 != 0
	h.TC = data[2]&0x02 != 0
	h.RD = data[2]&0x01 != 0
	h.RA = data[3]&0x80 != 0
	h.RCode = ResponseCode(data[3] & 0x0f)
	h.QDCount = uint16(data[4])<<8 + uint16(data[5])
	h.ANCount = uint16(data[6])<<8 + uint16(data[7])
	h.NSCount = uint16(data[8])<<8 + uint16(data[9])
	h.ARCount = uint16(data[10])<<8 + uint16(data[11])

	return 12, nil
}
