package infra

import (
	"math"
	"sort"

	"github.com/rom8726/airo/config"
)

type Registry struct {
	dbs    map[string]*DBInfo
	infras map[string]*InfraInfo
}

type DBInfo struct {
	Code      string
	Title     string
	Processor Processor

	order int
}

type InfraInfo struct {
	Code      string
	Title     string
	Processor Processor

	order int
}

type Opt func(*Registry)

func NewRegistry(opts ...Opt) *Registry {
	reg := &Registry{
		dbs:    map[string]*DBInfo{},
		infras: map[string]*InfraInfo{},
	}

	for _, opt := range opts {
		opt(reg)
	}

	return reg
}

func (r *Registry) UpdateConfig(cfg *config.ProjectConfig) {
	for _, info := range r.ListDBs() {
		info.Processor.SetConfig(cfg)
	}
	for _, info := range r.ListInfras() {
		info.Processor.SetConfig(cfg)
	}
}

func (r *Registry) ListDBs() []*DBInfo {
	list := make([]*DBInfo, 0, len(r.dbs))
	for _, info := range r.dbs {
		list = append(list, info)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].order < list[j].order
	})

	return list
}

func (r *Registry) ListInfras() []*InfraInfo {
	list := make([]*InfraInfo, 0, len(r.infras))
	for _, info := range r.infras {
		list = append(list, info)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].order < list[j].order
	})

	return list
}

func (r *Registry) GetDB(code string) DBInfo {
	return *r.dbs[code]
}

func (r *Registry) GetInfra(code string) InfraInfo {
	return *r.infras[code]
}

func (r *Registry) addDB(code string, info *DBInfo) {
	if info.order == 0 {
		info.order = math.MaxInt
	}

	r.dbs[code] = info
}

func (r *Registry) addInfra(code string, info *InfraInfo) {
	if info.order == 0 {
		info.order = math.MaxInt
	}

	r.infras[code] = info
}
