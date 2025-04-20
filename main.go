package main

import (
	"main.go/handler"
)

func main() {
	// // Sample raw JSON (could also come from a file, DB, etc.)
	// rawJSON := `[
	// 	{
	// 		"data": {"temperature": 35.7},
	// 		"device_id": "sensor-1",
	// 		"timestamp": "2025-03-28T11:59:08+05:30"
	// 	},
	// 	{
	// 		"data": {"temperature": 28.2},
	// 		"device_id": "sensor-2",
	// 		"timestamp": "2025-03-28T11:59:08+05:30"
	// 	},
	// 	{
	// 		"data": {"temperature": 36.1},
	// 		"device_id": "sensor-2",
	// 		"timestamp": "2025-03-28T11:59:08+05:30"
	// 	}
	// ]`

	// var rawData []map[string]interface{}
	// json.Unmarshal([]byte(rawJSON), &rawData)

	// // Group by device_id
	// grouped := make(map[string][]map[string]interface{})
	// for _, obj := range rawData {
	// 	if deviceID, ok := obj["device_id"].(string); ok {
	// 		grouped[deviceID] = append(grouped[deviceID], obj)
	// 	}
	// }
	// handler.WsHandler()
	go handler.StartWebSocketListener()
	// Start Gin server
	// r := gin.Default()

	// r.GET("/sensor/:topic", handler.GetDeviceData)

	// r.Run(":8082") // Run on http://localhost:8080
	select {} // Block forever
}
