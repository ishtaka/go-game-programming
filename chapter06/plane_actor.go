package chapter06

type PlaneActor struct {
	Actor
	mc MeshComponent
}

func NewPlaneActor(game *Game) *PlaneActor {
	a := NewActor(game)
	a.SetScale(10)

	mc := NewMeshComponent(a, DefaultUpdateOrder)
	m := game.GetRenderer().GetMesh("Assets/Plane.gpmesh")
	mc.SetMesh(m)
	a.AddComponent(mc)

	return &PlaneActor{
		Actor: a,
		mc:    mc,
	}
}

func (p *PlaneActor) Destroy() {
	p.GetGame().RemoveActor(p)
	p.DestroyComponents()
}
