package golinq

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Hash(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	hash := sha256.New()
	hash.Write(bytes)

	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}
