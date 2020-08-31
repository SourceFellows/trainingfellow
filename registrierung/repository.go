package registrierung

type RegistrierungsRepository interface {
	SaveRegistrierung(*Registrierung) error
	GetUnconfirmedRegistrierung() ([]*Registrierung, error)
	ConfirmedRegistrierung(registrierungId string) (*Registrierung, error)
}
