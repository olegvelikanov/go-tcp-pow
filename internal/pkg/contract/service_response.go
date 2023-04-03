package contract

import "encoding/binary"

// | offset       | description | length
// |--------------|-------------|-----------
// |            0 | quoteLen    | 8 bytes
// |            8 | quote       | quoteLen
// |   8+quoteLen |             |

const (
	srvRspQuoteLenOffset = 0
	srvRspQuoteOffset    = 8
)

type ServiceResponse struct {
	Quote []byte
}

func (*ServiceResponse) isMessage() {}

func serializeServiceResponse(m *ServiceResponse, b []byte) (int, error) {
	if len(b) < srvRspQuoteOffset+len(m.Quote) {
		return 0, ErrNotEnoughBytes
	}
	binary.LittleEndian.PutUint64(b[srvRspQuoteLenOffset:srvRspQuoteOffset], uint64(len(m.Quote)))

	copy(b[srvRspQuoteOffset:], m.Quote)

	return srvRspQuoteOffset + len(m.Quote), nil
}

func deserializeServiceResponse(b []byte) (*ServiceResponse, error) {
	if len(b) < srvRspQuoteOffset {
		return nil, ErrNotEnoughBytes
	}

	quoteLen := int(binary.LittleEndian.Uint64(b[srvRspQuoteLenOffset:srvRspQuoteOffset]))

	if len(b) < srvRspQuoteOffset+quoteLen {
		return nil, ErrNotEnoughBytes
	}

	return &ServiceResponse{
		Quote: b[srvRspQuoteOffset : srvRspQuoteOffset+quoteLen],
	}, nil

}
