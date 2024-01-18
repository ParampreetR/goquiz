package main

type FormQuestion struct {
	Question string `json:"question" xml:"question" form:"question"`
	Option1  string `json:"1" xml:"1" form:"1"`
	Option2  string `json:"2" xml:"2" form:"2"`
	Option3  string `json:"3" xml:"3" form:"3"`
	Option4  string `json:"4" xml:"4" form:"4"`
	Answer   string `json:"correct" xml:"correct" form:"correct"`
}

type User struct {
	Name string `json:"user" xml:"user" form:"user"`
}

type QuestionOption struct {
	Option    string
	IsCorrect bool
}

type QuestionDefination struct {
	Id       int
	Question string
	Options  []QuestionOption
}

type QuestionResult struct {
	Id        int
	IsCorrect bool
}
