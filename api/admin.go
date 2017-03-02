package api

import (
	"net/http"
	"github.com/herlegs/Undercover/logic"
	"github.com/herlegs/Undercover/api/dto"
	dao "github.com/herlegs/Undercover/storage"
)

const(
	ServerError = http.StatusInternalServerError
	OK = http.StatusOK
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request,_ := reqDto.(*dto.CreateRoomRequest)
	resp,err := logic.CreateNewRoom(request)
	if err != nil {
		WriteResponse(w, ServerError, err.Error())
	}
	WriteResponse(w, OK, resp)
}

func StartGameHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	//TODO
	request := reqDto.(*dto.StartGameRequest)
	request.AdminID = GetUserIDFromRequest(r)
	WriteResponse(w, OK, request)
}

func EndGameHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	//TODO
	request := reqDto.(*dto.EndGameRequest)
	request.AdminID = GetUserIDFromRequest(r)
	WriteResponse(w, OK, request)
}

func CloseRoomHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.CloseRoomRequest)
	request.AdminID = GetUserIDFromRequest(r)
	logic.CloseRoom(request)
}

func ValidateAdminHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	isAdmin := logic.IsRoomAdmin(request)
	resp := &dto.ValidateUserResponse{
		RoomID: request.RoomID,
		RoomStatus: dao.GetRoomStatus(request.RoomID),
	}
	if isAdmin {
		WriteResponse(w, OK, resp)
	}else{
		WriteResponse(w, http.StatusForbidden, "")
	}
}

func ValidatePlayerHandler(w http.ResponseWriter, r *http.Request, reqDto dto.Request){
	request := reqDto.(*dto.UserIdentityRequest)
	isAdmin := logic.IsRoomAdmin(request)
	resp := &dto.ValidateUserResponse{
		RoomID: request.RoomID,
		RoomStatus: dao.GetRoomStatus(request.RoomID),
	}
	if isAdmin {
		WriteResponse(w, http.StatusForbidden, "")
	}else{
		WriteResponse(w, OK, resp)
	}
}