package main

func (q *QuestionDefination) Init(question string, id int) {
	q.Question = question
	q.Id = id
}

func (q *QuestionDefination) AddOption(option string, is_correct bool) {
	option_to_add := new(QuestionOption)
	option_to_add.IsCorrect = is_correct
	option_to_add.Option = option
	q.Options = append(q.Options, *option_to_add)
}

func GetQuestionsFromDatabase() []QuestionDefination {
	var all_questions []QuestionDefination
	question := new(QuestionDefination)
	question.Init("What comes after 2?", 1)
	question.AddOption("3", true)
	question.AddOption("4", false)
	question.AddOption("5", false)
	question.AddOption("6", false)
	all_questions = append(all_questions, *question)

	question = new(QuestionDefination)
	question.Init("What comes after 8?", 2)
	question.AddOption("3", false)
	question.AddOption("4", false)
	question.AddOption("10", false)
	question.AddOption("9", true)
	all_questions = append(all_questions, *question)

	question = new(QuestionDefination)
	question.Init("Capital of a?", 3)
	question.AddOption("A", true)
	question.AddOption("a", false)
	question.AddOption("b", false)
	question.AddOption("E", false)
	all_questions = append(all_questions, *question)

	question = new(QuestionDefination)
	question.Init("India starts with which word?", 4)
	question.AddOption("I", true)
	question.AddOption("a", false)
	question.AddOption("d", false)
	question.AddOption("n", false)
	all_questions = append(all_questions, *question)

	return all_questions
}

func SaveQuestionToDatabase(question FormQuestion) {

}

func DeleteQuestionFromDatabase(id int) {

}
