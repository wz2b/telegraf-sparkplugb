package spbcache

type SparkplugNodeCache struct {
	nodeCache map[string]*SparkplugNode
}

func NewSparkplugNodeCache() *SparkplugNodeCache {
	return &SparkplugNodeCache{
		nodeCache: make(map[string]*SparkplugNode),
	}
}

func (spc *SparkplugNodeCache) GetSparkplugNode(namespace string, nodeId string) *SparkplugNode {
	key := namespace + "/" + nodeId
	return spc.nodeCache[key]
}

// NBIRTH Process NBIRTH message.  This will either add or replace the node in the cache.
func (spc *SparkplugNodeCache) NBIRTH(groupId string, nodeId string) *SparkplugNode {
	key := groupId + "/" + nodeId
	node := NewSparkplugNode()
	spc.nodeCache[key] = node
	return node
}

// NDEATH Process NDEATH message.  This will remove the node from the cache.
func (spc *SparkplugNodeCache) NDEATH(groupId string, nodeId string) {
	key := groupId + "/" + nodeId
	delete(spc.nodeCache, key)
}