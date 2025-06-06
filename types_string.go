// Code generated by "stringer -output types_string.go -type=MapType,ProgramType,PinType"; DO NOT EDIT.

package ebpf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnspecifiedMap-0]
	_ = x[Hash-1]
	_ = x[Array-2]
	_ = x[ProgramArray-3]
	_ = x[PerfEventArray-4]
	_ = x[PerCPUHash-5]
	_ = x[PerCPUArray-6]
	_ = x[StackTrace-7]
	_ = x[CGroupArray-8]
	_ = x[LRUHash-9]
	_ = x[LRUCPUHash-10]
	_ = x[LPMTrie-11]
	_ = x[ArrayOfMaps-12]
	_ = x[HashOfMaps-13]
	_ = x[DevMap-14]
	_ = x[SockMap-15]
	_ = x[CPUMap-16]
	_ = x[XSKMap-17]
	_ = x[SockHash-18]
	_ = x[CGroupStorage-19]
	_ = x[ReusePortSockArray-20]
	_ = x[PerCPUCGroupStorage-21]
	_ = x[Queue-22]
	_ = x[Stack-23]
	_ = x[SkStorage-24]
	_ = x[DevMapHash-25]
	_ = x[StructOpsMap-26]
	_ = x[RingBuf-27]
	_ = x[InodeStorage-28]
	_ = x[TaskStorage-29]
	_ = x[BloomFilter-30]
	_ = x[UserRingbuf-31]
	_ = x[CgroupStorage-32]
	_ = x[Arena-33]
	_ = x[WindowsHash-268435457]
	_ = x[WindowsArray-268435458]
	_ = x[WindowsProgramArray-268435459]
	_ = x[WindowsPerCPUHash-268435460]
	_ = x[WindowsPerCPUArray-268435461]
	_ = x[WindowsHashOfMaps-268435462]
	_ = x[WindowsArrayOfMaps-268435463]
	_ = x[WindowsLRUHash-268435464]
	_ = x[WindowsLPMTrie-268435465]
	_ = x[WindowsQueue-268435466]
	_ = x[WindowsLRUCPUHash-268435467]
	_ = x[WindowsStack-268435468]
	_ = x[WindowsRingBuf-268435469]
}

const (
	_MapType_name_0 = "UnspecifiedMapHashArrayProgramArrayPerfEventArrayPerCPUHashPerCPUArrayStackTraceCGroupArrayLRUHashLRUCPUHashLPMTrieArrayOfMapsHashOfMapsDevMapSockMapCPUMapXSKMapSockHashCGroupStorageReusePortSockArrayPerCPUCGroupStorageQueueStackSkStorageDevMapHashStructOpsMapRingBufInodeStorageTaskStorageBloomFilterUserRingbufCgroupStorageArena"
	_MapType_name_1 = "WindowsHashWindowsArrayWindowsProgramArrayWindowsPerCPUHashWindowsPerCPUArrayWindowsHashOfMapsWindowsArrayOfMapsWindowsLRUHashWindowsLPMTrieWindowsQueueWindowsLRUCPUHashWindowsStackWindowsRingBuf"
)

var (
	_MapType_index_0 = [...]uint16{0, 14, 18, 23, 35, 49, 59, 70, 80, 91, 98, 108, 115, 126, 136, 142, 149, 155, 161, 169, 182, 200, 219, 224, 229, 238, 248, 260, 267, 279, 290, 301, 312, 325, 330}
	_MapType_index_1 = [...]uint8{0, 11, 23, 42, 59, 77, 94, 112, 126, 140, 152, 169, 181, 195}
)

