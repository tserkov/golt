package messages

const (
	SignatureInit       = 0x01
	SignatureAckFailure = 0x0E
	SignatureReset      = 0x0F
	SignatureRun        = 0x10
	SignatureDiscardAll = 0x2F
	SignaturePullAll    = 0x3F
	SignatureSuccess    = 0x70
	SignatureRecord     = 0x71
	SignatureIgnored    = 0x7E
	SignatureFailure    = 0x7F
)
