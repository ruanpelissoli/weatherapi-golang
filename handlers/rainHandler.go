package handlers

type RainHandler struct {
}

func NewRainHandler() *RainHandler {
	return &RainHandler{}
}

func (w RainHandler) IsRaining(city string) (string, error) {
	return "It is raining dude", nil
}
