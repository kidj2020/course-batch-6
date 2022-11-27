package domain

import (
	"errors"
	"time"
)

type QuestionNew struct {
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
	CreatorID     int    `json:"creator_id"`
	ExerciseID    int    `json:"exercise_id"`
}

type Question struct {
	ID            int       `json:"id"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"correct_answer"`
	Score         int       `json:"score"`
	CreatorID     int       `json:"creator_id"`
	ExerciseID    int       `json:"exercise_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func CreateNewquestion(exerciseId, creatorId int,
	body, option_a, option_b, option_c, option_d, correct_answer string) (*Question, error) {
	if body == "" {
		return nil, errors.New("body is required")
	}

	if option_a == "" {
		return nil, errors.New("option_a is required")
	}

	if option_b == "" {
		return nil, errors.New("option_b is required")
	}

	if option_c == "" {
		return nil, errors.New("option_c is required")
	}

	if option_d == "" {
		return nil, errors.New("option_d is required")
	}

	if correct_answer == "" {
		return nil, errors.New("correct_answer is required")
	}

	return &Question{
		ExerciseID:    exerciseId,
		Body:          body,
		OptionA:       option_a,
		OptionB:       option_b,
		OptionC:       option_c,
		OptionD:       option_d,
		CorrectAnswer: correct_answer,
		CreatorID:     creatorId,
	}, nil
}
