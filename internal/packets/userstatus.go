package packets

import (
	"bytes"
)

type NPS_CUSTOMERID uint32

const NPS_SESSION_KEY_LEN = 32

type UserStatus struct {
	// User status
	isBan bool
	isGagged bool
	customerId NPS_CUSTOMERID
	 sessionKeyStr [NPS_SESSION_KEY_LEN]byte
	 sessionKeyLen int
	 sessionKey bytes.Buffer
}

