package db

type User struct {
	Id       int
	UserName string
	Password string `json:"-"`
	Token    string
}

type DBMessage struct {
	Msg_id            int
	Sender_id         int
	Receiver_id       int
	Conversatio_type  int
	Message_body_type int
	Content           string
}

type ReuqestOfAddingFriend struct {
	Sender_id    int
	Receiver_id  int
	Id           int
	Msg          string
	Requst_state RequestState
}

type RequestState int8

const (
	Defualt       RequestState = 0
	AlreadyAgree  RequestState = 1
	AlreadyRefuse RequestState = 2
	Timeout       RequestState = 3
	NotWork       RequestState = 4
)
