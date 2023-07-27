package authenticationService

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

type FirebaseAuthenticationService struct {
	context     context.Context
	firebaseApp *firebase.App
}

func NewFirebaseAuthenticationService(context context.Context, firebaseApp *firebase.App) *FirebaseAuthenticationService {
	return &FirebaseAuthenticationService{
		context:     context,
		firebaseApp: firebaseApp,
	}
}

func (fas *FirebaseAuthenticationService) IsValidToken(token string) bool {
	client, err := fas.firebaseApp.Auth(fas.context)
	var result bool

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	firebaseToken, err := client.VerifyIDToken(fas.context, token)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		result = false
	}

	if firebaseToken != nil {
		result = true
	}

	return result
}
