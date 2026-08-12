package P207

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "无环可完成",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "存在环不可完成",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "无先修课程",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "单课程无先修",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "零课程",
			numCourses:    0,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "链式依赖无环",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "链式依赖有环",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}},
			want:          false,
		},
		{
			name:          "多分支无环",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {4, 2}},
			want:          true,
		},
		{
			name:          "多分支部分成环",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {4, 2}, {2, 4}},
			want:          false,
		},
		{
			name:          "自环依赖",
			numCourses:    3,
			prerequisites: [][]int{{1, 1}},
			want:          false,
		},
		{
			name:          "重复依赖无环",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {1, 0}, {2, 1}},
			want:          true,
		},
		{
			name:          "经典示例可完成",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish(%d, %v) = %v, want %v",
					tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
