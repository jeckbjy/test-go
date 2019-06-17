package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"
	guid "github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"github.com/rs/xid"
	sshortid "github.com/skahack/go-shortid"
	tshortid "github.com/teris-io/shortid"
	"math/big"
)

func main() {
	uuid4, _ := uuid.NewV4()

	fmt.Printf("teris id   :%s\n", tshortid.MustGenerate())
	fmt.Printf("skahack id :%s\n", sshortid.Generator().Generate())
	fmt.Printf("xid        :%s\n", xid.New().String())
	fmt.Printf("shortuuid  :%s\n", shortuuid.New())
	fmt.Printf("gofrs_uuid :%s\n", uuid4.String())
	fmt.Printf("google_uuid:%s\n", guid.New().String())

	str := []byte("example_Encoding_DecodeString")
	data32 := base32.StdEncoding.EncodeToString(str)
	fmt.Printf("base32:%+v,%s\n",len(data32), data32)
	data64 := base64.StdEncoding.EncodeToString(str)
	fmt.Printf("base64:%+v,%s\n",len(data64), data64)

	alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	//alphabet := "23456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"
	datan := BaseN(alphabet, str)
	fmt.Printf("base58:%+v,%s\n",len(datan), datan)

	num := make([]byte, 8)
	binary.LittleEndian.PutUint64(num,12345678900000)
	num32 := base32.StdEncoding.EncodeToString(num)
	fmt.Printf("num32:%s\n", num32)
	num64 := base64.StdEncoding.EncodeToString(num)
	fmt.Printf("num64:%s\n", num64)
}

// base n虽然可以实现，但是效率貌似并不高,SetBytes会多一次拷贝,DivMod也会涉及到除法运算，
// 肯定不如base64的位运算快
func BaseN(alphabet string, input []byte) string {
	zero := big.NewInt(0)
	radix := big.NewInt(int64(len(alphabet)))

	x := new(big.Int)
	x.SetBytes(input)
	answer := make([]byte, 0, len(input))
	for x.Cmp(zero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, radix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// padding?

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
