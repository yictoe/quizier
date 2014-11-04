package coreQuiz

import (
	"bufio"
	"encoding/json"
	//"io"
	//"fmt"
	"io/ioutil"
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

type bytes []byte

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
	w.Write(data)
	w.Flush()
	defer f.Close()
}

func (quizier Quiz) SaveFile(Addr string) {
	var data bytes
	data, _ = json.Marshal(quizier)
	data.SaveFile(Addr)
	return
}

func (ans Answer) SaveFile(Addr string) {
	var data bytes
	data, _ = json.Marshal(ans)
	data.SaveFile(Addr)
	return
}

func ReadQuiz(Addr string) *Quiz {
	reading, err := ioutil.ReadFile(Addr)
	Check(err)
	var data map[string][]json.RawMessage
	var quizier Quiz
	err = json.Unmarshal(reading, &data)
	Check(err)
	for _, i := range data["SelectQuiz"] {
		var temp SelectQuizUnit
		err = json.Unmarshal(i, &temp)
		Check(err)
		quizier.SelectQuiz = append(quizier.SelectQuiz, &temp)
	}
	for _, i := range data["TypeQuiz"] {
		var temp TypeQuizUnit
		err = json.Unmarshal(i, &temp)
		quizier.TypeQuiz = append(quizier.TypeQuiz, &temp)
	}
	return &quizier
}

func ReadAns(Addr string) *Answer {
	reading, err := ioutil.ReadFile(Addr)
	Check(err)
	var data map[string][]json.RawMessage
	var ans Answer
	err = json.Unmarshal(reading, &data)
	Check(err)
	for _, i := range data["StrAnswers"] {
		var temp StrAnswerUnit
		err = json.Unmarshal(i, &temp)
		Check(err)
		ans.StrAnswers = append(ans.StrAnswers, &temp)
	}
	for _, i := range data["IntAnswers"] {
		var temp IntAnswerUnit
		err = json.Unmarshal(i, &temp)
		ans.IntAnswers = append(ans.IntAnswers, &temp)
	}
	return &ans
}
