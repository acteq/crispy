package protocol

import (
	// "errors"
	"time"
	"encoding/binary"
)

const (
	MsgCycleData    byte = 0x3d
)

func init() {
	if _, ok := msgHanlders[MsgCycleData]; ok == true {
		
	} else {
		msgHanlders[MsgCycleData] = cycleHandle
	}
	
}

func cycleHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]

	// .setCrane_protocol_version(msg[7]);
	//版本号（1byte）
	protocolVersion := msg[7]
	// .setCrane_manufacturer_and_device_type( (short)(msg[8]&0xFF) );
	//厂家及设备类型（1byte）
	manufacturer_and_device_type := msg[8]

	// .setStart_time( getTimeStampBy6ByteArrayTime( new byte[] {msg[9], msg[10], msg[11], msg[12], msg[13], msg[14]} ) );
	//开始时间（6byte年月日时分秒)
	startTime := time.Date(2000 + int(msg[9]), time.Month(msg[10]), int(msg[11]), int(msg[12]), int(msg[13]), int(msg[14]), 0,  time.UTC)

	// .setEnd_time( getTimeStampBy6ByteArrayTime( new byte[] {msg[15], msg[16], msg[17], msg[18], msg[19], msg[20]} ) );
	//结束时间（6byte年月日时分秒)
	endTime := time.Date(2000 + int(msg[15]), time.Month(msg[16]), int(msg[17]), int(msg[18]), int(msg[19]), int(msg[20]), 0, time.UTC) 

	// .setCycle_maximum_weight( getShort(msg[21], msg[22]) );
	//本次循环最大吊重（2byte）
	maximum_weight := binary.BigEndian.Uint16(msg[21:])

	// .setCycle_maximum_torque( getShort(msg[23], msg[24]) );
	//本次循环最大力矩（2byte）
	maxTorque := binary.BigEndian.Uint16(msg[23:]) 
	
	// .setMaximum_height( getShort(msg[25], msg[26]) );
	//最大高度（2byte）
	maxHeight := binary.BigEndian.Uint16(msg[25:]) 
	
	// .setMinimum_height( getShort(msg[27], msg[28]) );
	//最小高度（2byte）
	miniHeight := binary.BigEndian.Uint16(msg[27:])

	// .setMaximum_range( getShort(msg[29], msg[30]) );
	//最大幅度（2byte）
	maxRange := binary.BigEndian.Uint16(msg[29:])

	// .setMinimum_range( getShort(msg[31], msg[32]) );
	//最小幅度（2byte）
	minRange := binary.BigEndian.Uint16(msg[31:])

	// .setLifting_point_angle( getShort(msg[33], msg[34]) );
	//起吊点角度（2byte）
	liftingPointAngle := binary.BigEndian.Uint16(msg[33:])

	// .setLifting_point_range( getShort(msg[35], msg[36]) );
	//起吊点幅度（2byte）
	liftingPointRange := binary.BigEndian.Uint16(msg[35:])

	// .setLifting_point_height( getShort(msg[37], msg[38]) );
	//起吊点高度(2byte)
	liftingPointHeight := binary.BigEndian.Uint16(msg[37:]) 

	// .setUnloading_point_angle( getShort(msg[39], msg[40]) );
	//卸吊点角度（2byte）
	unloadingPointAngle := binary.BigEndian.Uint16(msg[39:]) 

	// .setUnloading_point_range( getShort(msg[41], msg[42]) );
	//卸吊点幅度（2byte）
	unloadingPointRange := binary.BigEndian.Uint16(msg[41:])

	// .setUnloading_point_height( getShort(msg[43], msg[44]) );
	//卸吊点高度（2byte）
	unloadingPointHeight := binary.BigEndian.Uint16(msg[43:]) 

	cycleData()

	return nil
}

func cycleData()  {
	
}
