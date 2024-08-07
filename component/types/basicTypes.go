package types

func NewServer(listenAddr string)*Server{
	return &Server{
		listenAddr: listenAddr,
	}
}


type Server struct{
	listenAddr string
}

type serverError struct{
	Error string `json:"error"`
}



type NewLoginRequest struct {
	Email    string `json:"email"`
}