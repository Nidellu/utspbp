package controller

import (
	"database/sql"
	"log"
	"net/http"

	m "uts/model"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, 404, "Error: Parsing data")
		return
	}

	gameID := r.URL.Query().Get("id")

	if gameID == "" {
		sendErrorResponse(w, 404, "Bad Request: Incomplete input data")
		return
	}

	data, err := db.Begin()
	if err != nil {
		sendErrorResponse(w, 404, "Error: Database not found")
		return
	}
	defer data.Rollback()

	query := ("SELECT r.id, r.room_name from rooms r	WHERE r.id_game = ?")

	rows, err := db.Query(query, gameID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 404, "Can't connect to database")
		return
	}

	var listsRoom m.ListRoom
	for rows.Next() {
		var room m.Room
		if err := rows.Scan(&room.ID, &room.Room_name); err != nil {
			log.Println(err)
			sendErrorResponse(w, 400, "Can't get the data")
			return
		}
		listsRoom.Room = append(listsRoom.Room, room)
	}

	sendListRoomSuccessResponse(w, 200, "Success Print Data", listsRoom)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, 404, "Error: Parsing data")
		return
	}

	roomID := r.URL.Query().Get("id")

	if roomID == "" {
		sendErrorResponse(w, 404, "Bad Request: Incomplete input data")
		return
	}

	data, err := db.Begin()
	if err != nil {
		sendErrorResponse(w, 404, "Error: Database not found")
		return
	}
	defer data.Rollback()

	query := "SELECT r.id, r.room_name, p.id, p.id_account, c.username FROM rooms r JOIN participants p ON p.id_room = r.id JOIN accounts c ON c.id = p.id_account WHERE r.id = ?"

	rows, err := db.Query(query, roomID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 404, "Can't connect to database")
		return
	}

	var detailRoom m.RoomDetailResponse
	for rows.Next() {
		var participants m.Participant
		if err := rows.Scan(&detailRoom.RoomDetail.ID, &detailRoom.RoomDetail.Room_name, &participants.ID, &participants.ID_account, &participants.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w, 400, "Can't get the data")
			return
		}
		detailRoom.RoomDetail.Participant = append(detailRoom.RoomDetail.Participant, participants)
	}

	sendDetailRoomSuccessResponse(w, 200, "Success Print Data", detailRoom)
}

func InsertParticipants(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, 400, "Error: error parsing data")
		return
	}

	roomID := r.Form.Get("idRoom")
	accountID := r.Form.Get("idAccount")

	if accountID == "" || roomID == "" {
		sendErrorResponse(w, 400, "Bad request: Incomplete Data!")
		return
	}

	data, err := db.Begin()
	if err != nil {
		sendErrorResponse(w, 400, "Error: Database not Found!")
		return
	}
	defer data.Rollback()

	var game m.Game
	err = db.QueryRow("SELECT g.max_player FROM rooms r JOIN games g ON r.id_game = g.id WHERE r.id = ?", roomID).Scan(&game.Max_player)
	if err == sql.ErrNoRows {
		sendErrorResponse(w, 400, "Error room not found")
		return
	} else if err != nil {
		sendErrorResponse(w, 500, "Can't connect to database")
		return
	}

	var totalParticipants int
	err = db.QueryRow("SELECT COUNT(id) FROM participants WHERE id_room = ?", roomID).Scan(&totalParticipants)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "Can't connect to database")
		return
	}

	if totalParticipants >= game.Max_player {
		sendErrorResponse(w, 400, "Room is full")
		return
	}

	_, err = db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?, ?)", roomID, accountID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	sendParticipantsSuccessResponse(w, 200, "Entered the room successfully")
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	if err := r.ParseForm(); err != nil {
		sendErrorResponse(w, 400, "Error parsing form data")
		return
	}

	roomID := r.Form.Get("idRoom")
	accountID := r.Form.Get("idAccount")

	data, err := db.Begin()
	if err != nil {
		sendErrorResponse(w, 400, "Error: Database not Found!")
		return
	}
	defer data.Rollback()

	var participantID int
	err = db.QueryRow("SELECT id FROM participants WHERE id_room = ? AND id_account = ?", roomID, accountID).Scan(&participantID)
	if err == sql.ErrNoRows {
		sendErrorResponse(w, 400, "Can't find participant")
		return
	} else if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "Can't connect to database")
		return
	}

	_, err = db.Exec("DELETE FROM participants WHERE id = ?", participantID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "Internal Server Error")
		return
	}

	sendParticipantsSuccessResponse(w, 200, "Success left the room")
}
