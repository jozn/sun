package galaxy

import (
    "net/http"
    "errors"
)

func Run() {
    http.Handle("/getFile/", &httpHandler{})
    http.Handle("/", &httpHandler{})
}

type fileFormat int

const (
    IMAGE = fileFormat(1)
    VIDEO = fileFormat(2)
    AUDIO = fileFormat(3)
    GIF   = fileFormat(4)
)

var errBadReq = errors.New("bad request")
var errWrongFileName = errors.New("wrong file name")
var errFileNameTooShort = errors.New("file name too short")


