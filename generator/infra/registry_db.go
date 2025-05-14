package infra

import (
	"math"
	"sort"
	"sync"
)

type RegistryDB map[string]DBInfo

type DBInfo struct {
	Code      string
	Title     string
	Processor Processor

	order int
}

var registryDB = RegistryDB{}
var registryDBMu sync.Mutex

func addDB(code string, info DBInfo) {
	registryDBMu.Lock()
	defer registryDBMu.Unlock()

	if info.order == 0 {
		info.order = math.MaxInt
	}

	registryDB[code] = info
}

func GetDB(code string) DBInfo {
	registryDBMu.Lock()
	defer registryDBMu.Unlock()

	return registryDB[code]
}

func ListDBInfos() []DBInfo {
	registryDBMu.Lock()
	defer registryDBMu.Unlock()

	infos := make([]DBInfo, 0, len(registryDB))
	for _, info := range registryDB {
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].order < infos[j].order
	})

	return infos
}
