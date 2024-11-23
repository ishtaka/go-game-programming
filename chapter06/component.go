package chapter06

const DefaultUpdateOrder = 100

type Component interface {
	// Update this component by delta time
	Update(deltaTime float32)
	// ProcessInput input for this component
	ProcessInput(keyState []uint8)
	// OnUpdateWorldTransform called when world transform changes
	OnUpdateWorldTransform()
	GetOwner() Actor
	GetUpdateOrder() int
	Destroy() // must override if embedded in a struct and call owner RemoveComponent
}

type component struct {
	// Owning actor
	owner Actor
	// Update order of component
	updateOrder int
}

func NewComponent(owner Actor, updateOrder int) Component {
	c := &component{
		owner:       owner,
		updateOrder: updateOrder,
	}

	return c
}

func (c *component) Update(deltaTime float32) {}

func (c *component) ProcessInput(keyState []uint8) {}

func (c *component) OnUpdateWorldTransform() {}

func (c *component) GetOwner() Actor {
	return c.owner
}

func (c *component) GetUpdateOrder() int {
	return c.updateOrder
}

func (c *component) Destroy() {
	c.owner.RemoveComponent(c)
}
