package emailist

func NewIdentity(name, email string) *Identity {
	return &Identity{Name: name, Email: email}
}
