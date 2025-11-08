package util

import (
	"strconv"

	"github.com/sony/sonyflake"
)

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

func NextIDStr() (string, error) {
	logger := NewLogger("id_util")
	
	id, err := NextID()
	if err != nil {
		logger.Error("failed to generate next id", ErrorField(err))
		return "", err
	}
	
	logger.Debug("generated new id", StringField("id", strconv.FormatUint(id, 10)))
	return strconv.FormatUint(id, 10), nil
}
