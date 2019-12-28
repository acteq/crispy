package protocol

import (
	// "errors"
	"encoding/binary"
	"log"
	"time"
)

const (
	MsgCycleData    byte = 0x3d
)

func init() {
	if _, ok := msgHanlders[MsgCycleData]; ok == true {
		log.Printf("can't add function cycleHandle, the key %2X has exits\n", MsgCycleData)   
	} else {
		msgHanlders[MsgCycleData] = cycleHandle
	}
	
}


type Cycle struct {
	towerCraneId byte
	deviceId string
	protocolVersion byte
	manufacturer_and_device_type byte
	startTime time.Time
	endTime time.Time
	maxWeight uint16
	maxTorque uint16
	maxHeight uint16
	minHeight uint16
	maxRange uint16
	minRange uint16
	liftingPointAngle uint16
	liftingPointRange uint16
	liftingPointHeight uint16
	unloadingPointAngle uint16
	unloadingPointRange uint16
	unloadingPointHeight uint16
}


func cycleHandle(msg []byte) (resp []byte) {
	if len(msg) < 43 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}
	cycle := Cycle{}
	//塔吊编号
	cycle.towerCraneId = msg[1];
	//设备编号 
	cycle.deviceId = BytetoH(msg[2:5])

	// .setCrane_protocol_version(msg[7]);
	//版本号（1byte）
	cycle.protocolVersion = msg[5]
	// .setCrane_manufacturer_and_device_type( (short)(msg[8]&0xFF) );
	//厂家及设备类型（1byte）
	cycle.manufacturer_and_device_type = msg[6]

	// .setStart_time( getTimeStampBy6ByteArrayTime( new byte[] {msg[9], msg[10], msg[11], msg[12], msg[13], msg[14]} ) );
	//开始时间（6byte年月日时分秒)
	cycle.startTime = BytesToTime(msg[7:13])

	// .setEnd_time( getTimeStampBy6ByteArrayTime( new byte[] {msg[15], msg[16], msg[17], msg[18], msg[19], msg[20]} ) );
	//结束时间（6byte年月日时分秒)
	cycle.endTime = BytesToTime(msg[13:19]) 

	// .setCycle_maximum_weight( getShort(msg[21], msg[22]) );
	//本次循环最大吊重（2byte）
	cycle.maxWeight = binary.BigEndian.Uint16(msg[19:])

	// .setCycle_maximum_torque( getShort(msg[23], msg[24]) );
	//本次循环最大力矩（2byte）
	cycle.maxTorque = binary.BigEndian.Uint16(msg[21:]) 
	
	// .setMaximum_height( getShort(msg[25], msg[26]) );
	//最大高度（2byte）
	cycle.maxHeight = binary.BigEndian.Uint16(msg[23:]) 
	
	// .setMinimum_height( getShort(msg[27], msg[28]) );
	//最小高度（2byte）
	cycle.minHeight = binary.BigEndian.Uint16(msg[25:])

	// .setMaximum_range( getShort(msg[29], msg[30]) );
	//最大幅度（2byte）
	cycle.maxRange = binary.BigEndian.Uint16(msg[27:])

	// .setMinimum_range( getShort(msg[31], msg[32]) );
	//最小幅度（2byte）
	cycle.minRange = binary.BigEndian.Uint16(msg[29:])

	// .setLifting_point_angle( getShort(msg[33], msg[34]) );
	//起吊点角度（2byte）
	cycle.liftingPointAngle = binary.BigEndian.Uint16(msg[31:])

	// .setLifting_point_range( getShort(msg[35], msg[36]) );
	//起吊点幅度（2byte）
	cycle.liftingPointRange = binary.BigEndian.Uint16(msg[33:])

	// .setLifting_point_height( getShort(msg[37], msg[38]) );
	//起吊点高度(2byte)
	cycle.liftingPointHeight = binary.BigEndian.Uint16(msg[35:]) 

	// .setUnloading_point_angle( getShort(msg[39], msg[40]) );
	//卸吊点角度（2byte）
	cycle.unloadingPointAngle = binary.BigEndian.Uint16(msg[37:]) 

	// .setUnloading_point_range( getShort(msg[41], msg[42]) );
	//卸吊点幅度（2byte）
	cycle.unloadingPointRange = binary.BigEndian.Uint16(msg[39:])

	// .setUnloading_point_height( getShort(msg[43], msg[44]) );
	//卸吊点高度（2byte）
	cycle.unloadingPointHeight = binary.BigEndian.Uint16(msg[41:]) 

	reportCycleData(cycle)

	return nil
}

func reportCycleData(cycle Cycle)  {
	log.Printf("cycle:%+v\n", cycle)
}
