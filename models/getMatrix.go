package models

const (
	//BasicWidth is a default map width
	BasicWidth = 9
	//BasicHeight is a default map height
	BasicHeight = 9
	//BasicBlockWidth is a default blockwith
	BasicBlockWidth = 3
)

//GetMatrix is a simple constructor for Matrix object
func GetMatrix() Matrix {
	var matrix Matrix = Matrix{
		Width:      BasicWidth,
		Height:     BasicHeight,
		BlockWidth: BasicBlockWidth,
	}
	matrix.GenerateMap()
	return matrix
}
