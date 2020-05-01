package netflow

type V5Header struct {

	version 			uint //NetFlow export format version number
	count 				uint //Number of flows exported in this packet (1-30)
	sysUptime 			uint //Current time in milliseconds since the export device booted
	unixSecs 			uint //Current count of seconds since 0000 UTC 1970
	unixNsecs 			uint //Residual nanoseconds since 0000 UTC 1970
	flowSequence 		uint //Sequence counter of total flows seen
	engineType 			uint //Type of flow-switching engine
	engineId 			uint //Slot number of the flow-switching engine
	samplingInterval	uint //First two bits hold the sampling mode; remaining 14 bits hold value of sampling interval
}

type V5FlowRecord struct {
	srcaddr 	uint //Source IP address
	dstaddr 	uint //Destination IP address
	nexthop 	uint //IP address of next hop router
	input 		uint //SNMP index of input interface
	output 		uint //SNMP index of output interface
	dPkts 		uint //Packets in the flow
	dOctets 	uint //Total number of Layer 3 bytes in the packets of the flow
	First 		uint //SysUptime at start of flow
	Last 		uint //SysUptime at the time the last packet of the flow was received
	srcport 	uint //TCP/UDP source port number or equivalent
	dstport 	uint //TCP/UDP destination port number or equivalent
	pad1 		uint //Unused (zero) bytes
	tcpFlags 	uint //Cumulative OR of TCP flags
	prot 		uint //IP protocol type (for example, TCP = 6; UDP = 17)
	tos 		uint //IP type of service (ToS)
	srcAs 		uint //Autonomous system number of the source, either origin or peer
	dstAs 		uint //Autonomous system number of the destination, either origin or peer
	srcMask 	uint //Source address prefix mask bits
	dstMask 	uint //Destination address prefix mask bits
	pad2 		uint //Unused (zero) bytes
}



