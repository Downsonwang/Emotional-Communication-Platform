package service

import "Gin/service/note"

type ServiceGroup struct {
    NoteServiceGroup note.ServiceGroup
}

var ServiceGroupInfo = new(ServiceGroup)