package models

import "ms/sun/models/x"

type ReqParams struct {
	UserId      int
	SessionUuid string
	Session     *x.Session
	Page        int
	Last        int
	Limit       int
}

func (r *ReqParams) GetOffset() int {
	if r.Page < 1 {
		r.Page = 1
	}
	return (r.Page - 1) * r.Limit
}
