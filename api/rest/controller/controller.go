package controller

func MapControllerHandler(_type string) Con {
	switch _type {
	case "home":
		return NewAppAPI()
	case "about":
		return NewAboutAPI()
	}
	return nil
}