package controller

//go:generate mockery --name=Controller --case=snake
type Controller interface {
	UtilityController() UtilityController
	SessionController() SessionController
}

type controllerImpl struct {
	utilityController UtilityController
	sessionController SessionController
}

func New(
	utilityController UtilityController,
	sessionController SessionController,
) Controller {
	return &controllerImpl{
		utilityController,
		sessionController,
	}
}
func (i *controllerImpl) UtilityController() UtilityController {
	return i.utilityController
}
func (i *controllerImpl) SessionController() SessionController {
	return i.sessionController
}
