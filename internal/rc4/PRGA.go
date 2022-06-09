package rc4

func PRGA(S []byte, text []byte) []byte {
	i := 0
	j := 0
	enc := make([]byte, len(text))

	for index, value := range text {
		i = (i + 1) % 256
		j = (j + int(S[i])) % 256

		S[j], S[i] = S[i], S[j]

		K := S[(S[i]+S[j])%255]

		enc[index] = K ^ value
	}

	return enc
}
