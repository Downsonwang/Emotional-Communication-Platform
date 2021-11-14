package v1

import "Gin/api/v1/note"

type ApiV1Group struct {
  NoteApiGroup note.ApiGroup
}

var ApiGroupInfo = new(ApiV1Group)