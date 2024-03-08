package controller

import (
	"encoding/json"
	"net/http"

	m "uts/model"
)

func sendListRoomSuccessResponse(w http.ResponseWriter, code int, message string, listRooms m.ListRoom) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ListRoomResponse
	response.Status = code
	response.Message = message
	response.Data = listRooms
	json.NewEncoder(w).Encode(response)
}

func sendDetailRoomSuccessResponse(w http.ResponseWriter, code int, message string, listRooms m.RoomDetailResponse) {
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsDetailResponse
	response.Status = code
	response.Message = message
	response.Data = listRooms
	json.NewEncoder(w).Encode(response)
}

func sendParticipantsSuccessResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ParticipantResponse
	response.Status = code
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
