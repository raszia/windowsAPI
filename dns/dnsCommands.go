package dns

import (
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
)

//microsoft don't support edit. we must delete and add the record again
func (reqEditRecord *EditRecordStruct) execute(action string) error {

	if action != "addRecord" {
		//delete a record
		if err := command.CmdCommand(command.CMDdns).
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
