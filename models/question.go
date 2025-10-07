package models

import (
	"goLandCRUD/config"
	"goLandCRUD/utils"
)

type Question struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	UserID      int64  `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UpVotes     int64  `json:"up_votes"`
	DownVotes   int64  `json:"down_votes"`
	Status      string `json:"status"`
	AnswerCount int64  `json:"answer_count"`
}

type QuestionWithUser struct {
	Question       Question         `json:"question"`
	User           User             `json:"user"`
	AnswerWithUser []AnswerWithUser `json:"answers"`
	UpVotesList    []Vote           `json:"up_votes_list"`
	DownVotesList  []Vote           `json:"down_votes_list"`
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
func GetUserQuestionsList(userId int64) ([]QuestionWithUser, error) {
	query := utils.GetQuestionsByUserQuery
	rows, err := config.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer closeQuery(rows)
	var questions []QuestionWithUser
	var answersChan = make(chan *QuestionWithUser)
	index := 0
	for rows.Next() {
		var question QuestionWithUser
		err := rows.Scan(&question.Question.ID,
			&question.Question.Title, &question.Question.Body, &question.Question.UserID, &question.Question.CreatedAt, &question.Question.UpdatedAt, &question.Question.UpVotes, &question.Question.DownVotes, &question.Question.Status, &question.Question.AnswerCount)
		if err != nil {
			return nil, err
		}
		index++
		go func(question *QuestionWithUser) {
			answers, err := GetAllAnswersByQuestionId(question.Question.ID)
			if err != nil {
				answers = nil
			}
			question.AnswerWithUser = answers
			answersChan <- question
		}(&question)
	}
	for answers := range answersChan {
		questions = append(questions, *answers)
		if index == len(questions) {
			close(answersChan)
		}
	}
	return questions, nil
}
func (question *Question) GetQuestionById() error {
	query := utils.GetQuestionByIdQuery
	row := config.DB.QueryRow(query, question.ID)
	err := row.Scan(&question.ID, &question.Title, &question.Body, &question.UserID, &question.CreatedAt, &question.UpdatedAt, &question.UpVotes, &question.DownVotes, &question.Status, &question.AnswerCount)
	if err != nil {
		return err
	}
	return nil
}

func (question *QuestionWithUser) GetQuestionDetails() error {
	err := question.Question.GetQuestionById()
	if err != nil {
		return err
	}
	question.AnswerWithUser, err = GetAllAnswersByQuestionId(question.Question.ID)
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

func UpdateQuestionAnswerCount(questionId int64) error {
	query := utils.UpdateQuestionIncrementAnswerCountQuery
	_, err := config.DB.Exec(query, questionId)
	return err
}

func (question *Question) DeleteQuestion() error {
	question.Status = "deleted"
	query := utils.DeleteQuestionQuery
	_, err := config.DB.Exec(query, question.Status, question.ID)
	return err
}

func (question *Question) UpdateQuestion() error {
	query := utils.UpdateQuestionQuery
	result, err := config.DB.Exec(query, question.Title, question.Body)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
