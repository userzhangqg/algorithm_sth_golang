package day07

/*
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。

*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	oldColor := image[sr][sc]
	//if image[sr][sc] != newColor {
	//	image[sr][sc] = newColor
	//}

	//if sr > 0 && oldColor == image[sr - 1][sc]{
	//	floodFill(image, sr - 1, sc, newColor)
	//}
	//if sr < len(image) - 1 && oldColor == image[sr + 1][sc]{
	//	floodFill(image, sr + 1, sc, newColor)
	//}
	//if sc > 0 && oldColor == image[sr][sc - 1]{
	//	floodFill(image, sr, sc - 1, newColor)
	//}
	//if sc < len(image[0]) - 1 && oldColor == image[sr][sc + 1]{
	//	floodFill(image, sr, sc + 1, newColor)
	//}
	if oldColor != newColor {
		dfs(image, sr, sc, oldColor, newColor)
	}

	//dfs(image, sr, sc, oldColor, newColor)

	return image
}

func dfs(image [][]int, sr int, sc int, color int, newColor int) {
	if image[sr][sc] == color {
		image[sr][sc] = newColor
		if sr > 0 {
			dfs(image, sr-1, sc, color, newColor)
		}
		if sr < len(image)-1 {
			dfs(image, sr+1, sc, color, newColor)
		}
		if sc > 0 {
			dfs(image, sr, sc-1, color, newColor)
		}
		if sc < len(image[0])-1 {
			dfs(image, sr, sc+1, color, newColor)
		}
	}
}
