package coreQuiz

import (
	"strings"
)

type SelectQuizUnit struct {
	No       int
	Score    float32
	BodyText string
	Example  [10]string
	Hint     string
	Answer   int
	Email    string
}

type TypeQuizUnit struct {
	No       int
	Score    float32
	BodyText string
	Hint     string
	Answer   string
	Email    string
}

type WriteQuizUnit struct {
	No       int
	Score    float32
	BodyText string
	Hint     string
	Email    string
}

type Quiz struct {
	SelectQuiz []*SelectQuizUnit
	TypeQuiz   []*TypeQuizUnit
	WriteQuiz  []*WriteQuizUnit
}

type IntAnswerUnit struct {
	No     int
	Answer int
}

type StrAnswerUnit struct {
	No     int
	Answer string
}

type Answer struct {
	IntAnswers []*IntAnswerUnit
	StrAnswers []*StrAnswerUnit
}

func (ans Answer) Score(quizier Quiz) float32 {
	var sc float32 = 0
	for _, i := range ans.IntAnswers {
		for _, j := range quizier.SelectQuiz {
			if (*i).No == (*j).No {
				if (*i).Answer == (*j).Answer {
					sc = sc + (*j).Score
				}
			}
		}
	}
	for _, i := range ans.StrAnswers {
		for _, j := range quizier.TypeQuiz {
			if (*i).No == (*j).No {
				if strings.EqualFold((*i).Answer, (*j).Answer) {
					sc = sc + (*j).Score
				}
			}
		}
	}
	return sc
}

func (quizier Quiz) TotalScore() float32 {
	var sc float32 = 0
	for _, i := range quizier.SelectQuiz {
		sc = sc + (*i).Score
	}
	for _, i := range quizier.TypeQuiz {
		sc = sc + (*i).Score
	}
	return sc
}
