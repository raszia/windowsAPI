package system

import (
	"context"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/disk"
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

func GetDiskUsage(ctx context.Context, path string) (*disk.UsageStat, error) {
	dStat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return dStat, nil
}
func GetDiskPartitions(all bool) ([]disk.PartitionStat, error) {
	dStat, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}
	return dStat, nil
}

func GetDiskIocounters(names ...string) (map[string]disk.IOCountersStat, error) {
	dStat, err := disk.IOCounters(names...)
	if err != nil {
		return nil, err
	}
	return dStat, nil
}
