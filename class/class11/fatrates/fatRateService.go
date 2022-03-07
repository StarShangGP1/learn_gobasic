package fatrates

import "log"

type FatRateService struct {
	S *Suggestion
	Calc
}

func (frsvc *FatRateService) GiveSuggestionToPerson(p *Person) string {
	if err := frsvc.calcBmi(p); err != nil {
		log.Println("err", err)
		return "error"
	}
	frsvc.calcFatRate(p)
	return frsvc.S.GetSuggestion(p)
}

func (frsvc *FatRateService) GiveSuggestionToPersons(ps ...*Person) map[*Person]string {
	out := map[*Person]string{}
	for _, item := range ps {
		out[item] = frsvc.S.GetSuggestion(item)
	}
	return out
}
