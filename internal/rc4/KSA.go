package rc4

// KSA
// Key Schedule Algorithm
func KSA(key []byte) []byte {
	length := len(key)

	// length must be 1 < length < 256
	if length < 1 || length > 256 {
		return nil
	}

	S := make([]byte, 256)

	for i := 0; i < 256; i++ {
		S[i] = byte(i)
	}

	j := uint8(0)

	for i := 0; i < 256; i++ {
		j += S[i] + key[i%length]
		S[i], S[j] = S[j], S[i]
	}

	return S
}
