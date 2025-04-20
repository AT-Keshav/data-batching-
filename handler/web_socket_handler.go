package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	weUrl         = "ws://localhost:8080/ws"
	DeviceDataMap = make(map[string][]map[string]interface{})
	dataLock      sync.Mutex
)

func StartWebSocketListener() {
	fmt.Println("Starting WebSocket listener...")
	conn, _, err := websocket.DefaultDialer.Dial(weUrl, nil)
	if err != nil {
		fmt.Println("Failed to connect to WebSocket server:", err)
		return
	}
	defer conn.Close()

	for {
		time.Sleep(1 * time.Second)
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message from WebSocket:", err)
			return
		}

		var rawData []map[string]interface{}
		if err := json.Unmarshal(msg, &rawData); err != nil {
			fmt.Println("Failed to unmarshal JSON:", err)
			continue
		}

		dataLock.Lock()
		for _, obj := range rawData {
			if deviceID, ok := obj["topic"].(string); ok {
				// fmt.Println("Device ID: ", deviceID)
				DeviceDataMap[deviceID] = append(DeviceDataMap[deviceID], obj)
				// fmt.Println("Length of the cash of topic ", deviceID, " is ", len(DeviceDataMap[deviceID]))
				// fmt.Println("Data: ", DeviceDataMap[deviceID])
				if len(DeviceDataMap[deviceID]) >= Batch {
					fmt.Println("sending to blockchain: ", deviceID, " and data: ", DeviceDataMap[deviceID])
					// fmt.Println(len(DeviceDataMap["mqtt_device_1_channel_2"]))
					data := make([]map[string]interface{}, len(DeviceDataMap[deviceID]))
					copy(data, DeviceDataMap[deviceID])
					go GetDeviceData(data)
					// dataLock.Lock()
					DeviceDataMap[deviceID] = make([]map[string]interface{}, 0)
					// dataLock.Unlock()
					fmt.Println("After sending to blockchain cash lenght ", len(DeviceDataMap[deviceID]), " for device ", deviceID)

				}
			}
			// dataMap, ok := obj["data"].(map[string]interface{})
			// if !ok {
			// 	fmt.Println("Failed to assert data field as string")
			// 	continue
			// }
			// if error_code, ok := dataMap["error_code"].(string); ok {
			// 	fmt.Println("Error Code: ", error_code)
			// 	DeviceDataMap[error_code] = append(DeviceDataMap[error_code], obj)
			// }
		}
		dataLock.Unlock()
	}
}

// func WebSocketHandler(c *gin.Context) {
// 	conn, _, err := websocket.DefaultDialer.Dial(WeUrl, nil)
// 	if err != nil {
// 		fmt.Println("Failed to connect to WebSocket server:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for {
// 		time.Sleep(1 * time.Second)
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Failed to read message from WebSocket:", err)
// 			return
// 		}
// 		var jsonData interface{}
// 		if err := json.Unmarshal(msg, &jsonData); err != nil {
// 			fmt.Println("Failed to unmarshal JSON:", err)
// 			continue
// 		}
// 		CacheLock.Lock()
// 		data, err := json.Marshal(jsonData)
// 		if err != nil {
// 			fmt.Println("Failed to marshal JSON:", err)
// 			continue
// 		}
// 		jsonString := string(data)
// 		Cache = append(Cache, jsonString)
// 		if len(Cache) >= BatchSize {
// 			go EndBlocker(c)
// 		}
// 		CacheLock.Unlock()
// 	}
// }

// func EndBlocker(c *gin.Context) {
// 	CacheLock.Lock()
// 	if len(Cache) == 0 {
// 		CacheLock.Unlock()
// 		return
// 	}
// 	batch := Cache
// 	Cache = []string{}
// 	CacheLock.Unlock()
// 	var rawData []map[string]interface{}
// 	json.Unmarshal([]byte(batch[0]), &rawData)
// 	grouped := make(map[string][]map[string]interface{})
// 	for _, obj := range rawData {
// 		if deviceID, ok := obj["device_id"].(string); ok {
// 			grouped[deviceID] = append(grouped[deviceID], obj)
// 		}
// 	}
// 	deviceId := c.Param("id")
// 	data := grouped[deviceId] // Get all packets for device
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": data,
// 	})
// }
