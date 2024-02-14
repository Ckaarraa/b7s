import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

type HashAndData struct {
	Hash string      
	Data interface{} 
}

func getDigest(rec interface{}) (HashAndData, error) {
	payload, err := json.Marshal(rec)
	if err != nil {
		return HashAndData{}, err 
	}

	hash := sha256.Sum256(payload)
	hashHex := hex.EncodeToString(hash[:])
	return HashAndData{Hash: hashHex, Data: rec}, nil
}
