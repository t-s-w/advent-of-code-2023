package d05

import "testing"

type sampleCase struct {
	source int
	expectedDest int
}

type mapTestCase struct {
	M ResourceMap
	cases []sampleCase
}
func TestResourceMaps(t *testing.T) {
	cases := []mapTestCase{
		{
			M: ResourceMap{
				Data: []mapData{
					{49, 53, 8},
					{0, 11, 42},
					{42, 0, 7},
					{57, 7, 4},
				},
			},
			cases: []sampleCase{
				{source: 2, expectedDest: 44},
				{source: 7, expectedDest: 57},
				{source: 60, expectedDest: 56},
			},
		},
	}

	for i, c := range cases {
		for j, d := range c.cases {
			got := c.M.Map(d.source)
			if got != d.expectedDest {
				t.Errorf("Testcase ID %d:%d failed: expected %d got %d",i,j,got,d.expectedDest)
			}
		} 
	}
}