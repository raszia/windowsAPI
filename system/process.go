package system

import (
	"context"

	"github.com/shirou/gopsutil/v3/process"
)

type NumCtxSwitchesStat struct {
	Voluntary   int64 `json:"voluntary"`
	Involuntary int64 `json:"involuntary"`
}
type MemoryInfoStat struct {
	RSS    uint64 `json:"rss"`    // bytes
	VMS    uint64 `json:"vms"`    // bytes
	HWM    uint64 `json:"hwm"`    // bytes
	Data   uint64 `json:"data"`   // bytes
	Stack  uint64 `json:"stack"`  // bytes
	Locked uint64 `json:"locked"` // bytes
	Swap   uint64 `json:"swap"`   // bytes
}

type Process struct {
	Pid            int32               `json:"pid"`
	Name           string              `json:"name"`
	Status         []string            `json:"status"`
	ParentPID      int32               `json:"parentPID"`
	Uids           []int32             `json:"uids"`
	Gids           []int32             `json:"gids"`
	Groups         []int32             `json:"groups"`
	NumThreads     int32               `json:"numThreads"`
	NumCtxSwitches *NumCtxSwitchesStat `json:"numCtxSwitches"`
	MemInfo        *MemoryInfoStat     `json:"memInfo"`
	CreateTime     int64               `json:"createTime"`
	Tgid           int32               `json:"tgid"`
}

func GetProcess(ctx context.Context) ([]*Process, error) {
	processlist, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}
	var plist []*Process
	for _, p := range processlist {
		name, _ := p.NameWithContext(ctx)
		status, _ := p.StatusWithContext(ctx)
		parent, _ := p.ParentWithContext(ctx)
		uids, _ := p.UidsWithContext(ctx)
		gids, _ := p.GidsWithContext(ctx)
		numT, _ := p.NumThreadsWithContext(ctx)
		memInfo, _ := p.MemoryInfoWithContext(ctx)
		ctime, _ := p.CreateTimeWithContext(ctx)
		tgid, _ := p.TgidWithContext(ctx)
		numctx, _ := p.NumCtxSwitchesWithContext(ctx)
		proc := &Process{
			Pid:            p.Pid,
			Name:           name,
			Status:         status,
			ParentPID:      parent.Pid,
			Uids:           uids,
			Gids:           gids,
			NumThreads:     numT,
			MemInfo:        (*MemoryInfoStat)(memInfo),
			CreateTime:     ctime,
			Tgid:           tgid,
			NumCtxSwitches: (*NumCtxSwitchesStat)(numctx),
		}
		plist = append(plist, proc)
	}
	return plist, nil
}
