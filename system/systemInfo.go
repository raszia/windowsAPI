package system

import (
	"context"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/host"
)

func GetInfo(ctx context.Context) (*host.InfoStat, error) {
	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func Getvmem(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	vmem, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return vmem, nil
}

func GetInterfaces(ctx context.Context) ([]net.InterfaceStat, error) {
	iStat, err := net.InterfacesWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return iStat, nil
}

func GetDisk(ctx context.Context) ([]net.InterfaceStat, error) {
	iStat, err := net.InterfacesWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return iStat, nil
}
