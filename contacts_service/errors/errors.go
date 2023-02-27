package errors

var (
	EmptyFieldError        = "object cannot have any empty field"
	ContactAlreadyExist    = "contact with this informations already exist"
	ServerInternalErrorMsg = "server internal error, please try again"
	ContactNotFoundError   = "contact with this id not found"
	WrongIdFormatError     = "id format is wrong, check id"
	BadRequestMsg          = "bad request format"
)
