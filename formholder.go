package main

type FormHolder struct {
	Map map[string]interface{}
}

func (m *FormHolder) Fill(headArr, formArr []string) {
	m.Map = make(map[string]interface{})
	for i, v := range headArr {
		if i < len(formArr) {
			m.Map[v] = formArr[i]
		} else {
			m.Map[v] = ""
		}
	}
}

func (m *FormHolder) IsEmpty() bool {
	for _, v := range m.Map {
		if v.(string) != "" {
			return false
		}
	}
	return true
}
