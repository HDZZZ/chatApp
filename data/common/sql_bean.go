package common

type User struct {
	Id       int    `mapstructure:"uid"`
	UserName string `mapstructure:"user_name"`
	Password string `json:"-" mapstructure:"pass_word"`
	Token    string `json:"Token,omitempty"`
}

type DBMessage struct {
	Msg_id            int    `mapstructure:"msg_id"`
	Sender_id         int    `mapstructure:"sender_id"`
	Receiver_id       int    `mapstructure:"receiver_id"`
	Conversatio_type  int    `mapstructure:"conversatio_type"`
	Message_body_type int    `mapstructure:"message_body_type"`
	Content           string `mapstructure:"content"`
}

type ReuqestOfAddingFriend struct {
	Sender_id    int          `mapstructure:"sender_uid"`
	Receiver_id  int          `mapstructure:"receiver_uid"`
	Id           int          `mapstructure:"id"`
	Msg          string       `mapstructure:"msg"`
	Requst_state RequestState `mapstructure:"request_state"`
}

type RequestState int8

const (
	Defualt       RequestState = 0
	AlreadyAgree  RequestState = 1
	AlreadyRefuse RequestState = 2
	Timeout       RequestState = 3
	NotWork       RequestState = 4
)

type Group struct {
	Gid         int    `mapstructure:"gid"`
	GroupName   string `mapstructure:"group_name"`
	OwnerId     int    `mapstructure:"ownerId"`
	Description string `mapstructure:"description"`
	MemberCount int    `mapstructure:"member_count"`
}

type GroupMember struct {
	Gid      int            `mapstructure:"gid"`
	Alias    string         `mapstructure:"alias"`
	Uid      int            `mapstructure:"uid"`
	Identity MemberIdentity `mapstructure:"identity"`
}
type MemberIdentity int8

const (
	Member  MemberIdentity = 0
	Owner   MemberIdentity = 1
	Manager MemberIdentity = 2
	NONE    MemberIdentity = 3
)
