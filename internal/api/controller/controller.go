package controller

//go:generate mockery --name=Controller --case=snake
type Controller interface {
	UtilityController() UtilityController
}

func New(
	utilityController UtilityController,
) Controller {
	return &controllerImpl{
		utilityController,
	}
}

type controllerImpl struct {
	utilityController UtilityController
}

func (i *controllerImpl) UtilityController() UtilityController {
	return i.utilityController
}
