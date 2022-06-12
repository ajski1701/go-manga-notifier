package models

type AuthOutput struct {
	Result string
	Token  authTokens
}

type authTokens struct {
	Session string
	Refresh string
}
