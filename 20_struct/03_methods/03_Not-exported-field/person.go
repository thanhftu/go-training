package person

// Person has unexported firstName
type Person struct {
	firstName string
	lastName  string
}

// FirstNameCaller used to call firstName
func (p Person) FirstNameCaller() string {
	return p.firstName
}

// FirstNameSettler used to set firstName
func (p *Person) FirstNameSettler(newName string) {
	p.firstName = newName
}
