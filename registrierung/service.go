package registrierung

//RegistrierungsService stellt die Businessfunktionalität für Registrierungen bereit
type RegistrierungsService struct {
	Repository RegistrierungsRepository
	Notifier   RegistrierungsNotifier
}

//HandleNewRegistrierung behandelte neue Registrierungen
func (rs *RegistrierungsService) HandleNewRegistrierung(reg *Registrierung) error {
	return rs.Repository.SaveRegistrierung(reg)
}

//ConfirmRegistration bestätigt eine Registrierung
func (rs *RegistrierungsService) ConfirmRegistration(regID string) error {
	registrierung, err := rs.Repository.ConfirmedRegistrierung(regID)
	if err != nil {
		return err
	}

	return rs.Notifier.InformAboutNewRegistrierung(registrierung)
}

//GetUnconfirmedRegistrierungen liefert alle noch nicht bestätigten Registrierungen
func (rs *RegistrierungsService) GetUnconfirmedRegistrierungen() ([]*Registrierung, error) {
	return rs.Repository.GetUnconfirmedRegistrierungen()
}
