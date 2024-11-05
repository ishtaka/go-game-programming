package chapter05

import (
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var sizeOfUint32 = int(unsafe.Sizeof(uint32(0)))
var sizeOfFloat32 = int(unsafe.Sizeof(float32(0)))

type VertexArray struct {
	// How many vertices in the vertex buffer?
	numVerts int
	// How many indices in the index buffer
	numIndices int
	// OpenGL ID of the vertex buffer
	vertexBuffer uint32
	// OpenGL ID of the index buffer
	indexBuffer uint32
	// OpenGL ID of the vertex array object
	vertexArray uint32
}

// NewVertexArray create a new vertex array
func NewVertexArray(vertices []float32, numVerts int, indices []uint32, numIndices int) *VertexArray {
	v := &VertexArray{
		numVerts:   numVerts,
		numIndices: numIndices,
	}

	// Create vertex array
	gl.GenVertexArrays(1, &v.vertexArray)
	gl.BindVertexArray(v.vertexArray)

	// Create vertex buffer
	gl.GenBuffers(1, &v.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, v.vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, numVerts*5*sizeOfFloat32, gl.Ptr(vertices), gl.STATIC_DRAW)

	// Create index buffer
	gl.GenBuffers(1, &v.indexBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, numIndices*sizeOfUint32, gl.Ptr(indices), gl.STATIC_DRAW)

	// Specify the vertex attributes
	// (For now, assume one vertex format)
	// Position is 3 floats starting at offset 0
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*int32(sizeOfFloat32), gl.PtrOffset(0))
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5*int32(sizeOfFloat32), gl.PtrOffset(3*sizeOfFloat32))

	return v
}

// SetActive activate this vertex array (so we can draw it)
func (v *VertexArray) SetActive() {
	gl.BindVertexArray(v.vertexArray)
}

func (v *VertexArray) GetNumIndices() int {
	return v.numIndices
}

func (v *VertexArray) GetNumVerts() int {
	return v.numVerts
}

func (v *VertexArray) Destroy() {
	gl.DeleteBuffers(1, &v.vertexBuffer)
	gl.DeleteBuffers(1, &v.indexBuffer)
	gl.DeleteVertexArrays(1, &v.vertexArray)
}
