package main

type TableNode struct {
	Id        string   `json:"key"`
	Name      string   `json:"name"`
	DependsOn []string `json:"depends_on"`
	Columns   []Column `json:"columns"`
}
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (m *WritableManifest) CreateGraph() []TableNode {
	result := []TableNode{}
	for _, n := range m.Nodes {
		if n.ResourceType != "model" {
			continue
		}
		node := TableNode{Name: n.Name, Id: n.UniqueID, DependsOn: []string{}}
		for _, c := range n.Columns {
			dataType := ""
			if c.DataType != nil {
				dataType = *c.DataType
			}
			col := Column{Name: c.Name, Type: dataType}
			node.Columns = append(node.Columns, col)
		}
		for _, r := range n.DependsOn.Nodes {
			node.DependsOn = append(node.DependsOn, r)
		}
		result = append(result, node)
	}
	return result
}
