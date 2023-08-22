package authenticationService

type IAuthenticationService interface {
	IsValidToken(token string) bool
}
