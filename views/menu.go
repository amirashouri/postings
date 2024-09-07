package views

const (
	INDEX_TAB    = "postings"
	HOME_TAB     = "home"
	MY_ACCOUNT   = "my account"
	ABOUT_TAB    = "about"
	CONTACT_TAB  = "contact"
	PROJECTS_TAB = "projects"
	LOGIN_TAB    = "login"
	SIGNUP_TAB   = "signup"
	POSTS_TAB    = "posts"
)

type MenuState struct {
	IsAuthenticated bool
	IsOpen          bool
	SelectedPage    string
}

func NewMenuState(isOpen bool, selectedState string) *MenuState {
	return &MenuState{IsOpen: isOpen, SelectedPage: selectedState}
}
