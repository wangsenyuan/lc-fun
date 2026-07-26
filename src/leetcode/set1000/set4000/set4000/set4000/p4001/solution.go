package p4001

func aggregateTimeSeries(series1 [][]int, series2 [][]int) [][]int {
	var res [][]int

	for i, j := 0, 0; i < len(series1) || j < len(series2); {
		if i < len(series1) && j < len(series2) {
			if series1[i][0] == series2[j][0] {
				res = append(res, []int{series1[i][0], series1[i][1] + series2[j][1]})
				i++
				j++
			} else if series1[i][0] < series2[j][0] {
				res = append(res, []int{series1[i][0], series1[i][1] + series2[j][1]})
				i++
			} else {
				res = append(res, []int{series2[j][0], series2[j][1] + series1[i][1]})
				j++
			}
		} else if i < len(series1) {
			res = append(res, series1[i])
			i++
		} else {
			res = append(res, series2[j])
			j++
		}
	}
	return res
}
