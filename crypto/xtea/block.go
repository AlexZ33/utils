package xtea

// XTEA is based on 64 rounds.
const numRounds = 64

// blockToUint32 读取8 byte 边出 2个 uint32
func blockToUint32(src []byte) (uint32, uint32) {
	r0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	r1 := uint32(src[4])<<24 | uint32(src[5])<<16 | uint32(src[6])<<8 | uint32(src[7])
	return r0, r1
}

//uint32ToBlock 2个 uint32 边出 8 bytes
func uint32ToBlock(v0, v1 uint32, dst []byte) {
	dst[0] = byte(v0 >> 24)
	dst[1] = byte(v0 >> 16)
	dst[2] = byte(v0 >> 8)
	dst[3] = byte(v0)
	dst[4] = byte(v1 >> 24)
	dst[5] = byte(v1 >> 16)
	dst[6] = byte(v1 >> 8)
	dst[7] = byte(v1 >> 0)
}

// encryptBlock 使用XTEA加密单个8 byte block
func encryptBlock(c *Cipher, dst, src []byte) {
	v0, v1 := blockToUint32(src)

	// Two rounds of XTEA applied per loop
	for i := 0; i < numRounds; {
		v0 += ((v1<<4 ^ v1>>5) + v1) ^ c.table[i]
		i++
		v1 += ((v0<<4 ^ v0>>5) + v0) ^ c.table[i]
		i++
	}

	uint32ToBlock(v0, v1, dst)
}

// decryptBlock 使用XTEA解密单个8 byte block
func decryptBlock(c *Cipher, dst, src []byte) {
	v0, v1 := blockToUint32(src)

	// Two rounds of XTEA applied per loop
	for i := numRounds; i > 0; {
		i--
		v1 -= ((v0<<4 ^ v0>>5) + v0) ^ c.table[i]
		i--
		v0 -= ((v1<<4 ^ v1>>5) + v1) ^ c.table[i]
	}

	uint32ToBlock(v0, v1, dst)
}
