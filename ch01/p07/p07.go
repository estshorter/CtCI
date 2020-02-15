package p07

// RotateMatrix rotate matrix
func RotateMatrix(mat [][]int, n int, positiveRotation bool) [][]int {
	for layer := 0; layer < n/2; layer++ {
		for i := 0; i < n-2*layer-1; i++ {
			top := &mat[layer][layer+i]
			right := &mat[layer+i][n-1-layer]
			bottom := &mat[n-1-layer][n-i-1-layer]
			left := &mat[n-1-i-layer][layer]

			if positiveRotation {
				tmp := *top
				*top = *left
				*left = *bottom
				*bottom = *right
				*right = tmp
			} else {
				tmp := *top
				*top = *right
				*right = *bottom
				*bottom = *left
				*left = tmp
			}
		}
	}
	return mat
}
