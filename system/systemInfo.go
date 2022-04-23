package system

import (
	"context"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"
)

func GetInfo(ctx context.Context) (*host.InfoStat, error) {
	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func Getvmem(ctx context.Context) (*mem.VirtualMemoryExStat, error) {
	vmem, err := mem.VirtualMemoryExWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return vmem, nil
}
