package v1

import (
	"Gin/api/v1/note"
	"Gin/api/v1/tag"
)

type ApiV1Group struct {
	NoteApiGroup note.ApiGroup
	TagApiGroup  tag.ApiGroup
}

var ApiGroupInfo = new(ApiV1Group)
