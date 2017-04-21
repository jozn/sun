package models

type ReqParams struct {
	UserId      int
	SessionUuid string
	Session     *Session
	Page        int
	Last        int
	Limit       int
}
