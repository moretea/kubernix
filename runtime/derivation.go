package runtime

type Derivation struct {
	Id    string
	Url   string
	Size  uint64
	State string
}

const DERIVATION_STATE_PULLING = "PULLING"
const DERIVATION_STATE_READY = "READY"
const DERIVATION_STATE_ERROR = "ERROR"
