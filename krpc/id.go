package krpc

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/anacrolix/torrent/bencode"

	"github.com/anacrolix/dht/v2/int160"
)

func RandomNodeID() (id ID) {
	crand.Read(id[:])
	return
}

type ID [20]byte

var (
	_ interface {
		bencode.Unmarshaler
	} = (*ID)(nil)
	_ bencode.Marshaler = ID{}
)

func IdFromString(s string) (id ID) {
	if n := copy(id[:], s); n != 20 {
		panic(n)
	}
	return
}

func (id ID) MarshalBencode() ([]byte, error) {
	return []byte("20:" + string(id[:])), nil
}

func (id *ID) UnmarshalBencode(b []byte) error {
	var s string
	if err := bencode.Unmarshal(b, &s); err != nil {
		return err
	}
	if n := copy(id[:], s); n != 20 {
		return fmt.Errorf("string has wrong length: %d", n)
	}
	return nil
}

func (id ID) String() string {
	return hex.EncodeToString(id[:])
}

func (id ID) Int160() int160.T {
	return int160.FromByteArray(id)
}
