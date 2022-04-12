package sparkplugb

import "regexp"

var spRegex = regexp.MustCompile("(?P<namespace>{\\$sparkplug|spBv1.0})/(?P<group_id>[^/]*)/(?P<msg_type>[^/]*)/(?P<node_id>[^/]*)(?:/(?P<device_id>.*))?")

type SparkplugTopic struct {
	Namespace string
	GroupId   string
	MsgType   string
	NodeId    string
	DeviceId  string
}

func ParseTopic(topic string) *SparkplugTopic {
	matches := spRegex.FindStringSubmatch(topic)

	if matches != nil {
		return &SparkplugTopic{
			Namespace: matches[1],
			GroupId:   matches[2],
			MsgType:   matches[3],
			NodeId:    matches[4],
			DeviceId:  matches[5],
		}
	} else {
		return nil
	}
}
