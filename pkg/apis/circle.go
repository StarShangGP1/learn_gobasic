package apis

type TopPost struct {
	ID            uint32
	ReleaseTime   int64
	PId           uint32
	PName         string
	Content       string
	ByTimeTall    float32
	ByTimeWeight  float32
	ByTimeFatRate float32
}

func (*Circle) TableName() string {
	return "learn_go.circle"
}
