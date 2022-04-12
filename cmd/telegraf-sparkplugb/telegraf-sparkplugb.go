package main

import (
	"flag"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"telegraf-sparkplugb/pkg/sparkplugb"
	"telegraf-sparkplugb/pkg/spbcache"
	"time"
)


var spc *spbcache.SparkplugNodeCache

func main() {
	spc = spbcache.NewSparkplugNodeCache()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	configFile := flag.String("config", "telegraf-sparkplugb.yaml", "Path to config file (default is telegraf-sparkplugb.con")
	cmdLineLogLevel := flag.String("log-level", "warn", "trace, debug, info, warning, error, fatal, or panic" )
	flag.Parse()

	if cmdLineLogLevel != nil {
		level, err := log.ParseLevel(*cmdLineLogLevel)
		if err == nil {
			log.SetLevel(level)
		} else {
			log.SetLevel(log.WarnLevel)
			log.Warnf("Unknown log level '%s'", level)
		}
	}

	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}


	var config config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	options := mqtt.NewClientOptions()
	options.SetAutoReconnect(true)
	options.Username = config.Username
	options.Password = config.Password

	if config.ClientId != "" {
		options.ClientID = config.ClientId
	} else {
		options.ClientID = "telegraf"
	}

	options.SetAutoReconnect(true)
	options.SetOrderMatters(true)
	options.SetCleanSession(true)
	options.OnConnect = connectionEstablishedHandler
	options.OnConnectionLost = connectionLostHandler

	for _, server := range config.Brokers {
		options.AddBroker(server)
	}

	client := mqtt.NewClient(options)
	ct := client.Connect()

	success := ct.WaitTimeout(2 * time.Minute)
	if !success {
		log.Fatal("Timeout connecting to server")
	}

	log.Info("Connected to broker")

	<-sigs
}

func connectionEstablishedHandler(client mqtt.Client) {
	log.Info("Connected to broker")
	client.Subscribe("$sparkplug/#", 2, mqttIn)
	client.Subscribe("spBv1.0/#", 2, mqttIn)
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Info("Connection to broker lost")
}

func mqttIn(client mqtt.Client, message mqtt.Message) {
	sptopic := sparkplugb.ParseTopic(message.Topic())

	if sptopic != nil {
		var payload = &sparkplugb.Payload{}
		err := proto.Unmarshal(message.Payload(), payload)
		if err == nil {
			if sptopic.Namespace == "spBv1.0" {

			} else if sptopic.Namespace == "$sparkplug" {
			}
		} else {
			log.Errorf("Sparkplug message could nto be parsed: %s", err.Error());
		}
	} else {
		log.Errorf("Topic '%s' doesn't match", message.Topic())
	}
}


func handleNodeMessage( groupId string, messageType string, nodeId string, payload *sparkplugb.Payload) {
	if messageType == "NBIRTH" {
		spc.NBIRTH(groupId, nodeId)
	} else if messageType == "NDEATH" {
		spc.NDEATH(groupId, nodeId)
	} else {
		log.Warnf("Unknown message type %s for node: %s", messageType, nodeId)
	}
}

func handleDeviceMessage( groupId string, messageType string, nodeId string, deviceId string, payload *sparkplugb.Payload) {
	log.Infof("%s for device: %s", messageType, deviceId)


}
