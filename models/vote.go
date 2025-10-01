package models

type Vote struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"user_id"`
	AnswerID   int64 `json:"answer_id"`
	QuestionID int64 `json:"question_id"`
	Type       int   `json:"type"`
}
