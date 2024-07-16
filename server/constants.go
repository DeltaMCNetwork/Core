package server

const VERSION = "0.0.1"
const TICK_TIME = 50
const PROTOCOL_VERSION = 47
const PROTOCOL_NAME = "GoCraft " + VERSION
const BUFFER_SIZE = 2048
const USE_COMPRESSION = false
const USE_ENCRYPTION = false
const USE_PROXY = false

const (
	PacketModePing   = iota
	PacketModeStatus = iota
	PacketModeLogin  = iota
	PacketModePlay   = iota
)
