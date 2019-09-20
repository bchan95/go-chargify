package chargify

import (
	"crypto/sha1"
	"fmt"
	"log"
)

func (c *client) GenerateSelfServiceLink(method string, subscriptionID int64) string {
	if c.siteSharedKey == "" {
		return ""
	}
	toHash := fmt.Sprintf("%s--%d--%s", method, subscriptionID, c.siteSharedKey)
	log.Printf(toHash)
	hash := sha1.Sum([]byte(toHash))
	return fmt.Sprintf("%x", hash)
}
