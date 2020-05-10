package main

import (
	"adhoc.com/mqtt-server/db"
	"adhoc.com/mqtt-server/model"
	"encoding/json"
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type SensorDaten = model.SensorDaten



var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	var ss SensorDaten
	err := json.Unmarshal(msg.Payload(), &ss)
	if err != nil {
		panic(err);
	}
	fmt.Printf("MSG: %s\n", ss.SensorID)
	for _, data :=range ss.Data{
		fmt.Printf("MSG: %d\n", data)
	}

	res, err := db.Insert("sensor", ss.ToBSON())
	id := res.InsertedID
	fmt.Println("Inserted ID: ",id)

	db.GetDataByName("sensor", ss.SensorID)

}

func main() {

	mqttBroker := flag.String("broker", "tcp://0.0.0.0:1883", "your mqtt broker link. Example: tcp://0.0.0.0:1883")
	subTopic := flag.String("topic", "sensors", "your subcription topic. Default: sensors")

	flag.Parse()

	chanel1 := make(chan os.Signal, 1)

	opts := MQTT.NewClientOptions().AddBroker(*mqttBroker)
	opts.SetClientID("123456")
	opts.SetDefaultPublishHandler(messageHandler)
	topic := *subTopic

	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to Mqtt broker\n")
	}
	<-chanel1
}
