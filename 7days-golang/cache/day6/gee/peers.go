package gee

// PeerPicker 根据传入的key选择相应节点的PeerGetter
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 从对应group中查找缓存值
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
