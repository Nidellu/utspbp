package model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type Participant struct {
	ID         int    `json:"id"`
	ID_account int    `json:"id_account"`
	Username   string `json:"username"`
}

type Game struct {
	ID         int `json:"id"`
	ID_account int `json:"game_name"`
	Max_player int `json:"max_player"`
}

type AccountResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

type AccounstResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Account `json:"data"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type ListRoom struct {
	Room []Room `json:"rooms"`
}

type ListRoomResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    ListRoom `json:"data"`
}

type RoomDetail struct {
	ID          int           `json:"id"`
	Room_name   string        `json:"room_name"`
	Participant []Participant `json:"participants"`
}

type RoomDetailResponse struct {
	RoomDetail RoomDetail `json:"room"`
}

type RoomsDetailResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    RoomDetailResponse `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
