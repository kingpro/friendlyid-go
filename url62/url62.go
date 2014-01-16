package url62

import (
	"github.com/nu7hatch/gouuid"
	"math/big"
)

const base62alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func FromUUID(shex string) string {
	u, _ := uuid.ParseHex(shex)

	convertUp(u, base62alphabet)
}

// ConvertUp converts a hexadecimal UUID string to a base alphabet greater than
// 16. It is used here to compress a 32 character UUID down to 23 URL friendly
// characters.
func convertUp(oldNumber string, baseAlphabet string) string {
	n := big.NewInt(0)
	n.SetString(oldNumber, 16)

	base := big.NewInt(int64(len(baseAlphabet)))

	newNumber := make([]byte, 23) //converted size of max base-62 uuid
	i := len(newNumber)

	for n.Int64() != 0 {
		i--
		_, r := n.DivMod(n, base, big.NewInt(0))
		newNumber[i] = baseAlphabet[r.Int64()]
	}
	return string(newNumber[i:])
}
