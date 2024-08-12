package server

import (
	"encoding/json"
	"net/deltamc/server/component"
)

type ClientKeepAlive struct {
	ClientPacket
	KeepAliveId int32
}

func (packet *ClientKeepAlive) GetPacketId(conn IConnection) int {
	return ClientKeepAlivePacket
}

func (packet *ClientKeepAlive) Read(buffer IBuffer) {
	packet.KeepAliveId = buffer.ReadVarInt()
}

type ClientChatMessage struct {
	ClientPacket
	Text string
}

func (packet *ClientChatMessage) GetPacketId() int {
	return ClientChatMessagePacket
}

func (packet *ClientChatMessage) Read(buffer IBuffer) {
	packet.Text = buffer.ReadString()
}

type ClientUseEntity struct {
	ClientPacket
	Target    int32
	ClickType int32
	TargetX   float32
	TargetY   float32
	TargetZ   float32
}

func (packet *ClientUseEntity) GetPacketId() int {
	return ClientUseEntityPacket
}

func (packet *ClientUseEntity) Read(buffer IBuffer) {
	packet.Target = buffer.ReadVarInt()
	packet.ClickType = buffer.ReadVarInt()

	if packet.ClickType == 2 {
		packet.TargetX = buffer.ReadFloat()
		packet.TargetY = buffer.ReadFloat()
		packet.TargetZ = buffer.ReadFloat()
	}
}

type ClientPlayerMovement struct {
	ClientPacket
	OnGround bool
}

func (packet *ClientPlayerMovement) GetPacketId() int {
	return ClientPlayerPacket
}

func (packet *ClientPlayerMovement) Read(buffer IBuffer) {
	packet.OnGround = buffer.ReadBool()
}

type ClientPlayerPosition struct {
	ClientPacket
	PosX     double
	FeetY    double
	PosZ     double
	OnGround bool
}

func (packet *ClientPlayerPosition) GetPacketId() int {
	return ClientPlayerPositionPacket
}

func (packet *ClientPlayerPosition) Read(buffer IBuffer) {
	packet.PosX = buffer.ReadDouble()
	packet.FeetY = buffer.ReadDouble()
	packet.PosZ = buffer.ReadDouble()
	packet.OnGround = buffer.ReadBool()
}

