package infra

import (
	"sort"
	"sync"
)

type Registry map[string]InfraInfo

type InfraInfo struct {
	Code      string
	Title     string
	Processor Processor
}

var registry = Registry{}
var registryMu sync.Mutex

func addInfra(code string, info InfraInfo) {
	registryMu.Lock()
	defer registryMu.Unlock()

	registry[code] = info
}

func GetInfra(code string) InfraInfo {
	registryMu.Lock()
	defer registryMu.Unlock()

	return registry[code]
}

func ListInfraInfos() []InfraInfo {
	registryMu.Lock()
	defer registryMu.Unlock()

	infos := make([]InfraInfo, 0, len(registry))
	for _, info := range registry {
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Code < infos[j].Code
	})

	return infos
}
