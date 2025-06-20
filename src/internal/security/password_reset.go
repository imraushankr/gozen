package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// GeneratePasswordResetToken generates a password reset token
func GeneratePasswordResetToken() (string, string, time.Time, error) {
	// Generate random bytes
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", time.Time{}, err
	}

	// Convert to hex string (this is what we send to user)
	resetToken := hex.EncodeToString(bytes)

	// Hash the token (this is what we store in database)
	hash := sha256.Sum256([]byte(resetToken))
	hashedToken := hex.EncodeToString(hash[:])

	// Set expiration time (10 minutes from now)
	expireTime := time.Now().Add(10 * time.Minute)

	return resetToken, hashedToken, expireTime, nil
}

// HashResetToken hashes a reset token for comparison
func HashResetToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}