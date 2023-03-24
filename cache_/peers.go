package cache_

import pb "minicache/cache_/cachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
type PeerGetter interface {
	Get(in *pb.Request, key *pb.Response) error
}
