package dns

// Header defines the DNS header.
type Header struct {
	ID      uint16
	QR      bool
	Opcode  byte
	AA      bool
	TC      bool
	RD      bool
	RA      bool
	RCode   byte
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

type ResponseCode byte

const (
	NoError ResponseCode = iota
	FormError
	ServFail
	NXDomain
	NotImplemented
	Refused
	YXDomain
	YXRRSet
	NSRRSet
	NotAuthorized
	NotZone
)

// HeaderMarshal will return a byte slide containing the binary representation of the header.
func HeaderMarshal(h *Header) ([]byte, error) {
	ret := make([]byte, 12)
	ret[0] = byte(h.ID & 0xff00 >> 8)
	ret[1] = byte(h.ID & 0x00ff)
	if h.QR {
		ret[2] |= 0x80
	}
	ret[2] |= (h.Opcode & 0x0f << 3)
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
	ret[3] |= (h.RCode & 0x0f)

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
	h.Opcode = data[2] & 0x78 >> 3
	h.AA = data[2]&0x04 != 0
	h.TC = data[2]&0x02 != 0
	h.RD = data[2]&0x01 != 0
	h.RA = data[3]&0x80 != 0
	h.RCode = data[3] & 0x0f
	h.QDCount = uint16(data[4])<<8 + uint16(data[5])
	h.ANCount = uint16(data[6])<<8 + uint16(data[7])
	h.NSCount = uint16(data[8])<<8 + uint16(data[9])
	h.ARCount = uint16(data[10])<<8 + uint16(data[11])

	return 12, nil
}
