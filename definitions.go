package main

type WritableManifest struct {
	Metadata  ManifestMetadata    `json:"metadata"`
	Nodes     map[string]Node     `json:"nodes"`
	Sources   map[string]Source   `json:"sources"`
	Exposures map[string]Exposure `json:"exposures"`
	sourceMap map[string]string
}

type ManifestMetadata struct {
	ProjectName *string `json:"project_name,omitempty"`
	ProjectID   *string `json:"project_id,omitempty"`
}
type Node struct {
	Database     *string                `json:"database,omitempty"`
	Schema       string                 `json:"schema"`
	Name         string                 `json:"name"`
	ResourceType string                 `json:"resource_type"` // Example: "model", "analysis", "test", etc.
	PackageName  string                 `json:"package_name"`
	Path         string                 `json:"path"`
	UniqueID     string                 `json:"unique_id"`
	Alias        *string                `json:"alias,omitempty"`
	Config       NodeConfig             `json:"config"`
	Tags         []string               `json:"tags"`
	Description  string                 `json:"description"`
	Meta         map[string]interface{} `json:"meta"`
	CreatedAt    float64                `json:"created_at"`
	RelationName *string                `json:"relation_name,omitempty"`
	Language     string                 `json:"language"`
	Refs         []RefArgs              `json:"refs"`
	Sources      [][]string             `json:"sources"`
	Metrics      [][]string             `json:"metrics"`
	DependsOn    DependsOn              `json:"depends_on"`
	AttachedNode string                 `json:"attached_node"`
}

type NodeConfig struct {
	Enabled             bool                   `json:"enabled"`
	Alias               *string                `json:"alias,omitempty"`
	Schema              *string                `json:"schema,omitempty"`
	Database            *string                `json:"database,omitempty"`
	Tags                []string               `json:"tags"`
	Meta                map[string]interface{} `json:"meta"`
	Group               *string                `json:"group,omitempty"`
	Materialized        string                 `json:"materialized"`
	IncrementalStrategy *string                `json:"incremental_strategy,omitempty"`
	ColumnTypes         map[string]string      `json:"column_types"`
	Packages            []string               `json:"packages"`
}

type DependsOn struct {
	Macros []string `json:"macros"`
	Nodes  []string `json:"nodes"`
}
type RefArgs struct {
	Name    string  `json:"name"`
	Package *string `json:"package,omitempty"`
	Version *string `json:"version,omitempty"`
}
type Source struct {
	Name        string                 `json:"name"`               // Source name
	Description string                 `json:"description"`        // Description of the source
	Meta        map[string]interface{} `json:"meta"`               // Additional metadata
	Identifier  string                 `json:"identifier"`         // Identifier for the source
	Schema      string                 `json:"schema"`             // Schema where the source is located
	SourceName  string                 `json:"source_name"`        // Schema where the source is located
	Database    *string                `json:"database,omitempty"` // Database for the source (optional)
	Loader      *string                `json:"loader,omitempty"`   // Loader for the source (optional)
	Tags        []string               `json:"tags"`               // Tags for categorization
	Columns     map[string]ColumnInfo  `json:"columns"`            // Column metadata for the source
	Tables      []Table                `json:"tables"`             // List of tables within the source
	UniqueID    string                 `json:"unique_id"`          // Unique identifier for the source
}

type Table struct {
	Name        string                 `json:"name"`        // Table name
	Description string                 `json:"description"` // Description of the table
	Meta        map[string]interface{} `json:"meta"`        // Additional metadata for the table
	Identifier  string                 `json:"identifier"`  // Identifier for the table
	Tags        []string               `json:"tags"`        // Tags for the table
	Columns     map[string]ColumnInfo  `json:"columns"`     // Column metadata for the table
	UniqueID    string                 `json:"unique_id"`   // Unique identifier for the table
}

type ColumnInfo struct {
	Name        string                 `json:"name"`                  // Column name
	Description string                 `json:"description"`           // Column description
	Meta        map[string]interface{} `json:"meta"`                  // Additional metadata
	DataType    string                 `json:"data_type,omitempty"`   // Column data type
	Tags        []string               `json:"tags"`                  // Tags for the column
	Granularity string                 `json:"granularity,omitempty"` // Column granularity level
}
type Macro struct {
	Name    string            `json:"name"`
	Package string            `json:"package"`
	Meta    map[string]string `json:"meta"`
	Path    string            `json:"path"`
}
type Exposure struct {
	Label       string    `json:"label"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Maturity    string    `json:"maturity"`
	Owner       Owner     `json:"owner"`
	DependsOn   DependsOn `json:"depends_on"`
}

type Owner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
