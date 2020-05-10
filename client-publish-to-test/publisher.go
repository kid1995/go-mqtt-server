package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type SensorDaten struct {
	SensorID string       `json:"id"`
	Data     []	int `json:"data"`
}


func main() {

	ssData := new(SensorDaten)


	opts := MQTT.NewClientOptions().AddBroker("tcp://0.0.0.0:1883")
	opts.SetClientID("1156")

	topic := "sensors"
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')

	ssData.SensorID = msg
	ssData.Data = []int{23, 44, 22, 11}

	ssPacket, _ :=json.Marshal(ssData)
	fmt.Println(string(ssPacket))


	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Sample Publisher Started")
	fmt.Println("---- doing publish ----")
	token := client.Publish(topic, 0, false, string(ssPacket))
	token.Wait()

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")
}
