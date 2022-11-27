package domain

import (
	"errors"
)

type CreateAnswersRequest struct {
	Answer string `json:"answer"`
}

func CreateNewAnswer(exerciseId, questionID, userID int,
	answer string) (*Answer, error) {
	if answer == "" {
		return nil, errors.New("answer is required")
	}

	return &Answer{
		ExerciseID: exerciseId,
		QuestionID: questionID,
		UserID:     userID,
		Answer:     answer,
	}, nil
}
