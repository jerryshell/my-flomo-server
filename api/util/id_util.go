package util

import (
	"strconv"

	"github.com/sony/sonyflake/v2"
)

var sf *sonyflake.Sonyflake

func GetMachineID() (int, error) {
	return 1, nil
}

func init() {
	var st = sonyflake.Settings{
		MachineID: GetMachineID,
	}
	var err error
	if sf, err = sonyflake.New(st); err != nil {
		panic("sonyflake init failed: " + err.Error())
	}
}

func NextID() (int64, error) {
	return sf.NextID()
}

func NextIDStr() (string, error) {
	logger := NewLogger("id_util")
	
	id, err := NextID()
	if err != nil {
		logger.Error("failed to generate next id", ErrorField(err))
		return "", err
	}
	
	logger.Debug("generated new id", StringField("id", strconv.FormatInt(id, 10)))
	return strconv.FormatInt(id, 10), nil
}
