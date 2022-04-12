package spbcache

import (
	log "github.com/sirupsen/logrus"
	"telegraf-sparkplugb/pkg/sparkplugb"
)

type SparkplugNode struct {

}

func NewSparkplugNode() *SparkplugNode {
	return &SparkplugNode{}
}



// DBIRTH A device is being born under this node
func (d *SparkplugNode) DBIRTH(payload sparkplugb.Payload) {


	for _, metric := range(payload.Metrics) {
		//alias := metric.Alias
		name := metric.Name
		//value := metric.Value

		log.Tracef("device add metric %s", name)
	}


}

// NDEATH a device is terminating under this lnode
func (d *SparkplugNode) NDEATH() {

}