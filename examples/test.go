package main

import (
	"fmt"
	"github.com/yictoe/quizier/coreQuiz"
)

func main() {

	var a1, a2 coreQuiz.SelectQuizUnit
	a1.No = 1
	a1.Answer = 3
	a1.Score = 5
	a2.No = 3
	a2.Answer = 4
	a2.Score = 3.3
	var b1, b2 coreQuiz.TypeQuizUnit
	b1.No = 2
	b1.Answer = "홍길동"
	b1.Score = 5.2
	b2.No = 4
	b2.Answer = "Chair"
	b2.Score = 2.7
	var quiz coreQuiz.Quiz
	temp := []*(coreQuiz.SelectQuizUnit){&a1, &a2}
	quiz.SelectQuiz = temp
	temp2 := []*(coreQuiz.TypeQuizUnit){&b1, &b2}
	quiz.TypeQuiz = temp2
	fmt.Printf("%f\n", quiz.TotalScore())

	var ans coreQuiz.Answer
	var c1, c2 coreQuiz.IntAnswerUnit
	var d1, d2 coreQuiz.StrAnswerUnit
	c1.No = 3
	c1.Answer = 4
	c2.No = 1
	c2.Answer = 1
	d1.No = 2
	d1.Answer = "홍길동"
	d2.No = 4
	d2.Answer = "CcCcC"
	ans.IntAnswers = []*(coreQuiz.IntAnswerUnit){&c1, &c2}
	ans.StrAnswers = []*(coreQuiz.StrAnswerUnit){&d1, &d2}
	fmt.Printf("%f", ans.Score(quiz))
}