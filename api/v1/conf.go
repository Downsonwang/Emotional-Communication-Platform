package v1

import (
	"Gin/api/v1/note"
	"Gin/api/v1/pay"
	"Gin/api/v1/tag"
)

type ApiV1Group struct {
	NoteApiGroup note.ApiGroup
	TagApiGroup  tag.ApiGroup
	Pay          pay.ApiGroup
}

var ApiGroupInfo = new(ApiV1Group)
