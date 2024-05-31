package emailist

func NewMailingList(name string) *MailingList {
	return &MailingList{Name: name}
}

func (m *MailingList) Create(emails ...string) error {
	return nil
}

func (m *MailingList) Delete(emails ...string) error {
	return nil
}

func (m *MailingList) Rename() error {
	return nil
}

func (m *MailingList) Subscribe(emails ...string) error {
	return nil
}

func (m *MailingList) Unsubscribe(emails ...string) error {
	return nil
}
