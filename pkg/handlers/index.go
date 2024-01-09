package handlers

type Handler struct {
	Home  *HomeHandler
	About *AboutHandler
}

func NewHandler(home *HomeHandler, about *AboutHandler) *Handler {
	return &Handler{
		Home:  home,
		About: about,
	}
}
