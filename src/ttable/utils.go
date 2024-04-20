package ttable

import "encoding/binary"

// Apply sbox0 to each byte in w.
func subw(w uint32) uint32 {
	return uint32(sbox[w>>24])<<24 |
		uint32(sbox[w>>16&0xff])<<16 |
		uint32(sbox[w>>8&0xff])<<8 |
		uint32(sbox[w&0xff])
}

// Rotate
func rotw(w uint32) uint32 { return w<<8 | w>>24 }

func KeyExpansion(key []byte, nr int) []uint32 {
	nk := len(key) / 4
	xk := make([]uint32, nk*(nr+1))

	var i int

	// Copy key to xk
	for i = 0; i < nk; i++ {
		xk[i] = binary.BigEndian.Uint32(key[4*i:])
	}

	for ; i < len(xk); i++ {
		t := xk[i-1]
		if i%nk == 0 {
			t = subw(rotw(t)) ^ (uint32(powx[i/nk-1]) << 24)
		} else if nk > 6 && i%nk == 4 {
			t = subw(t)
		}
		xk[i] = xk[i-nk] ^ t
	}

	return xk
}
