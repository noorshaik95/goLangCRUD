package models

type Answer struct {
	ID         int64  `json:"id"`
	Body       string `json:"body"`
	UserID     int64  `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	UpVotes    int64  `json:"up_votes"`
	DownVotes  int64  `json:"down_votes"`
	Status     string `json:"status"`
}

type AnswerWithUser struct {
	Answer
	User          User   `json:"user"`
	UpVotesList   []Vote `json:"up_votes_list"`
	DownVotesList []Vote `json:"down_votes_list"`
}
