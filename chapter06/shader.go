package chapter06

import (
	"embed"
	_ "embed"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/ishtaka/go-game-programming/chapter06/math"
)

//go:embed shaders/*
var shaders embed.FS

type Shader struct {
	vertexShader  uint32
	fragShader    uint32
	shaderProgram uint32
}

func NewShader() *Shader {
	return &Shader{}
}

func (s *Shader) Load(vertName, fragName string) bool {
	var ok bool
	// Compile the vertex shader
	s.vertexShader, ok = s.compileShader("shaders/"+vertName, gl.VERTEX_SHADER)
	if !ok {
		return false
	}

	// Compile the fragment shader
	s.fragShader, ok = s.compileShader("shaders/"+fragName, gl.FRAGMENT_SHADER)
	if !ok {
		return false
	}

	// Now create a shader program that
	// links together the vertex/frag shaders
	s.shaderProgram = gl.CreateProgram()
	gl.AttachShader(s.shaderProgram, s.vertexShader)
	gl.AttachShader(s.shaderProgram, s.fragShader)
	gl.LinkProgram(s.shaderProgram)

	// Verify that the program linked successfully
	if !s.isValidProgram() {
		return false
	}

	return true
}

func (s *Shader) Unload() {
	// Delete the program/shaders
	gl.DeleteProgram(s.shaderProgram)
	gl.DeleteShader(s.vertexShader)
	gl.DeleteShader(s.fragShader)
}

func (s *Shader) SetActive() {
	// Set this program as the active one
	gl.UseProgram(s.shaderProgram)
}

func (s *Shader) SetMatrixUniform(name string, matrix *math.Matrix4) {
	// Find the uniform by this name
	loc := gl.GetUniformLocation(s.shaderProgram, gl.Str(name+"\x00"))
	// Send the matrix data to the uniform
	gl.UniformMatrix4fv(loc, 1, true, matrix.GetAsFloatPtr())
}

func (s *Shader) SetVectorUniform(name string, vector math.Vector3) {
	loc := gl.GetUniformLocation(s.shaderProgram, gl.Str(name+"\x00"))
	// Send the vector data
	gl.Uniform3fv(loc, 1, vector.AsFloatPtr())
}

func (s *Shader) SetFloatUniform(name string, value float32) {
	loc := gl.GetUniformLocation(s.shaderProgram, gl.Str(name+"\x00"))
	// Send the float data
	gl.Uniform1f(loc, value)
}

func (s *Shader) compileShader(fileName string, shaderType uint32) (uint32, bool) {
	f, err := shaders.ReadFile(fileName)
	if err != nil {
		sdl.Log("failed to open file: %s\n", err)
		return 0, false
	}

	contentChar, free := gl.Strs(string(f) + "\x00")
	defer free()

	// Create a shader of the specified type
	outShader := gl.CreateShader(shaderType)
	// Set the source characters and try to compile
	gl.ShaderSource(outShader, 1, contentChar, nil)
	gl.CompileShader(outShader)

	if !s.isCompiled(outShader) {
		sdl.Log("Failed to compile shader %s", fileName)
		return 0, false
	}

	return outShader, true
}

func (s *Shader) isCompiled(shader uint32) bool {
	var status int32
	// Query the compile status
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)

	if status != gl.TRUE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		sdl.Log("GLSL Compile Failed:\n%s", log)
		return false
	}

	return true
}

func (s *Shader) isValidProgram() bool {
	var status int32
	// Query the link status
	gl.GetProgramiv(s.shaderProgram, gl.LINK_STATUS, &status)

	if status != gl.TRUE {
		var logLength int32
		gl.GetProgramiv(s.shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.shaderProgram, logLength, nil, gl.Str(log))
		sdl.Log("GLSL Link Status:\n%s", log)
		return false
	}

	return true
}