func (i MapType) String() string {
	switch {
	case i <= 33:
		return _MapType_name_0[_MapType_index_0[i]:_MapType_index_0[i+1]]
	case 268435457 <= i && i <= 268435469:
		i -= 268435457
		return _MapType_name_1[_MapType_index_1[i]:_MapType_index_1[i+1]]
	default:
		return "MapType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnspecifiedProgram-0]
	_ = x[SocketFilter-1]
	_ = x[Kprobe-2]
	_ = x[SchedCLS-3]
	_ = x[SchedACT-4]
	_ = x[TracePoint-5]
	_ = x[XDP-6]
	_ = x[PerfEvent-7]
	_ = x[CGroupSKB-8]
	_ = x[CGroupSock-9]
	_ = x[LWTIn-10]
	_ = x[LWTOut-11]
	_ = x[LWTXmit-12]
	_ = x[SockOps-13]
	_ = x[SkSKB-14]
	_ = x[CGroupDevice-15]
	_ = x[SkMsg-16]
	_ = x[RawTracepoint-17]
	_ = x[CGroupSockAddr-18]
	_ = x[LWTSeg6Local-19]
	_ = x[LircMode2-20]
	_ = x[SkReuseport-21]
	_ = x[FlowDissector-22]
	_ = x[CGroupSysctl-23]
	_ = x[RawTracepointWritable-24]
	_ = x[CGroupSockopt-25]
	_ = x[Tracing-26]
	_ = x[StructOps-27]
	_ = x[Extension-28]
	_ = x[LSM-29]
	_ = x[SkLookup-30]
	_ = x[Syscall-31]
	_ = x[Netfilter-32]
	_ = x[CGroupSyscall-33]
	_ = x[WindowsXDP-268435457]
	_ = x[WindowsBind-268435458]
	_ = x[WindowsCGroupSockAddr-268435459]
	_ = x[WindowsSockOps-268435460]
	_ = x[WindowsXDPTest-268436454]
	_ = x[WindowsSample-268436455]
}

const (
	_ProgramType_name_0 = "UnspecifiedProgramSocketFilterKprobeSchedCLSSchedACTTracePointXDPPerfEventCGroupSKBCGroupSockLWTInLWTOutLWTXmitSockOpsSkSKBCGroupDeviceSkMsgRawTracepointCGroupSockAddrLWTSeg6LocalLircMode2SkReuseportFlowDissectorCGroupSysctlRawTracepointWritableCGroupSockoptTracingStructOpsExtensionLSMSkLookupSyscallNetfilterCGroupSyscall"
	_ProgramType_name_1 = "WindowsXDPWindowsBindWindowsCGroupSockAddrWindowsSockOps"
	_ProgramType_name_2 = "WindowsXDPTestWindowsSample"
)

var (
	_ProgramType_index_0 = [...]uint16{0, 18, 30, 36, 44, 52, 62, 65, 74, 83, 93, 98, 104, 111, 118, 123, 135, 140, 153, 167, 179, 188, 199, 212, 224, 245, 258, 265, 274, 283, 286, 294, 301, 310, 323}
	_ProgramType_index_1 = [...]uint8{0, 10, 21, 42, 56}
	_ProgramType_index_2 = [...]uint8{0, 14, 27}
)

func (i ProgramType) String() string {
	switch {
	case i <= 33:
		return _ProgramType_name_0[_ProgramType_index_0[i]:_ProgramType_index_0[i+1]]
	case 268435457 <= i && i <= 268435460:
		i -= 268435457
		return _ProgramType_name_1[_ProgramType_index_1[i]:_ProgramType_index_1[i+1]]
	case 268436454 <= i && i <= 268436455:
		i -= 268436454
		return _ProgramType_name_2[_ProgramType_index_2[i]:_ProgramType_index_2[i+1]]
	default:
		return "ProgramType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PinNone-0]
	_ = x[PinByName-1]
}

const _PinType_name = "PinNonePinByName"

var _PinType_index = [...]uint8{0, 7, 16}

func (i PinType) String() string {
	if i >= PinType(len(_PinType_index)-1) {
		return "PinType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PinType_name[_PinType_index[i]:_PinType_index[i+1]]
}
