package balance


type Balance interface {
	DoBalance([]*Instance) (*Instance, error)
}