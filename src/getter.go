package cache

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"log"
	"os"
	"time"
)

import (
	"consistenthash"
	pb "proto"
)

type Getter struct {
	conn  *nats.Conn
	hash  *consistenthash.Map
	fn    func(string) (interface{}, error)
	group *Group
	cache *cache
}

func New(addr string, replicas int, peers []string, fn func(string) (interface{}, error), timeout int64) *Getter {
	return newGetter(addr, peers, replicas, fn, newCache(timeout))
}

func newGetter(addr string, peers []string, replicas int, fn func(string) (interface{}, error), cache *cache) *Getter {
	conn, err := nats.Connect(addr)
	if err != nil {
		log.Fatal("connect nats addr", err)
		os.Exit(-1)
	}

	getter := &Getter{conn, consistenthash.New(replicas, nil), fn, &Group{}, cache}
	getter.hash.Add(peers...)
	getter.Sub(addr)
	return getter
}

func (getter *Getter) Get(key string) (interface{}, bool) {
	if v, ok := getter.cache.get(key); ok {
		return v, true
	}

	subject := getter.getSubject(key)
	if subject != "" {
		resp := getter.request(subject, &pb.Request{key})
		if resp != nil {
			return resp.Value, true
		}
	}
	return nil, false
}

func (getter *Getter) getSubject(key string) string {
	return getter.hash.Get(key)
}

func (getter *Getter) request(subject string, req *pb.Request) *pb.Response {
	val, _ := getter.group.Do(req.Key, func() (interface{}, error) {
		if v, ok := getter.cache.get(req.Key); ok {
			log.Println("Request get cache")
			return v, nil
		}

		b, _ := proto.Marshal(req)
		msg, err := getter.conn.Request(subject, b, 10*time.Millisecond)
		if err == nil {
			// TODO
			fmt.Println("Request get from other", "subject:", subject, "key:", req.Key, "value:", msg.Data)
			getter.cache.hotCache.Add(req.Key, msg.Data)
			//return msg.Data, nil
		}
		return msg.Data, err
	})

	log.Println("vvv:", val)
	resp := pb.Response{}
	proto.Unmarshal(val.([]byte), &resp)
	return &resp
}

func (getter *Getter) getLocal(key string) (interface{}, error) {
	val, err := getter.fn(key)
	log.Println("getLocal", "key:", key, "value:", val)
	return val, err
}

// async
func (getter *Getter) Sub(subject string) {
	getter.conn.Subscribe(subject, func(m *nats.Msg) {
		// TODO
		req := pb.Request{}
		proto.Unmarshal(m.Data, &req)

		var resp pb.Response
		val, err := getter.getLocal(req.Key)
		if err == nil {
			resp.Value = val.([]byte)
			resp.Ok = true
			getter.cache.mainCache.Add(req.Key, val)
		}

		b, _ := proto.Marshal(&resp)
		getter.conn.Publish(m.Reply, b)
	})
}
