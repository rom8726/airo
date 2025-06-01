package infra

import (
	"math"
	"sort"

	"github.com/rom8726/airo/config"
)

// ComponentType represents the type of a component (DB or Infra)
type ComponentType int

const (
	// DBComponent represents a database component
	DBComponent ComponentType = iota
	// InfraComponent represents an infrastructure component
	InfraComponent
)

// ComponentInfo contains information about a component
type ComponentInfo struct {
	Code      string
	Title     string
	Type      ComponentType
	Processor Processor
	Order     int
}

// Registry manages the available components
type Registry struct {
	dbs    map[string]*ComponentInfo
	infras map[string]*ComponentInfo
}

// RegistryOption is a function that configures a Registry
type RegistryOption func(*Registry)

// NewRegistry creates a new Registry with the given options
func NewRegistry(opts ...RegistryOption) *Registry {
	reg := &Registry{
		dbs:    map[string]*ComponentInfo{},
		infras: map[string]*ComponentInfo{},
	}

	for _, opt := range opts {
		opt(reg)
	}

	return reg
}

// WithComponent adds a component to the registry
func WithComponent(info ComponentInfo) RegistryOption {
	return func(r *Registry) {
		if info.Order == 0 {
			info.Order = math.MaxInt
		}

		infoCopy := info // Make a copy to avoid modifying the original

		switch info.Type {
		case DBComponent:
			r.dbs[info.Code] = &infoCopy
		case InfraComponent:
			r.infras[info.Code] = &infoCopy
		}
	}
}

// WithDB is a convenience function for adding a database component
func WithDB(code, title string, processor Processor, order int) RegistryOption {
	return WithComponent(ComponentInfo{
		Code:      code,
		Title:     title,
		Type:      DBComponent,
		Processor: processor,
		Order:     order,
	})
}

// WithInfra is a convenience function for adding an infrastructure component
func WithInfra(code, title string, processor Processor, order int) RegistryOption {
	return WithComponent(ComponentInfo{
		Code:      code,
		Title:     title,
		Type:      InfraComponent,
		Processor: processor,
		Order:     order,
	})
}

// RegisterComponent registers a component with the registry
func (r *Registry) RegisterComponent(info ComponentInfo) {
	if info.Order == 0 {
		info.Order = math.MaxInt
	}

	switch info.Type {
	case DBComponent:
		r.dbs[info.Code] = &info
	case InfraComponent:
		r.infras[info.Code] = &info
	}
}

// RegisterDB registers a database component
func (r *Registry) RegisterDB(code, title string, processor Processor, order int) {
	r.RegisterComponent(ComponentInfo{
		Code:      code,
		Title:     title,
		Type:      DBComponent,
		Processor: processor,
		Order:     order,
	})
}

// RegisterInfra registers an infrastructure component
func (r *Registry) RegisterInfra(code, title string, processor Processor, order int) {
	r.RegisterComponent(ComponentInfo{
		Code:      code,
		Title:     title,
		Type:      InfraComponent,
		Processor: processor,
		Order:     order,
	})
}

// UpdateConfig updates the configuration for all components
func (r *Registry) UpdateConfig(cfg *config.ProjectConfig) {
	for _, info := range r.ListDBs() {
		info.Processor.SetConfig(cfg)
	}
	for _, info := range r.ListInfras() {
		info.Processor.SetConfig(cfg)
	}
}

// ListDBs returns a sorted list of database components
func (r *Registry) ListDBs() []*ComponentInfo {
	list := make([]*ComponentInfo, 0, len(r.dbs))
	for _, info := range r.dbs {
		list = append(list, info)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Order < list[j].Order
	})

	return list
}

// ListInfras returns a sorted list of infrastructure components
func (r *Registry) ListInfras() []*ComponentInfo {
	list := make([]*ComponentInfo, 0, len(r.infras))
	for _, info := range r.infras {
		list = append(list, info)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Order < list[j].Order
	})

	return list
}

// GetDB returns information about a database component
func (r *Registry) GetDB(code string) ComponentInfo {
	return *r.dbs[code]
}

// GetInfra returns information about an infrastructure component
func (r *Registry) GetInfra(code string) ComponentInfo {
	return *r.infras[code]
}

// For backward compatibility
type DBInfo = ComponentInfo
type InfraInfo = ComponentInfo
type Opt = RegistryOption

// For backward compatibility
func (r *Registry) addDB(code string, info *DBInfo) {
	if info.Order == 0 {
		info.Order = math.MaxInt
	}
	info.Type = DBComponent
	r.dbs[code] = info
}

// For backward compatibility
func (r *Registry) addInfra(code string, info *InfraInfo) {
	if info.Order == 0 {
		info.Order = math.MaxInt
	}
	info.Type = InfraComponent
	r.infras[code] = info
}
