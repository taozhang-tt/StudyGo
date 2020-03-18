package main

import "fmt"

/**
836. 矩形重叠
	https://leetcode-cn.com/problems/rectangle-overlap/
题目描述
	矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。
	如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。
	给出两个矩形，判断它们是否重叠并返回结果。
示例 1：
	输入：rec1 = [0,0,2,2], rec2 = [1,1,3,3]
	输出：true
*/
func main() {
	rec1 := []int{0,0,2,2}
	rec2 := []int{1,1,3,3}
	fmt.Println(isRectangleOverlap1(rec1, rec2))
}


/**
检查位置
	固定 rect1，若想两个矩形无交集，rect2 只能出现在 rect1 四周
	rect1 在 rect2 左侧时，存在一条平行于y轴的线 x，满足 x >= rect1.x2 && x <= rect2.x1
	rect1 在 rect2 右侧时，存在一条平行于y轴的线 x，满足 x <= rect1.x1 && x >= rect2.x2
	rect1 在 rect2 上侧时，存在一条平行于x轴的线 y，满足 y <= rect1.y1 && x >= rect2.y2
	rect1 在 rect2 下侧时，存在一条平行于x轴的线 y，满足 y >= rect1.y2 && y <= rect2.y1
*/
func isRectangleOverlap1(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || //左
		rec1[0] >= rec2[2] || //右
		rec1[1] >= rec2[3] || //上
		rec2[1] >= rec1[3]) //下
//化简为
//	return rec1[1] < rec2[3] && rec1[3] > rec2[1] && rec1[0] < rec2[2] && rec1[2] > rec2[0]
}

/**
投影
	如果两个矩形有交集，那将其投影到y轴上，两个矩形对应的两条直线必然有交集，同时投影到x轴上的两条直线也必有交集
	投影到y轴，rect1 => (rect1.y1, rect1.y2), rect2 => (rect2.y1, rect2.y2)
	投影到x轴，rect1 => (rect1.x1, rect1.x2), rect2 => (rect2.x1, rect2.x2)
	y轴有交集表示为
		!(rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
	x轴有交集表示为
		!(rec1[0] >= rec2[2] || rec1[2] <= rec2[0])
	化简为
		rec1[1] < rec2[3] && rec1[3] > rec2[1] && rec1[0] < rec2[2] && rec1[2] > rec2[0]
*/
func isRectangleOverlap2(rec1 []int, rec2 []int) bool {
	return rec1[1] < rec2[3] && rec1[3] > rec2[1] && rec1[0] < rec2[2] && rec1[2] > rec2[0]
}
