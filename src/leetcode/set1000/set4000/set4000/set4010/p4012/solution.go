package p4012

import "sort"

func countTasks(tasks []int, shifts []int) []int {
	n := len(tasks)
	sum := make([]int, n+1)
	for i, v := range tasks {
		sum[i+1] = sum[i] + v
	}

	var pending int
	var taskId int

	play := func(tot int) int {
		if pending > tot {
			pending -= tot
			// taskId not finished yet
			return n - taskId
		}
		if pending > 0 {
			if tot < pending {
				pending -= tot
				return n - taskId
			}
			// tot >= pending
			tot -= pending
			taskId++
		}

		if sum[n]-sum[taskId] <= tot {
			pending = 0
			taskId = 0
			return 0
		}

		r := sort.Search(n, func(i int) bool {
			return sum[i]-sum[taskId] > tot
		})
		// sum[r] - sum[taskId] > tot
		pending = tasks[r-1] - (tot - sum[r-1] + sum[taskId])
		taskId = r - 1
		return n - taskId
	}

	ans := make([]int, len(shifts))

	for i, cur := range shifts {
		ans[i] = play(cur)
	}

	return ans
}
