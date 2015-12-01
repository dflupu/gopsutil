// +build windows

package load

import (
	"github.com/dflupu/gopsutil/internal/common"
)

func LoadAvg() (*LoadAvgStat, error) {
	ret := LoadAvgStat{}

	return &ret, common.NotImplementedError
}
