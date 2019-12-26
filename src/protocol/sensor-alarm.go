package protocol

import (
	// "errors"
	"time"
	"encoding/binary"
)

const (
    MsgAlarm    byte = 0x20 //报警信息
)

func init() {
	
	if _, ok := msgHanlders[MsgAlarm]; ok == true {

	} else {
			msgHanlders[MsgAlarm] = sensorSettingHandle
	}
}


func alarmHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]

	// .setCrane_protocol_version(msg[7]);
	protocolVersion := msg[7]

	var craneTime time.Time 
	//设备偶尔会发送2080年的数据，为了不影响查询，遇到年份大于50的数据时，用当前系统时间替换
	if msg[8]>50 {
		// .setCrane_time(getTimestamp(new Date()));
		craneTime = time.Now()
	}else {
		//yyMMddHHmmss格式时间转换成yyyy-MM-dd HH:mm:ss形式
		//.setCrane_time(getTimeStampBy6ByteArrayTime( new byte[] {msg[8], msg[9], msg[10], msg[11], msg[12], msg[13]} ) );
		craneTime = time.Date(2000 + int(msg[8]), time.Month(msg[9]), int(msg[10]), int(msg[11]), int(msg[12]), int(msg[13]), 0, time.UTC)
	}
	
	// .setCrane_manufacturer_and_device_type( msg[14] );
	manufacturer_and_device_type :=  msg[14] 

	// .setCrane_alarm_reason( msg[15] );
	alarmReason := msg[15]

	//高度数据 = 设备属性的大臂高度-原来的高度数据
	// .setCrane_height( (short) ( boom_height-getShort(msg[16], msg[17])&0xFFFF ) );
	height := binary.BigEndian.Uint16(msg[16:]) 
	
	// .setCrane_range( getShort(msg[18], msg[19]) );
	craneRange := binary.BigEndian.Uint16(msg[18:]) 

	// .setCrane_rotation( getShort(msg[20], msg[21]) );
	rotation := binary.BigEndian.Uint16(msg[20:])

	// .setCrane_elevating_capacity( getShort(msg[22], msg[23]) );
	elevatingCapacity :=  binary.BigEndian.Uint16(msg[22:])

	// .setCrane_wind_speed( getShort(msg[24], msg[25]) );
	windSpeed := binary.BigEndian.Uint16(msg[24:])
		
	// .setCrane_tilt_angle( getShort(msg[26], msg[27]) );
	tileAngle := binary.BigEndian.Uint16(msg[26:])

	// .setCrane_weight_percentage( msg[28] );
	weightPercentage := msg[28]

	// .setCrane_torque_percentage( msg[29] );
	torquePercentage := msg[29]

	// .setCrane_wind_speed_percentage( msg[30] );
	windSpeedPercentage := msg[30]

	// .setCrane_tilt_percentage( msg[31] );
	tiltPercentage := msg[31]

	// .setCrane_braking_state( msg[32] );
	brakingState := msg[32]
	
	alram() 
	
	return nil
}

func alram()  {


}
