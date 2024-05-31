package emailist

type Message struct {
	TextMessage string
	HtmlMessage string
}

type Identity struct {
	Name  string
	Email string
}

type MailingList struct {
	Name string
	Code string
}
