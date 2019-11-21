package pb //pb2需要指针，为了统一
func Bool(v bool) bool {
	return v
}
func Int32(v int32) int32 {
	return v
}
func Int(v int) int32 {
	return int32(v)
}
func Int64(v int64) int64 {
	return v
}
func Float32(v float32) float32 {
	return v
}
func Float64(v float64) float64 {
	return v
}
func Uint32(v uint32) uint32 {
	return v
}
func Uint64(v uint64) uint64 {
	return v
}
func String(v string) string {
	return v
}

func (x ServerStatus) Enum() ServerStatus {
	return x
}

func (x SDKChannel) Enum() SDKChannel {
	return x
}

func (x OrderStatus) Enum() OrderStatus {
	return x
}

func (x GamerStatus) Enum() GamerStatus {
	return x
}

func (x ChatChannel) Enum() ChatChannel {
	return x
}

func (x MailType) Enum() MailType {
	return x
}

func (x MailState) Enum() MailState {
	return x
}

func (x PVPType) Enum() PVPType {
	return x
}

func (x ItemType) Enum() ItemType {
	return x
}

func (x TaskState) Enum() TaskState {
	return x
}

func (x GashaponType) Enum() GashaponType {
	return x
}

func (x RoomState) Enum() RoomState {
	return x
}

func (x RoomGamerState) Enum() RoomGamerState {
	return x
}
