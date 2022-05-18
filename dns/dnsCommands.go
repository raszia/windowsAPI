package dns

import (
	"context"
	"net/http"
	"windows/command"
)

const (
	RecordTypeCNAME = "cname"
	RecordTypeA     = "A"
	RecordTypeNS    = "NS"
	RecordTypeMX    = "MX"
	RecordTypePTR   = "PTR"
	RecordTypeSRV   = "SRV"
	RecordTypeAAA   = "AAA"

	RecordArgDelCmd = "/recordDelete"
	RecordArgAddCmd = "/recordAdd"
	infoArg         = "/info"
	enumZonesArg    = "/enumZones"
	forceArg        = "/f"

	CMDdns = "dnscmd"
)

type recordStruct struct {
	ZoneName   string //"myco.local"
	RecordName string //"test"
	RecordType string `json:"recordType"` //"cname"
	RecordData string `json:"recordData"` //"test2.myco.local"
}

func getInfo(ctx context.Context) ([]byte, error) {
	res, err := command.CmdCommand(CMDdns).ArgAdd(infoArg).RunOutput(ctx)
	if err != nil {
		return nil, err
	}

	return res.GetStdOut().Bytes(), nil
}

func getZones(ctx context.Context) ([]byte, error) {
	res, err := command.CmdCommand(CMDdns).ArgAdd(enumZonesArg).RunOutput(ctx)
	if err != nil {
		return nil, err
	}

	return res.GetStdOut().Bytes(), nil
}

func recordAction(ctx context.Context, record *recordStruct, method string) error {

	switch method {
	case http.MethodDelete:
		if err := command.CmdCommand(CMDdns).
			ArgAdd(RecordArgDelCmd).
			ArgAdd(record.ZoneName).
			ArgAdd(record.RecordName).
			ArgAdd(record.RecordType).
			ArgAdd(forceArg).
			Run(); err != nil {
			return err
		}
	case http.MethodPost:
		if err := command.CmdCommand(CMDdns).
			ArgAdd(RecordArgAddCmd).
			ArgAdd(record.ZoneName).
			ArgAdd(record.RecordName).
			ArgAdd(record.RecordType).
			ArgAdd(record.RecordData).
			Run(); err != nil {
			return err
		}
	}

	return nil
}
