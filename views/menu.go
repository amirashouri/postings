package views

const (
	HOME_TAB     = "home"
	ABOUT_TAB    = "about"
	CONTACT_TAB  = "contact"
	PROJECTS_TAB = "projects"
	LOGIN_TAB    = "login"
	SIGNUP_TAB   = "signup"
)

type MenuState struct {
	IsAuthenticated bool
	IsOpen          bool
	SelectedPage    string
}

func NewMenuState(isOpen bool, selectedState string) *MenuState {
	return &MenuState{IsOpen: isOpen, SelectedPage: selectedState}
}
