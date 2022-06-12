package models

type AuthOutput struct {
	Result string
	Token  Tokens
}

type Tokens struct {
	Session string
	Refresh string
}
