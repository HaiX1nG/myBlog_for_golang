package utils

import (
	"fmt"
	"sync"
	"time"
)

const epoch = int64(1704038400) // 2023-12-31 23:59:59 UTC

type Snowflake struct {
	mu       sync.Mutex
	nodeID   int64
	sequence int64
}

func NewSnowflake(nodeID int64) *Snowflake {
	return &Snowflake{
		nodeID:   nodeID,
		sequence: 0,
	}
}

func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixNano()
	if timestamp < s.sequence {
		panic(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.sequence-timestamp))
	}

	if timestamp == s.sequence {
		s.sequence = (s.sequence + 1) % 4096
		if s.sequence == 0 {
			timestamp = s.waitNextMillis(timestamp)
		}
	} else {
		s.sequence = 0
	}

	return ((timestamp - epoch) << 22) | (s.nodeID << 12) | s.sequence
}

func (s *Snowflake) waitNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / 1e6
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixNano() / 1e6
	}
	return timestamp
}
