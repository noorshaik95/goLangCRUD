package models

import (
	"goLandCRUD/config"
	"goLandCRUD/utils"
)

type Answer struct {
	ID         int64  `json:"id"`
	Body       string `json:"body"`
	UserID     int64  `json:"user_id"`
	QuestionID int64  `json:"question_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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

func (answer *Answer) Save() error {
	query := utils.InsertAnswerQuery
	stmt, err := config.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(answer.Body, answer.UserID, answer.QuestionID)
	if err != nil {
		return err
	}
	answerId, err := result.LastInsertId()
	if err != nil {

	}
	answer.ID = answerId
	return nil
}

func GetAllAnswersByQuestionId(questionId int64) ([]AnswerWithUser, error) {
	query := utils.GetAnswersByQuestionIdQuery
	rows, err := config.DB.Query(query, questionId)
	if err != nil {
		return nil, err
	}
	defer closeQuery(rows)

	var answers []AnswerWithUser
	for rows.Next() {
		var answer AnswerWithUser
		err := rows.Scan(&answer.ID, &answer.Body, &answer.UserID, &answer.QuestionID, &answer.CreatedAt, &answer.UpdatedAt, &answer.UpVotes, &answer.DownVotes, &answer.Status,
			&answer.User.ID, &answer.User.Username, &answer.User.Email)
		if err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}
