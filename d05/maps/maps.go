package maaps

type MapData [3]int

type ResourceMap struct {
	Data []MapData
}

func (m *ResourceMap) Map(source int) (dest int) {
	for _, m := range m.Data {
		if source >= m[1] && source < m[1]+m[2] {
			return source - m[1] + m[0]
		}
	}
	return source
}