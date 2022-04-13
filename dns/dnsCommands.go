package dns

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	RecordTypeCNAME = "cname"
	RecordTypeA     = "A"
	RecordTypeNS    = "NS"
	RecordTypeMX    = "MX"
	RecordTypePTR   = "PTR"
	RecordTypeSRV   = "SRV"
	RecordTypeAAA   = "AAA"
)

type dnsCommandStruct struct {
	command string
	args    []string
}

func (dnscmd *dnsCommandStruct) recordType(name string) *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, name)
	return dnscmd
}

func (dnscmd *dnsCommandStruct) recordDelete() *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, "/recordDelete")
	return dnscmd
}
func (dnscmd *dnsCommandStruct) recordAdd() *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, "/recordAdd")
	return dnscmd
}
func (dnscmd *dnsCommandStruct) zoneName(name string) *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, name)
	return dnscmd
}

func (dnscmd *dnsCommandStruct) recordName(name string) *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, name)
	return dnscmd
}
func (dnscmd *dnsCommandStruct) recordData(data string) *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, data)
	return dnscmd
}

func (dnscmd *dnsCommandStruct) force() *dnsCommandStruct {
	dnscmd.args = append(dnscmd.args, "/f")
	return dnscmd
}

func (dnscmd *dnsCommandStruct) print() {
	fmt.Println(dnscmd.command, strings.Join(dnscmd.args, " "))
}

func (dnscmd *dnsCommandStruct) Run() error {
	fmt.Println(dnscmd.command, strings.Join(dnscmd.args, " "))
	x := exec.Command(dnscmd.command, dnscmd.args...)

	var buff, errBuf bytes.Buffer
	x.Stdout = &buff
	x.Stderr = &errBuf
	err := x.Run()
	return err

}

func dnsCommand() *dnsCommandStruct {

	dnscmd := &dnsCommandStruct{}
	dnscmd.command = "dnscmd"
	return dnscmd
}

//microsoft don't support edit. we must delete and add the record again
func (reqEditRecord *EditRecordStruct) execute(action string) error {

	if action != "addRecord" {
		err := dnsCommand().recordDelete().
			zoneName(reqEditRecord.ZoneName).recordName(reqEditRecord.RecordName).
			recordType(reqEditRecord.RecordType).force().Run()
		if err != nil {
			return err
		}
		if action == "deleteRecord" {
			return nil
		}
	}

	if action != "deleteRecord" {
		err := dnsCommand().recordAdd().
			zoneName(reqEditRecord.ZoneName).recordName(reqEditRecord.RecordName).
			recordType(reqEditRecord.RecordType).recordData(reqEditRecord.RecordData).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
