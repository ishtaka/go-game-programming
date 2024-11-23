package chapter06

import "github.com/go-gl/gl/v3.3-core/gl"

type MeshComponent interface {
	Component
	Draw(shader *Shader)
	SetMesh(mesh *Mesh)
	SetTextureIndex(index int)
}

func NewMeshComponent(owner Actor, updateOrder int) MeshComponent {
	c := NewComponent(owner, updateOrder)
	mc := &meshComponent{
		Component: c,
	}
	owner.GetGame().GetRenderer().AddMeshComp(mc)

	return mc
}

type meshComponent struct {
	Component
	mesh         *Mesh
	textureIndex int
}

func (m *meshComponent) Draw(shader *Shader) {
	if m.mesh != nil {
		// Set the world transform
		worldTrans := m.GetOwner().GetWorldTransform()
		shader.SetMatrixUniform("uWorldTransform", &worldTrans)
		// Set specular power
		shader.SetFloatUniform("uSpecPower", m.mesh.SpecPower())
		// Set the active texture
		tex := m.mesh.GetTexture(m.textureIndex)
		if tex != nil {
			tex.SetActive()
		}
		// Set the mesh's vertex array as active
		va := m.mesh.GetVertexArray()
		va.SetActive()
		// Draw
		gl.DrawElements(gl.TRIANGLES, int32(va.GetNumIndices()), gl.UNSIGNED_INT, nil)
	}
}

func (m *meshComponent) SetMesh(mesh *Mesh) {
	m.mesh = mesh
}

func (m *meshComponent) SetTextureIndex(index int) {
	m.textureIndex = index
}

func (m *meshComponent) Destroy() {
	owner := m.GetOwner()
	owner.GetGame().GetRenderer().RemoveMeshComp(m)
	owner.RemoveComponent(m)
}
