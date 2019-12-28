package protocol

import (
	// "errors"
	"log"
	"encoding/binary"
)

const MsgCraneInfo  byte = 0x31 //基本属性信息  0x31 = 49

func init() {
	_, ok := msgHanlders[MsgCraneInfo]
	if ok == true {
        log.Printf("can't add function craneInfoHandle, the key %2X has exits\n", MsgCraneInfo)
    } else {
        msgHanlders[MsgCraneInfo] = craneInfoHandle
    }
}

type TowerCrane struct {
	id byte
	deviceId string
	coordinate_x uint16
	coordinate_y uint16
	boomLength uint16
	balanceArmLength uint16
	towerHatHeight uint16
	boomHeight uint16
	maxLiftingWeight uint16
	maxTorque uint16
	propertyStatus byte
	craneType string
	manufacturer string
	hookWeight uint16
	articulatedLength uint16
	sensorInstallationStatus byte
}


func craneInfoHandle(msg []byte) (resp []byte) {
	if len(msg) < 87 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}
	crane := TowerCrane {}
	//塔吊编号
	crane.id = msg[1]
	//设备编号 
	crane.deviceId = BytetoH(msg[2:5])

	// .setCoordinate_x( new BigDecimal(getShort(data[7], data[8])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//坐标x
	crane.coordinate_x = binary.BigEndian.Uint16(msg[5:])
	
	// .setCoordinate_y( new BigDecimal(getShort(data[9], data[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//坐标y
	crane.coordinate_y = binary.BigEndian.Uint16(msg[7:])
	
	// .setBoom_length( new BigDecimal(getShort(data[11], data[12])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//起重臂长(前臂长)
	crane.boomLength = binary.BigEndian.Uint16(msg[9:])
	
	// .setBalance_arm_length( new BigDecimal(getShort(data[13], data[14])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//平衡臂长(后臂长)
	crane.balanceArmLength = binary.BigEndian.Uint16(msg[11:])
	
	// .setTower_hat_height( new BigDecimal(getShort(data[15], data[16])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//塔帽高
	crane.towerHatHeight = binary.BigEndian.Uint16(msg[13:])

	// .setBoom_height( new BigDecimal(getShort(data[17], data[18])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//起重臂高(塔臂高)
	crane.boomHeight = binary.BigEndian.Uint16(msg[15:])

	// .setMaximum_lifting_weight( new BigDecimal(getShort(data[19], data[20])&0xFFFF).setScale(4, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b1000) );
	//最大吊重 2byte
	crane.maxLiftingWeight = binary.BigEndian.Uint16(msg[17:])
	
	// .setMaximum_torque( new BigDecimal(getShort(data[21], data[22])&0xFFFF).setScale(2, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	//最大力矩 2byte
	crane.maxTorque = binary.BigEndian.Uint16(msg[19:])

	//.setMaximum_torque( new BigDecimal(getShort(data[21], data[22])&0xFFFF) );
	//最大力矩 2byte 因客户需求，暂时先将最大力矩除以10
	//maxRorque := binary.BigEndian.Uint16(data[21], data[22])
	
	// .setProperty_status( (short)(data[23]&0xFF) );
	//产权状态 1byte
	crane.propertyStatus = msg[21]
	
	// .setCrane_type( getStringTrim(data, 24, 54, "GB2312") );//塔机型号 30byte
	// craneType, _ :=  Utf82Gb2312(msg[22:52])
	// crane.craneType = string(craneType)

	// .setManufacturer( getStringTrim(data, 54, 84, "GB2312") );//生产厂商 30byte
	// manufacturer, _ := Utf82Gb2312(msg[52:82])
	// crane.manufacturer = string(manufacturer)
	
	//.setHook_weight( new BigDecimal(getShort(data[84], data[85])&0xFFFF) );
	//吊钩重量 2byte 暂时不解析
	crane.hookWeight = binary.BigEndian.Uint16(msg[82:])

	//.setArticulated_length( new BigDecimal(getShort(data[86], data[87])&0xFFFF) );
	//铰接长度 2byte 暂时不解析
	crane.articulatedLength = binary.BigEndian.Uint16(msg[84:])

	// .setSensor_installation_status( data[88] );
	//传感器安装状态 1byte
	crane.sensorInstallationStatus = msg[86]

	craneInfo(crane)

	return nil
}

func craneInfo(crane TowerCrane) {
	log.Printf("crane:%+v\n", crane)
} 
