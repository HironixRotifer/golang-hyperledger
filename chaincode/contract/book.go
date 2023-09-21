package contract

import (
	"fmt"
	"time"
)

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

// GenKeyID returns a unique ID
func GenKeyID(key string) string {
	now := time.Now()
	assetId := fmt.Sprintf(key, now.Unix()*1e3+int64(now.Nanosecond())/1e6)
	return assetId
}