type ClientPlayerLook struct {
	ClientPacket
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (packet *ClientPlayerLook) GetPacketId() int {
	return ClientPlayerLookPacket
}

func (packet *ClientPlayerLook) Read(buffer IBuffer) {
	packet.Yaw = buffer.ReadFloat()
	packet.Pitch = buffer.ReadFloat()
	packet.OnGround = buffer.ReadBool()
}

type ClientPlayerPositionLook struct {
	ClientPacket
	PosX     double
	FeetY    double
	PosZ     double
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (packet *ClientPlayerPositionLook) GetPacketId() int {
	return ClientPlayerPositionAndLookPacket
}

func (packet *ClientPlayerPositionLook) Read(buffer IBuffer) {
	packet.PosX = buffer.ReadDouble()
	packet.FeetY = buffer.ReadDouble()
	packet.PosZ = buffer.ReadDouble()
	packet.Yaw = buffer.ReadFloat()
	packet.Pitch = buffer.ReadFloat()
	packet.OnGround = buffer.ReadBool()
}

type ClientPlayerDigging struct {
	ClientPacket
	Status    DiggingStatus
	Location  *Vec3
	BlockFace BlockFace
}

func (packet *ClientPlayerDigging) GetPacketId() int {
	return ClientPlayerDiggingPacket
}

func (packet *ClientPlayerDigging) Read(buffer IBuffer) {
	packet.Status = buffer.ReadUInt8()
	packet.Location = buffer.ReadVec3()
	packet.BlockFace = buffer.ReadUInt8()
}

type ClientPlayerBlockPlacement struct {
	ClientPacket
	Location *Vec3
	Face     BlockFace
	//TODO: implement remaining fields? i dont wanna deal with this rn
	HeldItem   int16
	CursorPosX int8
	CursorPosY int8
	CursorPosZ int8
}

func (packet *ClientPlayerBlockPlacement) GetPacketId() int {
	return ClientPlayerBlockPlacementPacket
}

func (packet *ClientPlayerBlockPlacement) Read(buffer IBuffer) {
	packet.Location = buffer.ReadVec3()
	packet.Face = buffer.ReadUInt8()
	packet.HeldItem = buffer.ReadInt16()
	packet.CursorPosX = buffer.ReadInt8()
	packet.CursorPosY = buffer.ReadInt8()
	packet.CursorPosZ = buffer.ReadInt8()
}

type ClientHeldItemChange struct {
	ClientPacket
	Slot int16
}

func (packet *ClientHeldItemChange) GetPacketId() int {
	return ClientHeldItemChangePacket
}

func (packet *ClientHeldItemChange) Read(buffer IBuffer) {
	packet.Slot = buffer.ReadInt16()
}

type ClientAnimation struct {
	ClientPacket
}

func (packet *ClientAnimation) GetPacketId() int {
	return ClientAnimationPacket
}

func (packet *ClientAnimation) Read(buffer IBuffer) {

}

type ClientEntityAction struct {
	ClientPacket
	EntityId       int32
	ActionId       EntityAction
	HorseJumpBoost int32
}

func (packet *ClientEntityAction) GetPacketId() int {
	return ClientEntityActionPacket
}

func (packet *ClientEntityAction) Read(buffer IBuffer) {
	packet.EntityId = buffer.ReadVarInt()
	packet.ActionId = EntityAction(buffer.ReadVarInt())
	packet.HorseJumpBoost = buffer.ReadVarInt()
}

type ClientSteerVehicle struct {
	ClientPacket
	Sideways float32
	Forward  float32
	Flags    uint8
}

func (packet *ClientSteerVehicle) GetPacketId() int {
	return ClientSteerVehiclePacket
}

func (packet *ClientSteerVehicle) Read(buffer IBuffer) {
	packet.Sideways = buffer.ReadFloat()
	packet.Forward = buffer.ReadFloat()
	packet.Flags = buffer.ReadUInt8()
}

type ClientCloseWindow struct {
	ClientPacket
	WindowId uint8
}

func (packet *ClientCloseWindow) GetPacketId() int {
	return ClientCloseWindowPacket
}

func (packet *ClientCloseWindow) Read(buffer IBuffer) {
	packet.WindowId = buffer.ReadUInt8()
}

type ClientClickWindow struct {
	ClientPacket
	WindowId     uint8
	Slot         int16
	Button       int8
	ActionNumber int16
	Mode         int8
	ClickedItem  int16
}

func (packet *ClientClickWindow) GetPacketId() int {
	return ClientClickWindowPacket
}

func (packet *ClientClickWindow) Read(buffer IBuffer) {
	packet.WindowId = buffer.ReadUInt8()
	packet.Slot = buffer.ReadInt16()
	packet.Button = buffer.ReadInt8()
	packet.ActionNumber = buffer.ReadInt16()
	packet.Mode = buffer.ReadInt8()
	packet.ClickedItem = buffer.ReadInt16()
}

type ClientConfirmTransaction struct {
	ClientPacket
	WindowId     uint8
	ActionNumber int16
	Accepted     bool
}

func (packet *ClientConfirmTransaction) GetPacketId() int {
	return ClientConfirmTransactionPacket
}

func (packet *ClientConfirmTransaction) Read(buffer IBuffer) {
	packet.WindowId = buffer.ReadUInt8()
	packet.ActionNumber = buffer.ReadInt16()
	packet.Accepted = buffer.ReadBool()
}

type ClientCreativeInventoryAction struct {
	ClientPacket
	Slot int16
}

func (packet *ClientCreativeInventoryAction) GetPacketId() int {
	return ClientCreativeInventoryActionPacket
}

type ClientEnchantItem struct {
	ClientPacket
}

func (packet *ClientEnchantItem) GetPacketId() int {
	return ClientEnchantItemPacket
}

type ClientUpdateSign struct {
	ClientPacket
}

func (packet *ClientUpdateSign) GetPacketId() int {
	return ClientUpdateSignPacket
}

type ClientPlayerAbilities struct {
	ClientPacket
}

func (packet *ClientPlayerAbilities) GetPacketId() int {
	return ClientPlayerAbilitiesPacket
}

type ClientTabComplete struct {
	ClientPacket
}

func (packet *ClientTabComplete) GetPacketId() int {
	return ClientTabCompletePacket
}

type ClientSettings struct {
	ClientPacket
}

func (packet *ClientSettings) GetPacketId() int {
	return ClientSettingsPacket
}

type ClientStatus struct {
	ClientPacket
}

func (packet *ClientStatus) GetPacketId() int {
	return ClientStatusPacket
}

type ClientPluginMessage struct {
	ClientPacket
}

func (packet *ClientPluginMessage) GetPacketId() int {
	return ClientPluginMessagePacket
}

type ClientSpectate struct {
	ClientPacket
}

func (packet *ClientSpectate) GetPacketId() int {
	return ClientSpectatePacket
}

type ClientResourcePackStatus struct {
	ClientPacket
}

func (packet *ClientResourcePackStatus) GetPacketId() int {
	return ClientResourcePackPacket
}

// ══════════════════════════════════════════════════════════════════════
//
//                            SERVER PACKETS
//
// ══════════════════════════════════════════════════════════════════════

type ServerKeepAlive struct {
	KeepAliveId int32
}

func CreateServerKeepAlive(id int32) *ServerKeepAlive {
	return &ServerKeepAlive{KeepAliveId: id}
}

func (packet *ServerKeepAlive) GetPacketId(conn IConnection) int {
	return ClientKeepAlivePacket
}

func (packet *ServerKeepAlive) Write(buffer IBuffer) {
	buffer.WriteVarInt(packet.KeepAliveId)
}

type ServerJoinGame struct {
	EntityID         int
	Gamemode         uint8
	Dimension        byte
	Difficulty       uint8
	MaxPlayers       uint8
	LevelType        string
	ReducedDebugInfo bool
}

func CreateServerJoinGame(entityID int, gamemode uint8, dimension byte, difficulty uint8, maxPlayers uint8, levelType string, reducedDebugInfo bool) *ServerJoinGame {
	return &ServerJoinGame{
		EntityID:         entityID,
		Gamemode:         gamemode,
		Dimension:        dimension,
		Difficulty:       difficulty,
		MaxPlayers:       maxPlayers,
		LevelType:        levelType,
		ReducedDebugInfo: reducedDebugInfo,
	}
}

func (p *ServerJoinGame) GetPacketId(conn IConnection) int32 {
	return ServerJoinGamePacket
}

func (p *ServerJoinGame) Write(buf IBuffer) {
	buf.WriteInt(int32(p.EntityID))
	buf.WriteUInt8(p.Gamemode)
	buf.WriteByte(p.Dimension)
	buf.WriteUInt8(p.Difficulty)
	buf.WriteUInt8(p.MaxPlayers)
	buf.WriteString(p.LevelType)
	buf.WriteBool(p.ReducedDebugInfo)
}

type ServerChatMessage struct {
	JSONData *component.TextComponent
	Position byte
}

func CreateServerChatMessage(jsonData *component.TextComponent, position byte) *ServerChatMessage {
	return &ServerChatMessage{
		JSONData: jsonData,
		Position: position,
	}
}

func (p *ServerChatMessage) GetPacketId(conn IConnection) int32 {
	return ServerChatMessagePacket
}

func (p *ServerChatMessage) Write(buf IBuffer) {
	data, err := json.Marshal(p.JSONData)
	if err != nil {
		Error("Error serializing JSON Data (TextComponent) for ServerChatMessage packet: ", err)
		return
	}

	buf.WriteByteArray(data)
	buf.WriteByte(p.Position)
}
