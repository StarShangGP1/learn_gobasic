package fatrates

type Suggestion struct {
	suggestArr [][][]int
}

func GetFatRateSuggestion() *Suggestion {
	return &Suggestion{
		suggestArr: [][][]int{
			{ // 男
				{ // 18-39
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ // 40-59

				},
				{ // 60-

				},
			},
			{ // 女
				{ // 18-39

				},
				{ // 40-59

				},
				{ // 60-

				},
			},
		},
	}
}

func (s *Suggestion) GetSuggestion(p *Person) string {
	sexIdx := s.getIndexOfSex(p.Sex)
	ageIdx := s.getIndexOfAge(p.Age)
	maxFRSupported := len(s.suggestArr[sexIdx][ageIdx]) - 1
	frIdx := int(p.FatRate * 100)
	if frIdx > maxFRSupported {
		frIdx = maxFRSupported
	}
	suggestIdx := s.suggestArr[sexIdx][ageIdx][frIdx]
	result := s.translateResult(suggestIdx)
	return result
}

func (s *Suggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

func (s *Suggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 1
	case age >= 60:
		return 2
	default:
		return -1
	}
}

func (s *Suggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "非常肥胖"
	default:
		return "未知"
	}
}
