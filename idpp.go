package idpp

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"net"
	"os"
	"time"
)

var (
	timeEpoch = uint64(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC).
			UnixNano())

	workerID uint16
	sequence uint32
)

// ID describes a time based id
type ID interface {
	String() string
	Hex() string
	Time() time.Time
	WorkerID() uint16
	Sequence() uint16
	Validate() error
}

func init() {
	hash := sha256.New()

	hostname, _ := os.Hostname()
	hash.Write([]byte(hostname))

	pid := make([]byte, 4)
	binary.BigEndian.PutUint32(pid, uint32(os.Getpid()))
	hash.Write(pid)

	netInterfaces, _ := net.Interfaces()
	for _, netInterface := range netInterfaces {
		if addrs, err := netInterface.Addrs(); err == nil {
			for index, addr := range addrs {
				indexByte := make([]byte, 4)
				binary.BigEndian.PutUint32(indexByte, uint32(index))

				hash.Write(indexByte)
				hash.Write([]byte(netInterface.Name))
				hash.Write([]byte(addr.Network() + addr.String()))
				hash.Write(netInterface.HardwareAddr)
			}
		}
	}

	workerID = binary.BigEndian.Uint16(hash.Sum(nil))

	random := make([]byte, 4)
	rand.Read(random)
	sequence = binary.BigEndian.Uint32(random)
}
