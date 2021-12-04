package util

import "github.com/sony/sonyflake"

var sf *sonyflake.Sonyflake

func GetMachineID() (uint16, error) {
	return uint16(1), nil
}

func init() {
	var st = sonyflake.Settings{
		MachineID: GetMachineID,
	}
	if sf = sonyflake.NewSonyflake(st); sf == nil {
		panic("sonyflake init failed")
	}
}

func NextID() (uint64, error) {
	return sf.NextID()
}
