package service

import (
	"Gin/service/note"
	"Gin/service/register"
)

type ServiceGroup struct {
	NoteServiceGroup     note.ServiceGroup
	ResisterServiceGroup register.ServiceGroup
}

var ServiceGroupInfo = new(ServiceGroup)
