package server

type IPacket interface {
	GetPacketId(IConnection) int
}

type ServerPacket interface {
	IPacket
	Write(IBuffer)
}

type ClientPacket interface {
	IPacket
	Read(IBuffer)
}

type IPacketHandler interface {
	HandleKeepAlive(ClientKeepAlive, IPlayer)
	HandleChatMessage(ClientChatMessage, IPlayer)
	HandleUseEntity(ClientUseEntity, IPlayer)
	HandlePlayerMovement(ClientPlayerMovement, IPlayer)
	HandlePlayerPosition(ClientPlayerPosition, IPlayer)
	HandlePlayerLook(ClientPlayerLook, IPlayer)
	HandlePlayerPositionLook(ClientPlayerPositionLook, IPlayer)
	HandlePlayerDigging(ClientPlayerDigging, IPlayer)
	HandlePlayerBlockPlacement(ClientPlayerBlockPlacement, IPlayer)
	HandleHeldItemChange(ClientHeldItemChange, IPlayer)
	HandleAnimation(ClientAnimation, IPlayer)
	HandleEntityAction(ClientEntityAction, IPlayer)
	HandleSteerVehicle(ClientSteerVehicle, IPlayer)
	HandleCloseWindow(ClientCloseWindow, IPlayer)
	HandleClickWindow(ClientClickWindow, IPlayer)
	HandleConfirmTransaction(ClientConfirmTransaction, IPlayer)
	HandleCreativeInventoryAction(ClientCreativeInventoryAction, IPlayer)
	HandleEnchantItem(ClientEnchantItem, IPlayer)
	HandleUpdateSign(ClientUpdateSign, IPlayer)
	HandlePlayerAbilities(ClientPlayerAbilities, IPlayer)
	HandleTabComplete(ClientTabComplete, IPlayer)
	HandleClientSettings(ClientSettings, IPlayer)
	HandleClientStatus(ClientStatus, IPlayer)
	HandlePluginMessage(ClientPluginMessage, IPlayer)
	HandleSpectate(ClientSpectate, IPlayer)
	HandleResourcePackStatus(ClientResourcePackStatus, IPlayer)
}

type BasicPacketHandler struct {
}

func (handler *BasicPacketHandler) HandleKeepAlive(packet ClientKeepAlive, player IPlayer) {
	conn := player.GetConnection()
	conn.GetMinecraftServer().GetKeepAliveSender().ResetCounter(conn)
}

func (handler *BasicPacketHandler) HandleChatMessage(packet ClientChatMessage, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleUseEntity(packet ClientUseEntity, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerMovement(packet ClientPlayerMovement, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerPosition(packet ClientPlayerPosition, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerLook(packet ClientPlayerLook, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerPositionLook(packet ClientPlayerPositionLook, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerDigging(packet ClientPlayerDigging, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerBlockPlacement(packet ClientPlayerBlockPlacement, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleHeldItemChange(packet ClientHeldItemChange, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleAnimation(packet ClientAnimation, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleEntityAction(packet ClientEntityAction, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleSteerVehicle(packet ClientSteerVehicle, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleCloseWindow(packet ClientCloseWindow, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleClickWindow(packet ClientClickWindow, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleConfirmTransaction(packet ClientConfirmTransaction, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleCreativeInventoryAction(packet ClientCreativeInventoryAction, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleEnchantItem(packet ClientEnchantItem, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleUpdateSign(packet ClientUpdateSign, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePlayerAbilities(packet ClientPlayerAbilities, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleTabComplete(packet ClientTabComplete, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleClientSettings(packet ClientSettings, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleClientStatus(packet ClientStatus, player IPlayer) {

}

func (handler *BasicPacketHandler) HandlePluginMessage(packet ClientPluginMessage, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleSpectate(packet ClientSpectate, player IPlayer) {

}

func (handler *BasicPacketHandler) HandleResourcePackStatus(packet ClientResourcePackStatus, player IPlayer) {

}

func CreatePacketHandler() IPacketHandler {
	return &BasicPacketHandler{}
}
