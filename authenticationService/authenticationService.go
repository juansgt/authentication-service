package authenticationService

type IAuthenticationService interface {
	IsValidToken() bool
}
