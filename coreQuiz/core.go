package coreQuiz

import (
	"bufio"
	"encoding/json"
	"os"
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

type bytes [][]byte

func Check(e error) {
	if e != nil {
		panic(e)
	}
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

func (data bytes) SaveFile(Addr string) {
	f, err := os.Create(Addr)
	w := bufio.NewWriter(f)
	Check(err)
	defer f.Close()
	for _, i := range data {
		_, err = w.Write(i)
		w.Flush()
		Check(err)
	}
}

func (quizier Quiz) SaveFile(Addr string) {
	var data bytes
	for _, i := range quizier.SelectQuiz {
		temp, _ := json.Marshal(i)
		data = append(data, temp)
	}
	for _, i := range quizier.TypeQuiz {
		temp, err := json.Marshal(i)
		Check(err)
		data = append(data, temp)
	}
	data.SaveFile(Addr)
}

func (ans Answer) SaveFile(Addr string) {
	var data bytes
	for _, i := range ans.IntAnswers {
		temp, _ := json.Marshal(i)
		data = append(data, temp)
	}
	for _, i := range ans.StrAnswers {
		temp, err := json.Marshal(i)
		Check(err)
		data = append(data, temp)
	}
	data.SaveFile(Addr)
}
