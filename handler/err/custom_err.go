package err

import "net/http"

var (

	// == user == //
	ERR_CREATE_USER          = CustomErr{StatusCode: http.StatusInternalServerError, Code: 402, Msg: "create user query err"}
	ERR_GET_USER             = CustomErr{StatusCode: http.StatusInternalServerError, Code: 402, Msg: "get user query err"}
	ERR_ALREADY_EMAIL        = CustomErr{StatusCode: http.StatusBadRequest, Code: 403, Msg: "this email already exists"}
	ERR_INCORRECT_PASSWORD   = CustomErr{StatusCode: http.StatusBadRequest, Code: 401, Msg: "incorrect password"}
	ERR_USER_NOT_MATCH       = CustomErr{StatusCode: http.StatusBadRequest, Code: 401, Msg: "user not match"}
	ERR_UPDATE_USER_PASSWORD = CustomErr{StatusCode: http.StatusInternalServerError, Code: 402, Msg: "update user password query err"}
	ERR_SAME_PASSWORD        = CustomErr{StatusCode: http.StatusBadRequest, Code: 401, Msg: "the current password and the new password are the same"}

	// == jwt == //
	ERR_JWT_CREATE_FAIL     = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "jwt create fail"}
	ERR_JWT_INVALID         = CustomErr{StatusCode: http.StatusBadRequest, Code: 401, Msg: "jwt invalid"}
	ERR_UNAUTHORIZED        = CustomErr{StatusCode: http.StatusUnauthorized, Code: 401, Msg: "unauthorized"}
	ERR_GET_FAIL            = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "jwt get fail"}
	ERR_JWT_ACCESS_EXPIRED  = CustomErr{StatusCode: http.StatusUnauthorized, Code: 401, Msg: "access token expired"}
	ERR_JWT_REFRESH_EXPIRED = CustomErr{StatusCode: http.StatusUnauthorized, Code: 401, Msg: "refresh token expired"}
	ERR_JWT_NOT_FOUND       = CustomErr{StatusCode: http.StatusUnauthorized, Code: 401, Msg: "jwt not found"}

	// == redis == //
	ERR_REDIS_SET_FAIL  = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "redis set fail"}
	ERR_REDIS_DEL_FAIL  = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "redis delete fail"}
	ERR_REDIS_NOT_DATA  = CustomErr{StatusCode: http.StatusBadRequest, Code: 401, Msg: "redis not data"}
	ERR_REDIS_NOT_FOUND = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "redis not found"}
	ERR_REDIS_DELETE    = CustomErr{StatusCode: http.StatusInternalServerError, Code: 002, Msg: "redis delete fail"}

	// == transaction == //
	ERR_TRANSACTION_SET      = CustomErr{StatusCode: http.StatusInternalServerError, Code: 500, Msg: "transaction set err"}
	ERR_TRANSACTION_COMMIT   = CustomErr{StatusCode: http.StatusInternalServerError, Code: 500, Msg: "transaction commit err"}
	ERR_TRANSACTION_ROLLBACK = CustomErr{StatusCode: http.StatusInternalServerError, Code: 500, Msg: "transaction rollback err"}

	// == bcrypt == //
	ERR_BCRYPT_GENERATE_FAIL = CustomErr{StatusCode: http.StatusInternalServerError, Code: 003, Msg: "bcrypt hash password generate fail"}
	ERR_BCRYPT_PARSING_FAIL  = CustomErr{StatusCode: http.StatusInternalServerError, Code: 003, Msg: "bcrypt hash password parsing fail"}
	ERR_BCRYPT_MISMATCH      = CustomErr{StatusCode: http.StatusInternalServerError, Code: 003, Msg: "mismatch hash and password"}

	// == binding and validator == //
	ERR_BINDING   = CustomErr{StatusCode: http.StatusBadRequest, Code: 000, Msg: "request binding err"}
	ERR_VALIDATOR = CustomErr{StatusCode: http.StatusBadRequest, Code: 001, Msg: "request validator err"}

	// == default == //
	CREATE_SUCCESS     = CustomErr{StatusCode: http.StatusOK, Code: 200, Msg: "request success"}
	ERR_USER_NOT_FOUND = CustomErr{StatusCode: http.StatusBadRequest, Code: 400, Msg: "user not found"}
)

type CustomErr struct {
	StatusCode int
	Code       int
	Msg        string
}

type ErrRes struct {
	StatusCode int    `json:"statusCode" example:"500"`
	Code       int    `json:"code" example:"500"`
	Msg        string `json:"msg" example:"internal server err"`
}

func NewErrRes(err CustomErr) *ErrRes {
	return &ErrRes{
		StatusCode: err.StatusCode,
		Code:       err.Code,
		Msg:        err.Msg,
	}
}

type BadRequestRes struct {
	StatusCode int    `json:"statusCode" example:"400"`
	Code       int    `json:"code" example:"400"`
	Msg        string `json:"msg" example:"bad request"`
}

func NewBadRequestRes(err CustomErr) *BadRequestRes {
	return &BadRequestRes{
		Code: err.Code,
		Msg:  err.Msg,
	}
}

type InternalServerErrRes struct {
	StatusCode int    `json:"statusCode" example:"500"`
	Code       int    `json:"code" example:"500"`
	Msg        string `json:"msg" example:"internal server err"`
}

type OkRes struct {
	StatusCode int    `json:"statusCode" example:"200"`
	Code       int    `json:"code" example:"200"`
	Msg        string `json:"msg" example:"ok"`
}

func NewOkRes() *OkRes {
	return &OkRes{
		StatusCode: CREATE_SUCCESS.StatusCode,
		Code:       CREATE_SUCCESS.Code,
		Msg:        CREATE_SUCCESS.Msg,
	}
}
