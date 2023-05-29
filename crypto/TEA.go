package crypto

//TEA算法使用64位的明文分组和128位的密钥,它使用Feistel分组加密框架,需要进行 64 轮迭代,尽管作者认为 32 轮已经足够了.
//该算法使用了一个神秘常数δ作为倍数,它来源于黄金比率,以保证每一轮加密都不相同.但δ的精确值似乎并不重要,
//这里 TEA 把它定义为 δ=「(√5 - 1)231」(也就是程序中的 0×9E3779B9).

//TEA加密和解密时都使用一个常量值,这个常量值为0x9e3779b, 这个值是近似黄金分割率,注意,有些编程人员为了避免在程序中直接出现”mov变量,0x9e3779b”,
// 以免被破解者直接搜索0x9e3779b这个常数得知使用TEA算法,所以有时会使用”sub变量,0x61C88647”代替”mov变量,0x9e3779b”,0x61C88647=－(0x9e3779b).
//TEA算法每一次可以操作64bit(8byte),采用128bit(16byte)作为key,算法采用迭代的形式,推荐的迭代轮数是64轮,最少32轮.

import (
	"crypto/cipher"
	"encoding/binary"
	"errors"
)

const (
	// BlockSize TEA BlockSize 单位byte
	BlockSize = 8

	// KeySize TEA 算法 Key的byte长度.
	KeySize = 16

	// delta is the TEA key schedule 常量.
	delta = 0x9e3779b9

	// numRounds TEA算法 round 标准常量.
	numRounds = 64
)

//tea TEA加密算法.
type tea struct {
	key    [16]byte
	rounds int
}

// NewCipher 加密算法构造器,使用标准rounds, key 长度必须是 16 byte
func NewCipher(key []byte) (cipher.Block, error) {
	return NewCipherWithRounds(key, numRounds)
}

// NewCipherWithRounds 加密算法构造器,key 长度必须是 16 byte
func NewCipherWithRounds(key []byte, rounds int) (cipher.Block, error) {
	if len(key) != 16 {
		return nil, errors.New("tea: incorrect key size")
	}

	if rounds&1 != 0 {
		return nil, errors.New("tea: odd number of rounds specified")
	}

	c := &tea{
		rounds: rounds,
	}
	copy(c.key[:], key)

	return c, nil
}

//BlockSize 返回TEA block size,结果为常量. 方法来满足package "crypto/cipher" 的 Block interface
func (*tea) BlockSize() int {
	return BlockSize
}

//Encrypt 使用t.key 来加密 src参数8byte buffer内容,密文保存在dst里面.
// 注意data的长度大于block, 在连续的block上调用encrypt是不安全的,应该使用 CBC crypto/cipher/cbc.go 那种方式来encrypt
func (t *tea) Encrypt(dst, src []byte) {
	e := binary.BigEndian
	v0, v1 := e.Uint32(src), e.Uint32(src[4:])
	k0, k1, k2, k3 := e.Uint32(t.key[0:]), e.Uint32(t.key[4:]), e.Uint32(t.key[8:]), e.Uint32(t.key[12:])

	sum := uint32(0)
	delta := uint32(delta)

	for i := 0; i < t.rounds/2; i++ {
		sum += delta
		v0 += ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		v1 += ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
	}

	e.PutUint32(dst, v0)
	e.PutUint32(dst[4:], v1)
}

// Decrypt 使用t.key 来解密 src参数8byte buffer内容,明文保存在dst里面.
func (t *tea) Decrypt(dst, src []byte) {
	e := binary.BigEndian
	v0, v1 := e.Uint32(src), e.Uint32(src[4:])
	k0, k1, k2, k3 := e.Uint32(t.key[0:]), e.Uint32(t.key[4:]), e.Uint32(t.key[8:]), e.Uint32(t.key[12:])

	delta := uint32(delta)
	sum := delta * uint32(t.rounds/2) // in general, sum = delta * n

	for i := 0; i < t.rounds/2; i++ {
		v1 -= ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
		v0 -= ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		sum -= delta
	}

	e.PutUint32(dst, v0)
	e.PutUint32(dst[4:], v1)
}
