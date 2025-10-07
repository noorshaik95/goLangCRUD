package utils

var GetQuestionsByUserQuery = `SELECT
							id, title, body,
							user_id, created_at,
							updated_at, up_votes, down_votes, status, answer_count
							FROM questions WHERE user_id = ?`

var GetQuestionByIdQuery = `SELECT
							id, title, body,
							user_id, created_at,
							updated_at, up_votes, down_votes, status, answer_count
							FROM questions WHERE id = ?`

var GetAllQuestionsQuery = `SELECT
							id, title, body,
							user_id, created_at,
							updated_at, up_votes, down_votes
							FROM questions`

var GetAnswersByQuestionIdQuery = `SELECT
							answers.id as id, body,
							user_id, question_id,
							answers.created_at as created_at, answers.updated_at as updated_at,
							up_votes, down_votes,
							answers.status as status,
							users.id AS user_id,
							users.username as username,
							users.email AS email
							FROM answers INNER JOIN users ON users.id = answers.user_id WHERE question_id = ?`

var GetAnswerByIdQuery = `SELECT
							id, body,
							user_id, question_id,
							created_at, updated_at,
							up_votes, down_votes
							FROM answers WHERE id = ?`

var GetUserByIdQuery = `SELECT
						id, username,
						email, password,
						created_at, updated_at
						FROM users WHERE id = ?`

var GetUserByEmailQuery = `SELECT
						id, password
						FROM users WHERE email = ?`

var GetVotesByAnswerIdQuery = `SELECT
							id, user_id,
							answer_id, question_id,
							created_at, updated_at,
							type
							FROM votes WHERE answer_id = ?`

var GetVotesByQuestionIdQuery = `SELECT
							id, user_id,
							answer_id, question_id,
							created_at, updated_at,
							type
							FROM votes WHERE question_id = ?`

var InsertUserQuery = `INSERT INTO users
						(username, email, password)
						VALUES (?, ?, ?)`

var InsertQuestionQuery = `INSERT INTO questions
							(title, body, user_id)
							VALUES (?, ?, ?)`

var InsertAnswerQuery = `INSERT INTO answers
						(body, user_id, question_id)
						VALUES (?, ?, ?)`

var InsertVoteQuery = `INSERT INTO votes
						(user_id, answer_id, question_id, type)
						VALUES (?, ?, ?, ?)`

var UpdateQuestionVotesQuery = `UPDATE questions
								SET up_votes = up_votes + ?,
								down_votes = down_votes + ?,
								updated_at = CURRENT_TIMESTAMP
								WHERE id = ?`

var UpdateAnswerVotesQuery = `UPDATE answers
								SET up_votes = up_votes + ?,
								down_votes = down_votes + ?,
								updated_at = CURRENT_TIMESTAMP
								WHERE id = ?`

var DeleteVoteQuery = `DELETE FROM votes
						WHERE user_id = ? AND 
						((answer_id = ? AND question_id IS NULL) OR
						(question_id = ? AND answer_id IS NULL))`

var UpdateUserQuery = `UPDATE users
						SET username = ?, email = ?, password = ?, updated_at = CURRENT_TIMESTAMP
						WHERE id = ?`

var UpdateQuestionQuery = `UPDATE questions
							SET title = ?, body = ?, updated_at = CURRENT_TIMESTAMP
							WHERE id = ?`
var UpdateQuestionIncrementAnswerCountQuery = `UPDATE questions
							SET answer_count = answer_count + 1, updated_at = CURRENT_TIMESTAMP
							WHERE id = ?`

var UpdateAnswerQuery = `UPDATE answers
						SET body = ?, updated_at = CURRENT_TIMESTAMP
						WHERE id = ?`

var DeleteQuestionQuery = `UPDATE questions SET status = ? WHERE id = ?`

var DeleteAnswerQuery = `DELETE FROM answers WHERE id = ?`

var DeleteUserQuery = `DELETE FROM users WHERE id = ?`
