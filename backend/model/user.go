
package model

import(
	"time"
	"github.com/golang-jwt/jwt/v5"


)

type User struct {
	Id          int64  `json:"id"`
	Email	   string `json:"email"`
	Password   string `json:"-"`
}

// TokenClaims defines the claims for our JWT
type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// RefreshToken represents a refresh token stored in the database
type RefreshToken struct {
	Token     string    `json:"token"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type NewTokenRequest struct {
	Token_val string `json:"token_val"`
	Expiry string `json:"expiry"`
	Type_auth string `json:"type_auth"` // e.g., "access", "refresh"
	Scopes  []string `json:"scopes"` // e.g., ["read", "write"]
	}

// Credentials for login requests
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OauthToken struct{
	UserID   string `json:"user_id"`
	TokenValue string `json:"token_value"`
	Email    string `json:"email"`
	Provider string `json:"provider"` // e.g., "google", "github"
	ExpiresAt time.Time `json:"expires_at"`
	Scopes   []string `json:"scopes"` // e.g., ["read", "write"]
}

type Client struct {
	Provider	 string   `json:"provider"` // e.g., "google", "github"
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURIs []string `json:"redirect_uris"`
}
