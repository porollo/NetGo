package netflow

import "github.com/jinzhu/gorm"

type V5Header struct {
	gorm.Model
	version          uint16 //NetFlow export format version number
	count            uint16 //Number of flows exported in this packet (1-30)
	sysUptime        uint32 //Current time in milliseconds since the export device booted
	unixSecs         uint32 //Current count of seconds since 0000 UTC 1970
	unixNsecs        uint32 //Residual nanoseconds since 0000 UTC 1970
	flowSequence     uint32 //Sequence counter of total flows seen
	engineType       uint8  //Type of flow-switching engine
	engineId         uint8  //Slot number of the flow-switching engine
	samplingInterval uint16 //First two bits hold the sampling mode; remaining 14 bits hold value of sampling interval
}

type V5FlowRecord struct {
	gorm.Model
	srcaddr  uint32 //Source IP address
	dstaddr  uint32 //Destination IP address
	nexthop  uint32 //IP address of next hop router
	input    uint8  //SNMP index of input interface
	output   uint8  //SNMP index of output interface
	dPkts    uint32 //Packets in the flow
	dOctets  uint32 //Total number of Layer 3 bytes in the packets of the flow
	First    uint32 //SysUptime at start of flow
	Last     uint32 //SysUptime at the time the last packet of the flow was received
	srcport  uint8  //TCP/UDP source port number or equivalent
	dstport  uint8  //TCP/UDP destination port number or equivalent
	pad1     uint8  //Unused (zero) bytes
	tcpFlags uint8  //Cumulative OR of TCP flags
	prot     uint8  //IP protocol type (for example, TCP = 6; UDP = 17)
	tos      uint8  //IP type of service (ToS)
	srcAs    uint16 //Autonomous system number of the source, either origin or peer
	dstAs    uint16 //Autonomous system number of the destination, either origin or peer
	srcMask  uint8  //Source address prefix mask bits
	dstMask  uint8  //Destination address prefix mask bits
	pad2     uint16 //Unused (zero) bytes
}
