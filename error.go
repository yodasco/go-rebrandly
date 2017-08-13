package rebrandly

func (e *BadRequestResponse) Error() {
	return e.Message
}

func (e *UnauthorizedResponse) Error() {
	return e.Message
}

func (e *InvalidFormatResponse) Error() {
	return e.Message
}

func (e *AlreadyExistsResponse) Error() {
	return e.Message
}

func (e *NotFoundResponse) Error() {
	return e.Message
}

func (e *ServerErrorResponse) Error() {
	return e.Message
}
