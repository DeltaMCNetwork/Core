package server

// dont do slots like this, inventory will be an array of Item type. Check #items.go

type Slot struct {
	ID        int16
	ItemCount byte
	Damage    int16
	//TODO: NBT
}
