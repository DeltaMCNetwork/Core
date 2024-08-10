package server

const VERSION = "0.0.1"
const TICK_TIME = 50
const PROTOCOL_VERSION = 47
const PROTOCOL_NAME = "Moincraft " + VERSION
const BUFFER_SIZE = 2048
const USE_COMPRESSION = false
const USE_ENCRYPTION = true
const USE_PROXY = false
const VERIFY_TOKEN_LENGTH = 4

// no make it caps because every function from every part of the server is accessible
// thats why i dont use vars outside of structs

const (
	PacketModeHandshake = iota
	PacketModeStatus
	PacketModeLogin
	PacketModePlay
)

// Used in C07PlayerDigging
type DiggingStatus = byte

const (
	StartedDigging DiggingStatus = byte(iota)
	CancelledDigging
	FinishedDigging
	DropItemStack
	DropItem
	ShootArrowEat
)

type BlockFace = byte

const (
	NegY BlockFace = byte(iota)
	PosY
	NegZ
	PosZ
	NegX
	PosX
)

func ApplyBlockFace(vec *Vec3, face BlockFace) {
	switch face {
	case NegY:
		vec.y--
	case PosY:
		vec.y++
	case NegZ:
		vec.z--
	case PosZ:
		vec.z++
	case NegX:
		vec.x--
	case PosX:
		vec.x++
	}
}

// Used in 0X0B packet EntityAction

type EntityAction int32

const (
	StartSneaking EntityAction = iota
	StopSneaking
	LeaveBed
	StartSprinting
	StopSprinting
	JumpWithHorse
	OpenHorseInventory
)

// packet stuff for C07PlayerDigging, it has to be a byte
// fancy
// work on server packets

//ugly packetid constants

// Handshake client
const (
	ClientHandshakePacket      = iota
	LegacyServerListPingPacket = 254
)

// Play Server
const (
	ServerKeepAlivePacket = iota
	ServerJoinGamePacket
	ServerChatMessagePacket
	ServerTimeUpdatePacket
	ServerEntityEquipmentPacket
	ServerSpawnPositionPacket
	ServerUpdateHealthPacket
	ServerRespawnPacket
	ServerPlayerPositionAndLookPacket
	ServerHeldItemChangePacket
	ServerUseBedPacket
	ServerAnimationPacket
	ServerSpawnPlayerPacket
	ServerCollectItemPacket
	ServerSpawnMobPacket
	ServerSpawnPaintingPacket
	ServerSpawnExperienceOrb
	ServerEntityVelocityPacket
	ServerDestroyEntities
	ServerEntityPacket
	ServerEntityRelativeMovePacket
	ServerEntityLookPacket
	ServerEntityLookAndRelativeMovePacket
	ServerEntityTeleportPacket
	ServerEntityHeadLookPacket
	ServerEntityStatusPacket
	ServerAttachEntityPacket
	ServerEntityMetadataPacket
	ServerEntityEffectPacket
	ServerRemoveEntityEffectPacket
	ServerSetExperiencePacket
	ServerEntityPropertiesPacket
	ServerChunkDataPacket
	ServerMultiBlockChangePacket
	ServerBlockChangePacket
	ServerBlockActionPacket
	ServerBlockBreakAnimation
	ServerMapChunkBulkPacket
	ServerExplosionPacket
	ServerEffectPacket
	ServerSoundEffectPacket
	ServerParticlePacket
	ServerChangeGameStatePacket
	ServerGlobalEntityPacket
	ServerOpenWindowPacket
	ServerCloseWindowPacket
	ServerSetSlotPacket
	ServerWindowItemsPacket
	ServerWindowPropertyPacket
	ServerConfirmTransactionPacket
	ServerUpdateSignPacket
	ServerMapPacket
	ServerUpdateBlockEntityPacket
	ServerOpenSignEditorPacket
	ServerPlayListItemPacket
	ServerPlayerAbilitiesPacket
	ServerTabCompletePacket
	ServerScoreboardObjectivePacket
	ServerUpdateScorePacket
	ServerDisplayScoreboardPacket
	ServerTeamsPacket
	ServerPluginMessagePacket
	ServerDisconnectPacket
	ServerDifficultyPacket
	ServerCombatEventPacket
	ServerCameraPacket
	ServerWorldBorderPacket
	ServerTitlePacket
	ServerSetCompressionPacket
	ServerPlayListHeaderAndFooterPacket
	ServerResourcePackSendPacket
	ServerUpdateEntityNBTPacket
)

// Play Client
const (
	ClientKeepAlivePacket = iota
	ClientChatMessagePacket
	ClientUseEntityPacket
	ClientPlayerPacket
	ClientPlayerPositionPacket
	ClientPlayerLookPacket
	ClientPlayerPositionAndLookPacket
	ClientPlayerDiggingPacket
	ClientPlayerBlockPlacementPacket
	ClientHeldItemChangePacket
	ClientAnimationPacket
	ClientEntityActionPacket
	ClientSteerVehiclePacket
	ClientCloseWindowPacket
	ClientClickWindowPacket
	ClientConfirmTransactionPacket
	ClientCreativeInventoryActionPacket
	ClientEnchantItemPacket
	ClientUpdateSignPacket
	ClientPlayerAbilitiesPacket
	ClientTabCompletePacket
	ClientSettingsPacket
	ClientStatusPacket
	ClientPluginMessagePacket
	ClientSpectatePacket
	ClientResourcePackPacket
)

// Status Server
const (
	ServerResponsePacket = iota
	ServerPongPacket
)

// Status Client
const (
	ClientRequestPacket = iota
	ClientPingPacket
)

// Login Server
const (
	ServerLoginDisconnectPacket = iota
	ServerEncryptionRequestPacket
	ServerLoginSuccessPacket
	ServerLoginSetCompressionPacket
)

// Login Client
const (
	ClientLoginStartPacket = iota
	ClientEncryptionResponsePacket
)
