package contact

type User struct {
	PhoneNumber string // +1 123 456 789 00
}

func (u User) GetPhoneArea() {

}

func (u User) GetPhoneCountry() {

}

type Suggestion struct {
	User  User
	Score int
}
type SuggestionList []Suggestion

func (p SuggestionList) Len() int           { return len(p) }
func (p SuggestionList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p SuggestionList) Less(i, j int) bool { return p[i].Score > p[j].Score }
