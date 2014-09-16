package infrastructureFacade

func NewInfrastructureFacades() map[string]IInfrastructureFacade {
	return map[string]IInfrastructureFacade{
		"private_machine": PrivateMachineFacade{},
		"qsub":            QsubFacade{},
		"qcg":             QcgFacade{},
	}
}
