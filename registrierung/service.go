package registrierung

type RegistrierungsService struct {
	Repository RegistrierungsRepository
	Notifier   RegistrierungsNotifier
}

func (rs *RegistrierungsService) HandleNewRegistrierung(reg *Registrierung) error {
	return rs.Repository.SaveRegistrierung(reg)
}

func (rs *RegistrierungsService) ConfirmRegistration(regId string) error {
	registrierung, err := rs.Repository.ConfirmedRegistrierung(regId)
	if err != nil {
		return err
	}

	return rs.Notifier.InformAboutNewRegistrierung(registrierung)
}

func (rs *RegistrierungsService) GetUnconfirmedRegistrierung() ([]*Registrierung, error) {
	return rs.Repository.GetUnconfirmedRegistrierung()
}
