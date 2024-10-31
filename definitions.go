package main

type WritableManifest struct {
	Metadata  ManifestMetadata    `json:"metadata"`
	Nodes     map[string]Node     `json:"nodes"`
	Sources   map[string]Source   `json:"sources"`
	Macros    map[string]Macro    `json:"macros"`
	Exposures map[string]Exposure `json:"exposures"`
}

type ManifestMetadata struct {
	DBTSchemaVersion        string            `json:"dbt_schema_version"`
	DBTVersion              string            `json:"dbt_version"`
	GeneratedAt             string            `json:"generated_at"`
	InvocationID            *string           `json:"invocation_id,omitempty"`
	Env                     map[string]string `json:"env"`
	ProjectName             *string           `json:"project_name,omitempty"`
	ProjectID               *string           `json:"project_id,omitempty"`
	UserID                  *string           `json:"user_id,omitempty"`
	SendAnonymousUsageStats *bool             `json:"send_anonymous_usage_stats,omitempty"`
	AdapterType             *string           `json:"adapter_type,omitempty"`
}
type Node struct {
	Database          *string                `json:"database,omitempty"`
	Schema            string                 `json:"schema"`
	Name              string                 `json:"name"`
	ResourceType      string                 `json:"resource_type"` // Example: "model", "analysis", "test", etc.
	PackageName       string                 `json:"package_name"`
	Path              string                 `json:"path"`
	OriginalFilePath  string                 `json:"original_file_path"`
	UniqueID          string                 `json:"unique_id"`
	FQN               []string               `json:"fqn"`
	Alias             *string                `json:"alias,omitempty"`
	Checksum          FileHash               `json:"checksum"`
	Config            NodeConfig             `json:"config"`
	Tags              []string               `json:"tags"`
	Description       string                 `json:"description"`
	Columns           map[string]ColumnInfo  `json:"columns"`
	Meta              map[string]interface{} `json:"meta"`
	Group             *string                `json:"group,omitempty"`
	Docs              Docs                   `json:"docs"`
	PatchPath         *string                `json:"patch_path,omitempty"`
	BuildPath         *string                `json:"build_path,omitempty"`
	UnrenderedConfig  map[string]interface{} `json:"unrendered_config"`
	CreatedAt         float64                `json:"created_at"`
	RelationName      *string                `json:"relation_name,omitempty"`
	RawCode           string                 `json:"raw_code"`
	Language          string                 `json:"language"`
	Refs              []RefArgs              `json:"refs"`
	Sources           [][]string             `json:"sources"`
	Metrics           [][]string             `json:"metrics"`
	DependsOn         DependsOn              `json:"depends_on"`
	CompiledPath      *string                `json:"compiled_path,omitempty"`
	Compiled          bool                   `json:"compiled"`
	CompiledCode      *string                `json:"compiled_code,omitempty"`
	ExtraCTEsInjected bool                   `json:"extra_ctes_injected"`
	ExtraCTEs         []InjectedCTE          `json:"extra_ctes"`
	PreInjectedSQL    *string                `json:"_pre_injected_sql,omitempty"`
	Contract          *Contract              `json:"contract,omitempty"`
	AttachedNode      string                 `json:"attached_node"`
}

type FileHash struct {
	Name     string `json:"name"`
	Checksum string `json:"checksum"`
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
	PostHook            []Hook                 `json:"post-hook"`
	PreHook             []Hook                 `json:"pre-hook"`
	Quoting             map[string]interface{} `json:"quoting"`
	ColumnTypes         map[string]string      `json:"column_types"`
	FullRefresh         *bool                  `json:"full_refresh,omitempty"`
	OnSchemaChange      *string                `json:"on_schema_change,omitempty"`
	Grants              map[string]string      `json:"grants"`
	Packages            []string               `json:"packages"`
	Docs                Docs                   `json:"docs"`
	Contract            ContractConfig         `json:"contract"`
}

type Hook struct {
	SQL         string `json:"sql"`
	Transaction bool   `json:"transaction"`
	Index       *int   `json:"index,omitempty"`
}

type ColumnLevelConstraint struct {
	Type            string   `json:"type"`
	Name            *string  `json:"name,omitempty"`
	Expression      *string  `json:"expression,omitempty"`
	WarnUnenforced  bool     `json:"warn_unenforced"`
	WarnUnsupported bool     `json:"warn_unsupported"`
	To              *string  `json:"to,omitempty"`
	ToColumns       []string `json:"to_columns"`
}

type Docs struct {
	Show      bool    `json:"show"`
	NodeColor *string `json:"node_color,omitempty"`
}

type Contract struct {
	Enforced   bool    `json:"enforced"`
	AliasTypes bool    `json:"alias_types"`
	Checksum   *string `json:"checksum,omitempty"`
}

type DependsOn struct {
	Macros []string `json:"macros"`
	Nodes  []string `json:"nodes"`
}

type InjectedCTE struct {
	ID  string `json:"id"`
	SQL string `json:"sql"`
}

type RefArgs struct {
	Name    string  `json:"name"`
	Package *string `json:"package,omitempty"`
	Version *string `json:"version,omitempty"`
}

type ContractConfig struct {
	Enforced   bool `json:"enforced"`
	AliasTypes bool `json:"alias_types"`
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
	Quoting     Quoting                `json:"quoting"`            // Quoting information
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
	Quoting     Quoting                `json:"quoting"`     // Quoting information for the table
	Columns     map[string]ColumnInfo  `json:"columns"`     // Column metadata for the table
	UniqueID    string                 `json:"unique_id"`   // Unique identifier for the table
}

type Quoting struct {
	Database   bool `json:"database"`   // Whether to quote database names
	Schema     bool `json:"schema"`     // Whether to quote schema names
	Identifier bool `json:"identifier"` // Whether to quote identifiers
}

type ColumnInfo struct {
	Name        string                  `json:"name"`                  // Column name
	Description string                  `json:"description"`           // Column description
	Meta        map[string]interface{}  `json:"meta"`                  // Additional metadata
	DataType    *string                 `json:"data_type,omitempty"`   // Column data type
	Constraints []ColumnLevelConstraint `json:"constraints"`           // Constraints on the column
	Quote       *bool                   `json:"quote,omitempty"`       // Whether to quote the column name
	Tags        []string                `json:"tags"`                  // Tags for the column
	Granularity *string                 `json:"granularity,omitempty"` // Column granularity level
}
type Macro struct {
	Name    string            `json:"name"`
	Package string            `json:"package"`
	SQL     string            `json:"sql"`
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
