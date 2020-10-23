package registrierung

//RegistrierungsRepository verwaltet alle Registrierungen
type RegistrierungsRepository interface {
	//SaveRegistrierung speichert eine Registrierung
	SaveRegistrierung(*Registrierung) error
	//GetUnconfirmedRegistrierung liefert alle unbestätigten Registrierungen
	GetUnconfirmedRegistrierungen() ([]*Registrierung, error)
	//ConfirmedRegistrierung bestätigt eine Registrierung
	ConfirmedRegistrierung(registrierungID string) (*Registrierung, error)
}
