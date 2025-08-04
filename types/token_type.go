package types

type TokenType string

func (t TokenType) String() string {
	return string(t)
}

// TokenType values.
const (
	TokenTypeAccessToken  TokenType = "access_token"
	TokenTypeRefreshToken TokenType = "refresh_token"
)

type TokenPayload struct {
	UserID string    `json:"sub"`
	JTI    string    `json:"jti"`
	Type   TokenType `json:"type"`
	IAT    int64     `json:"iat"`
	Exp    int64     `json:"exp"`
}
