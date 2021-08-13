package view_models

type LoginVm struct {
	Token                 string `json:"token"`
	TokenExpiredAt        int64  `json:"token_expired_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiredAt int64  `json:"refresh_token_expired_at"`
}
