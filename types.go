package emailist

import "fmt"

type Message struct {
	Subject     string
	TextMessage string
	HtmlMessage string
}

type Identity struct {
	Name  string
	Email string
}

func (i *Identity) From() string {
	return fmt.Sprintf("%s <%s>", i.Name, i.Email)
}

type MailingList struct {
	Name string
	Code string
}
