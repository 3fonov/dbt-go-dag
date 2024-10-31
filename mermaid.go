package main

import (
	"fmt"
	"strings"
)

func (m *WritableManifest) CreateMermaidFCGraph() string {
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
	for _, n := range m.Nodes {
		if n.ResourceType != "model" {
			continue
		}
		n.ModelToMermaidFC(b, tests)
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

		b.WriteString(fmt.Sprintf("    subgraph %v\n", name))
		for _, s := range sources {
			s.SourceToMermaidFC(b)
		}
		b.WriteString("    end\n")
		b.WriteString(
			fmt.Sprintf("    style %v fill:#FFF1E6,stroke:#333,stroke-width:4px\n", name),
		)
	}
}
func (s *Source) SourceToMermaidFC(b *strings.Builder) {
	id := ToMermaidId(s.UniqueID)
	b.WriteString(fmt.Sprintf("        %v>%v]:::source", id, s.Name))
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
	if _, exists := tests[id]; exists {
		modelName = fmt.Sprintf("fa:fa-check-double %v", modelName)
	} else {
		modelClass = "modelError"
	}
	if n.Config.Materialized == "ephemeral" {
		modelName = fmt.Sprintf("fa:fa-ghost %v", modelName)
		modelClass = "modelEphemeral"
	} else if n.Config.Materialized == "incremental" {
		modelName = fmt.Sprintf("fa:fa-layer-group %v", modelName)
	} else if n.Config.Materialized == "view" {
		modelName = fmt.Sprintf("fa:fa-eye %v", modelName)
	}
	if n.IsProductNode() {

		modelClass = "modelProduct"
		b.WriteString(
			fmt.Sprintf("    %v[%v fa:fa-box-open]:::%v", id, modelName, modelClass),
		)
	} else {
		b.WriteString(fmt.Sprintf("    %v([%v]):::%v", id, modelName, modelClass))
	}
	b.WriteString("\n")
}
func (n *Node) RefsToMermaidFC(b *strings.Builder) {
	for _, d := range n.DependsOn.Nodes {
		b.WriteString("        ")
		b.WriteString(
			fmt.Sprintf("%v --> %v\n", ToMermaidId(d), ToMermaidId(n.UniqueID)),
		)
	}
}
func (m *WritableManifest) CreateMermaidERGraph() string {
	var builder strings.Builder
	builder.WriteString("erDiagram\n")
	for _, n := range m.Nodes {
		if n.ResourceType != "model" {
			continue
		}
		n.ToMermaidER(&builder)
	}
	return builder.String()
}
func (n *Node) ToMermaidER(b *strings.Builder) {
	id := ToMermaidId(n.UniqueID)
	b.WriteString(fmt.Sprintf("    %v[%v]", id, n.Name))
	n.ColumnsToMermaidER(b)
	b.WriteString("\n")
	n.RefsToMermaidER(b)
}
func ToMermaidId(s string) string {
	return strings.Replace(s, ".", "_", -1)
}
func (n *Node) RefsToMermaidER(b *strings.Builder) {
	for _, d := range n.DependsOn.Nodes {
		b.WriteString("        ")
		b.WriteString(
			fmt.Sprintf("%v }o--o{ %v: \"\"\n", ToMermaidId(d), ToMermaidId(n.UniqueID)),
		)
	}
}
func (n *Node) ColumnsToMermaidER(b *strings.Builder) {
	if len(n.Columns) > 0 {
		b.WriteString(" {\n")
		for _, c := range n.Columns {
			c.ToMermaidER(b)
		}
		b.WriteString("}")
	}
}

func (c *ColumnInfo) ToMermaidER(b *strings.Builder) {
	b.WriteString("        ")
	if c.DataType == nil {
		b.WriteString("unknown ")
	} else {
		b.WriteString(fmt.Sprintf("%v ", *c.DataType))
	}
	b.WriteString(fmt.Sprintf("%v", c.Name))
	b.WriteString("\n")
}
