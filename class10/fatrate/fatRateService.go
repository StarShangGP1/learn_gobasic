package fatrate

type fatRateService struct {
	c       Calc
	s       Suggestion
	persons []Person
}
