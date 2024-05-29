package emailist

type Identity struct {
	Name  string
	Email string
}

func NewIdentiy(name, email string) *Identity {
	return &Identity{Name: name, Email: email}
}

func (i *Identity) SendMail(to []string, title, body, htmlBody string) error {
	return nil
}

func (i *Identity) SendTextMail(to []string, title, body string) error {
	return i.SendMail(to, title, body, "")
}

func (i *Identity) SendHTMLMail(to []string, title, body string) error {
	return i.SendMail(to, title, "", body)
}
