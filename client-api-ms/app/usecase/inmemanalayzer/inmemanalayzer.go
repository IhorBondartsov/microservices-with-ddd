package inmemanalayzer

import (
	"github.com/shirou/gopsutil/mem"
)

type MemoryAnalyzer interface{
	GetTotalAndFreeMemory() (total uint64, free uint64, err error)
}

type memoryAnalyser struct{
	stat *mem.VirtualMemoryStat
}

func NewMemoryAnalyser() MemoryAnalyzer {
	return &memoryAnalyser{}
}

func (ma *memoryAnalyser)GetTotalAndFreeMemory() (total uint64, free uint64, err error){
	m, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}
	return m.Total, m.Free, nil
}

