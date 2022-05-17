package dns

import (
	"context"
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
	CMDdns          = "dnscmd"
)

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

//microsoft don't support edit. we must delete and add the record again
func (reqEditRecord *EditRecordStruct) execute(action string) error {

	if action != "addRecord" {
		//delete a record
		if err := command.CmdCommand(CMDdns).
			ArgAdd(RecordArgDelCmd).
			ArgAdd(reqEditRecord.ZoneName).
			ArgAdd(reqEditRecord.RecordName).
			ArgAdd(reqEditRecord.RecordType).
			ArgAdd(command.CMDArgForce).
			Run(); err != nil {
			return err
		}
		if action == "deleteRecord" {
			return nil
		}
	}

	if action != "deleteRecord" {
		if err := command.CmdCommand(command.CMDdns).
			ArgAdd(RecordArgAddCmd).
			ArgAdd(reqEditRecord.ZoneName).
			ArgAdd(reqEditRecord.RecordName).
			ArgAdd(reqEditRecord.RecordType).
			ArgAdd(reqEditRecord.RecordData).
			Run(); err != nil {
			return err
		}
	}
	return nil
}
