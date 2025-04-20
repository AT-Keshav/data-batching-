package handler

import (
	"fmt"
	"time"
)

const (
	Batch = 1000
)

func BatchTimeOut() {
	ticket := time.NewTicker(30 * time.Second)
	defer ticket.Stop()
	for {
		<-ticket.C
		EndBlocker()
	}
}

func GetDeviceData(data []map[string]interface{}) {
	dataLock.Lock()
	// id := deviceId
	// DeviceDataMap[id] = make([]map[string]interface{}, 0)
	// fmt.Println("✅ After sending to blockchain cash lenght ", len(DeviceDataMap[id]), " for device ", id)
	fmt.Println("Data Send to blockchain of length: ", len(data))
	dataLock.Unlock()

	// fmt.Println("Data Send to blockchain of length: ", len(data))

}
func EndBlocker() {
	fmt.Println("End Blocker")
}
