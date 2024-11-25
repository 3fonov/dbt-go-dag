package main

import (
	"fmt"
	"strings"
)

var idDict map[string]string

func (m *WritableManifest) CreateMermaidFCGraph() string {
	idDict = make(map[string]string, 0)
	var builder strings.Builder
	builder.WriteString(
		"%%{init: {'flowchart': {'defaultRenderer': 'elk', 'curve':'step' }} }%%\n",
	)
	builder.WriteString("flowchart LR\n")
	builder.WriteString(
		"     classDef default stroke:#383D3B,stroke-width:2px,text-align:left,white-space:nowrap;\n",
	)
	builder.WriteString("     classDef source fill:#EAE4E9,stroke-width:1px\n")
	builder.WriteString("     classDef model fill:#BEE1E6\n")
	builder.WriteString(
		"     classDef modelError fill:#E93577,stroke:#370618,color:#fefefe\n",
	)
	builder.WriteString("     classDef modelProduct fill:#DFE7FD\n")
	builder.WriteString("     classDef modelProductError fill:#E93577,color:#FEFEFE\n")
	builder.WriteString("     classDef exposure fill:#FFF1E6\n")
	builder.WriteString(
		"     classDef modelEphemeral fill:#fafafa,stroke-width:1px,stroke:#cdcdcd\n",
	)
	m.SourcesToMermaidFC(&builder)
	tests := m.GetTests()
	m.ModelsToMermaidFC(&builder, tests)
	m.ExposuresToMermaidFC(&builder)
	m.LinksToMermaidFC(&builder)
	return builder.String()
}
func (m *WritableManifest) GetTests() map[string]bool {
	result := make(map[string]bool)
	for _, n := range m.Nodes {
		if n.ResourceType != "test" {
			continue
		}
		id := ToMermaidId(n.AttachedNode)
		if _, exists := result[id]; !exists {
			result[id] = true
		}
	}
	return result
}
func (m *WritableManifest) ExposuresToMermaidFC(b *strings.Builder) {
	for e_id, e := range m.Exposures {
		id := ToMermaidId(e_id)
		b.WriteString(
			fmt.Sprintf(
				"        %v:::exposure@{ shape: curv-trap, label: \"%v\" }\n",
				id,
				e.Label,
			),
		)
	}
}
func (m *WritableManifest) LinksToMermaidFC(b *strings.Builder) {
	for _, n := range m.Nodes {
		if n.ResourceType != "model" {
			continue
		}
		n.RefsToMermaidFC(b)
	}
	for e_id, e := range m.Exposures {
		for _, d := range e.DependsOn.Nodes {
			b.WriteString("        ")
			b.WriteString(
				fmt.Sprintf("%v --> %v\n", ToMermaidId(d), ToMermaidId(e_id)),
			)
		}
	}
}
func (m *WritableManifest) ModelsToMermaidFC(b *strings.Builder, tests map[string]bool) {
	modelGroups := make(map[string][]Node)
	for _, n := range m.Nodes {
		if n.ResourceType != "model" {
			continue
		}
		groupName := n.GroupName()
		if _, exists := modelGroups[groupName]; exists {
			modelGroups[groupName] = append(modelGroups[groupName], n)
		} else {
			modelGroups[groupName] = []Node{n}
		}
	}

	for name, nodes := range modelGroups {
		if name == "-" {
			for _, n := range nodes {
				n.ModelToMermaidFC(b, tests)
			}
		} else {
			b.WriteString(fmt.Sprintf("    subgraph %v\n", name))
			for _, n := range nodes {
				n.ModelToMermaidFC(b, tests)
			}
			b.WriteString("    end\n")
			color := GetStringColor(name)
			b.WriteString(
				fmt.Sprintf("    style %v fill:%v,stroke:#333,stroke-width:1px\n", name, color),
			)
		}
	}
}
func (m *WritableManifest) SourcesToMermaidFC(b *strings.Builder) {

	sourceGroups := make(map[string][]Source)
	for _, s := range m.Sources {
		if _, exists := sourceGroups[s.SourceName]; exists {
			sourceGroups[s.SourceName] = append(sourceGroups[s.SourceName], s)
		} else {
			sourceGroups[s.SourceName] = []Source{s}
		}
	}
	for name, sources := range sourceGroups {

		b.WriteString(fmt.Sprintf("    subgraph %v_[%v]\n", name))
		var sourceNameList []string

		for _, s := range sources {
			sourceNameList = append(sourceNameList, s.Name)
		}
		sourceMap := CollapseStrings(sourceNameList)
		for _, s := range sources {
			s.SourceToMermaidFC(b, sourceMap)
		}
		b.WriteString("    end\n")
		b.WriteString(
			fmt.Sprintf("    style %v fill:#FFF1E6,stroke:#333,stroke-width:1px\n", name),
		)
	}
}
func (s *Source) SourceToMermaidFC(b *strings.Builder, sm map[string]string) {
	id := ToMermaidId(s.UniqueID)
	name := s.Name
	if cName, ok := sm[s.Name]; ok {
		id = ToMermaidId(cName)
		idDict[s.UniqueID] = id
		name = cName
	}
	b.WriteString(fmt.Sprintf("        %v>%v]:::source", id, name))
	b.WriteString("\n")
}
func (n *Node) IsProductNode() bool {
	if n.PackageName == "product" {
		return true
	}
	for _, t := range n.Tags {
		if t == "product" {
			return true
		}
	}
	return false
}
func (n *Node) ModelToMermaidFC(b *strings.Builder, tests map[string]bool) {
	id := ToMermaidId(n.UniqueID)
	modelName := n.Name
	modelClass := "model"
	if _, exists := tests[id]; !exists {
		modelClass = "modelError"
	}
	if n.Config.Materialized == "ephemeral" {
		modelClass = "modelEphemeral"
	} else if n.Config.Materialized == "incremental" {
		modelName = fmt.Sprintf("fa:fa-layer-group %v", modelName)
	} else if n.Config.Materialized == "view" {
		modelName = fmt.Sprintf("fa:fa-eye %v", modelName)
	}
	if n.IsProductNode() {

		if modelClass == "modelError" {
			modelClass = "modelProductError"
		} else {
			modelClass = "modelProduct"
		}
		b.WriteString(
			fmt.Sprintf("    %v[%v]:::%v", id, modelName, modelClass),
		)
	} else {
		b.WriteString(fmt.Sprintf("    %v([%v]):::%v", id, modelName, modelClass))
	}
	b.WriteString("\n")
}
func (n *Node) RefsToMermaidFC(b *strings.Builder) {
	visited := make(map[string]interface{})
	for _, d := range n.DependsOn.Nodes {
		sourceId := d
		if cId, ok := idDict[sourceId]; ok {
			sourceId = cId
		}
		if _, vis := visited[sourceId]; vis {

			continue
		}
		visited[sourceId] = sourceId
		b.WriteString("        ")
		b.WriteString(
			fmt.Sprintf("%v --> %v\n", ToMermaidId(sourceId), ToMermaidId(n.UniqueID)),
		)
	}
}
func (n *Node) GroupName() string {
	groupPrefix := "group:"
	for _, t := range n.Config.Tags {
		if strings.HasPrefix(t, groupPrefix) {
			return strings.TrimSpace(
				strings.TrimPrefix(t, groupPrefix),
			)
		}

	}
	return "-"
}
func ToMermaidId(s string) string {
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "(", "_", -1)
	s = strings.Replace(s, ")", "_", -1)
	s = strings.Replace(s, "\"", "_", -1)
	return s

}
