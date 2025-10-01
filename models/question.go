package models

import (
	"goLandCRUD/config"
	"goLandCRUD/utils"
)

type Question struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UpVotes   int64  `json:"up_votes"`
	DownVotes int64  `json:"down_votes"`
	Status    string `json:"status"`
}

type QuestionWithUser struct {
	Question      Question `json:"question"`
	User          User     `json:"user"`
	UpVotesList   []Vote   `json:"up_votes_list"`
	DownVotesList []Vote   `json:"down_votes_list"`
}

type Closeable interface {
	Close() error
}

func closeQuery(query Closeable) {
	err := query.Close()
	if err != nil {
		panic(err)
	}
}
func GetUserQuestionsList(userId int64) ([]Question, error) {
	query := utils.GetQuestionsByUserQuery
	rows, err := config.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer closeQuery(rows)
	var questions []Question
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.ID, &question.Title, &question.Body, &question.UserID, &question.CreatedAt, &question.UpdatedAt, &question.UpVotes, &question.DownVotes, &question.Status)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
func (question *Question) GetQuestionById() error {
	query := utils.GetQuestionByIdQuery
	row := config.DB.QueryRow(query, question.ID)
	err := row.Scan(&question.ID, &question.Title, &question.Body, &question.UserID, &question.CreatedAt, &question.UpdatedAt, &question.UpVotes, &question.DownVotes, &question.Status)
	if err != nil {
		return err
	}
	return nil
}
func (question *Question) CreateQuestion() error {
	query := utils.InsertQuestionQuery
	result, err := config.DB.Exec(query, question.Title, question.Body, question.UserID)
	if err != nil {
		return err
	}
	questionId, err := result.LastInsertId()
	question.ID = questionId
	if err != nil {
		return err
	}
	return nil
}
