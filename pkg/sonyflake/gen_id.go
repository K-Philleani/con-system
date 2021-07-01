package sonyflake

import (
	"fmt"
	sonyflake2 "github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake2.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需要传入当前的机器ID

func Init(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2021-07-01")
	settings := sonyflake2.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake2.NewSonyflake(settings)
	return
}

// GetId生成的id返回值

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
