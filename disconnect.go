package gologix

import "log"

// to disconect we send two items - a null item and an unconnected data item for the unregister service
func (client *Client) Disconnect() error {
	if !client.Connected {
		return nil
	}
	client.Connected = false
	var err error

	items := make([]CIPItem, 2)
	items[0] = CIPItem{} // null item

	reg_msg := CIPMessage_UnRegister{
		Service:                CIPService_ForwardClose,
		CipPathSize:            0x02,
		ClassType:              CIPClass_8bit,
		Class:                  0x06,
		InstanceType:           CIPInstance_8bit,
		Instance:               0x01,
		Priority:               0x0A,
		TimeoutTicks:           0x0E,
		ConnectionSerialNumber: client.ConnectionSerialNumber,
		VendorID:               CIP_VendorID,
		OriginatorSerialNumber: CIP_SerialNumber,
		PathSize:               3,                                           // 16 bit words
		Path:                   [6]byte{0x01, 0x00, 0x20, 0x02, 0x24, 0x01}, // TODO: generate paths automatically
	}

	items[1] = NewItem(CIPItem_UnconnectedData, reg_msg)

	err = client.Send(CIPCommandSendRRData, MarshalItems(items)) // 0x65 is register session
	if err != nil {
		log.Panicf("Couldn't send unconnect req %v", err)
	}
	return nil

}

type CIPMessage_UnRegister struct {
	Service                CIPService
	CipPathSize            byte
	ClassType              CIPClassSize
	Class                  byte
	InstanceType           CIPInstanceSize
	Instance               byte
	Priority               byte
	TimeoutTicks           byte
	ConnectionSerialNumber uint16
	VendorID               uint16
	OriginatorSerialNumber uint32
	PathSize               uint16
	Path                   [6]byte
}
