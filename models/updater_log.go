package models

import (
	"ms/sun/models/x"
	"time"
)

type updater struct {
	HereDirect   chan x.DirectLog
	StoredDirect chan x.DirectLog

	HereGroup   chan x.DirectLog
	StoredGroup chan x.DirectLog
}

var LogUpdater = updater{
	HereDirect:   make(chan x.DirectLog, 10000),
	StoredDirect: make(chan x.DirectLog, 10000),
}
