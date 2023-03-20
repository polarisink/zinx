package ziface

// @author lqs
// @date 2023/3/16 10:06 PM

type IDataPack interface {
	GetHeadLen() uint32

	Pack(message IMessage) ([]byte, error)

	Unpack([]byte) (IMessage, error)
}
