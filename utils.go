package main

import (
	"fmt"
	"os"
	"strconv"
)

func CheckAnswers(form_input map[string]string, db_answers []QuestionDefination) []QuestionResult {
	var result []QuestionResult
	for _, answer := range db_answers {
		if form_input[fmt.Sprintf("%d", answer.Id)] != "" {
			res, err := strconv.Atoi(form_input[fmt.Sprintf("%d", answer.Id)])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			result = append(result, QuestionResult{Id: answer.Id, IsCorrect: answer.Options[res].IsCorrect})
		}
	}

	return result
}
