package models

type VoteDataArgs struct {
	NoteID    string  `json:"note_id" binding:"required" `
	Direction float64 `json:"direction" binding:"oneof=1 0 -1"`
}
