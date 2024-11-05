package chapter05

type InputComponent interface {
	Component
	ProcessInput(keyState []uint8)

	GetMaxForward() float32
	SetMaxForwardSpeed(speed float32)

	GetMaxAngular() float32
	SetMaxAngularSpeed(speed float32)

	GetForwardKey() uint8
	SetForwardKey(key uint8)
	GetBackKey() uint8
	SetBackKey(key uint8)

	GetClockwiseKey() uint8
	SetClockwiseKey(key uint8)

	GetCounterClockwiseKey() uint8
	SetCounterClockwiseKey(key uint8)
}

type inputComponent struct {
	MoveComponent
	// The maximum forward/angular speeds
	maxForwardSpeed float32
	maxAngularSpeed float32
	// Keys for forward/back movement
	forwardKey uint8
	backKey    uint8
	// Keys for angular movement
	clockwiseKey        uint8
	counterClockwiseKey uint8
}

func NewInputComponent(owner Actor, updateOrder int) InputComponent {
	c := NewMoveComponent(owner, updateOrder)
	i := &inputComponent{
		MoveComponent: c,
	}

	return i
}

func (i *inputComponent) ProcessInput(keyState []uint8) {
	// Calculate forward speed for MoveComponent
	forwardSpeed := float32(0.0)
	if keyState[i.forwardKey] != 0 {
		forwardSpeed += i.maxForwardSpeed
	}
	if keyState[i.backKey] != 0 {
		forwardSpeed -= i.maxForwardSpeed
	}
	i.SetForwardSpeed(forwardSpeed)

	// Calculate angular speed for MoveComponent
	angularSpeed := float32(0.0)
	if keyState[i.clockwiseKey] != 0 {
		angularSpeed += i.maxAngularSpeed
	}
	if keyState[i.counterClockwiseKey] != 0 {
		angularSpeed -= i.maxAngularSpeed
	}
	i.SetAngularSpeed(angularSpeed)
}

func (i *inputComponent) GetMaxForward() float32 {
	return i.maxForwardSpeed
}

func (i *inputComponent) SetMaxForwardSpeed(speed float32) {
	i.maxForwardSpeed = speed
}

func (i *inputComponent) GetMaxAngular() float32 {
	return i.maxAngularSpeed
}

func (i *inputComponent) SetMaxAngularSpeed(speed float32) {
	i.maxAngularSpeed = speed
}

func (i *inputComponent) GetForwardKey() uint8 {
	return i.forwardKey
}

func (i *inputComponent) SetForwardKey(key uint8) {
	i.forwardKey = key
}

func (i *inputComponent) GetBackKey() uint8 {
	return i.backKey
}

func (i *inputComponent) SetBackKey(key uint8) {
	i.backKey = key
}

func (i *inputComponent) GetClockwiseKey() uint8 {
	return i.clockwiseKey
}

func (i *inputComponent) SetClockwiseKey(key uint8) {
	i.clockwiseKey = key
}

func (i *inputComponent) GetCounterClockwiseKey() uint8 {
	return i.counterClockwiseKey
}

func (i *inputComponent) SetCounterClockwiseKey(key uint8) {
	i.counterClockwiseKey = key
}

func (i *inputComponent) Destroy() {
	owner := i.GetOwner()
	owner.RemoveComponent(i)
}
