package main

import (
	"encoding/hex"
	"fmt"
	"github.com/akarregi/go-rc4/internal/rc4"
	"strings"
)

func main() {
	keystream := rc4.KSA([]byte("Secret"))
	fmt.Println(strings.ToUpper(hex.EncodeToString(keystream)))

	fmt.Println()

	encrypted := rc4.PRGA(keystream, []byte("Attack at dawn"))
	fmt.Println(strings.ToUpper(hex.EncodeToString(encrypted)))
}
