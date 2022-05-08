package api

type Chat struct {
	Sender          string `json:"sender"`
	SenderAccount   string `json:"sender_account"`
	Receiver        string `json:"receiver"`
	ReceiverAccount string `json:"receiver_account"`
	Message         string `json:"message"`
	Status          string `json:"status"`
}
