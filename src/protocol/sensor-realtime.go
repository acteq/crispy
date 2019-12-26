package protocol

import (
	// "errors"
	"time"
	"encoding/binary"
)

const (
	MsgRealtimeData     byte = 0x0c //工况数据上传（数据上传）
)

func init() {
	if _, ok := msgHanlders[MsgRealtimeData]; ok == true {
        
	} else {
		msgHanlders[MsgRealtimeData] = realtimeHandle
	}	
}

func realtimeHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]

	// .setCrane_protocol_version(msg[7]);//版本号
	protocolVersion := msg[7]

	//因为设备偶尔会发送2080年的数据，为了不影响查询，当遇到年份大于50的数据时，用当前时间替换
	if msg[8]>50 {
		// .setCrane_time(getTimestamp(new Date()));
		craneTime := time.Now()
	}else {
		//yyMMddHHmmss格式时间转换成yyyy-MM-dd HH:mm:ss形式
		//.setCrane_time(getTimeStampBy6ByteArrayTime( new byte[] {msg[8], msg[9], msg[10], msg[11], msg[12], msg[13]} ) );
		craneTime := time.Date(2000 + int(msg[8]), time.Month(msg[9]), int(msg[10]), int(msg[11]), int(msg[12]), int(msg[13]), 0, time.UTC)
	}
	// .setCrane_manufacturer_and_device_type( msg[14] );//
	manufacturer_and_device_type := msg[14]
	//高度数据 = 设备信息中的大臂高度-原来的高度数据
	//PS：boom_height的值在基本信息中
	// .setCrane_height( (short) ( boom_height-getShort(msg[15], msg[16])&0xFFFF) );} 
	crane_height := binary.BigEndian.Uint16(msg[15:])

	//幅度
	// .setCrane_range( getShort(msg[17], msg[18]) );
	craneRange := binary.BigEndian.Uint16(msg[17:])

	// .setCrane_rotation( getShort(msg[19], msg[20]) );
	//回转
	rotation := binary.BigEndian.Uint16(msg[19:])

	// .setCrane_elevating_capacity( getShort(msg[21], msg[22]) );
	//起重量数据
	elevatingCapacity := binary.BigEndian.Uint16(msg[21:])
	
	// .setCrane_wind_speed( getShort(msg[23], msg[24]) );
	//风速
	windSpeed := binary.BigEndian.Uint16(msg[23:])

	// .setCrane_tilt_angle( getShort(msg[25], msg[26]) );
	//倾角数据
	tiltAndle := binary.BigEndian.Uint16(msg[25:])

	// .setCrane_weight_percentage(msg[27]);
	//重量百分比
	weightPercentage := msg[27]

	// .setCrane_torque_percentage(msg[28]);
	//力矩百分比
	torquePercentage := msg[28]

	// .setCrane_wind_speed_percentage(msg[29]);//风速百分比
	windSpeedPercentage := msg[29]
	
	// .setCrane_tilt_percentage(msg[30]);
	//倾斜百分比
	tilePercentage := msg[30]
	
	// .setCrane_alarm_reason( msg[31] );
	//警报原因
	alarmReason := msg[31]
	
	// .setCrane_braking_state( msg[32] );//制动状态
	brakingState := msg[32]
	realtimeData()
	return nil
}
func realtimeData()  {
	



}
