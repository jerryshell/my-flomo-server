package util

import (
	"hash/fnv"
	"os"
	"strconv"

	"github.com/sony/sonyflake/v2"
)

var sf *sonyflake.Sonyflake

func GetMachineID() (int, error) {
	// 优先从环境变量获取机器ID
	if machineIDStr := os.Getenv("MACHINE_ID"); machineIDStr != "" {
		if machineID, err := strconv.Atoi(machineIDStr); err == nil && machineID >= 0 && machineID <= 1023 {
			return machineID, nil
		}
	}

	// 如果环境变量未设置或无效，尝试获取主机名哈希作为机器ID
	hostname, err := os.Hostname()
	if err != nil {
		// 如果获取主机名失败，使用默认值0
		return 0, nil
	}

	// 使用FNV-1a哈希算法生成稳定的机器ID
	h := fnv.New32a()
	h.Write([]byte(hostname))
	machineID := int(h.Sum32() % 1024) // Sonyflake限制机器ID在0-1023范围内

	return machineID, nil
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
