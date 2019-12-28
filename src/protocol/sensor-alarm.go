package protocol

import (
	// "errors"
	"time"
	"encoding/binary"
	"log"
)

const (
    MsgAlarm    byte = 0x20 //报警信息
)

func init() {
	
	if _, ok := msgHanlders[MsgAlarm]; ok == true {
		log.Printf("can't add function alarmHandle, the key %2X has exits\n", MsgAlarm)
	} else {
		msgHanlders[MsgAlarm] = alarmHandle
	}
}

type Alarm struct {
	towerCraneId byte
	deviceId string
	protocolVersion byte
	created time.Time
	manufacturer_and_device_type byte
	alarmReason byte
	height uint16 
	craneRange uint16 
	rotation uint16
	elevatingCapacity uint16
	windSpeed uint16
	tileAngle uint16
	weightPercentage byte
	torquePercentage byte
	windSpeedPercentage byte
	tiltPercentage byte
	brakingState byte
}

func alarmHandle(msg []byte) (resp []byte) {
	if len(msg) < 31 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}

	alarm := Alarm {}
	//塔吊编号
	alarm.towerCraneId = msg[1];
	//设备编号 
	alarm.deviceId = BytetoH(msg[2:5])

	// .setCrane_protocol_version(msg[7]);
	alarm.protocolVersion = msg[5]

	//设备偶尔会发送2080年的数据，为了不影响查询，遇到年份大于50的数据时，用当前系统时间替换
	if msg[8]>50 {
		// .setCrane_time(getTimestamp(new Date()));
		alarm.created = time.Now()
	}else {
		//yyMMddHHmmss格式时间转换成yyyy-MM-dd HH:mm:ss形式
		//.setCrane_time(getTimeStampBy6ByteArrayTime( new byte[] {msg[8], msg[9], msg[10], msg[11], msg[12], msg[13]} ) );
		alarm.created = BytesToTime(msg[6:12])
	}
	
	// .setCrane_manufacturer_and_device_type( msg[14] );
	alarm.manufacturer_and_device_type =  msg[12] 

	// .setCrane_alarm_reason( msg[15] );
	alarm.alarmReason = msg[13]

	//高度数据 = 设备属性的大臂高度-原来的高度数据
	// .setCrane_height( (short) ( boom_height-getShort(msg[16], msg[17])&0xFFFF ) );
	alarm.height = binary.BigEndian.Uint16(msg[14:]) 
	
	// .setCrane_range( getShort(msg[18], msg[19]) );
	alarm.craneRange = binary.BigEndian.Uint16(msg[16:]) 

	// .setCrane_rotation( getShort(msg[20], msg[21]) );
	alarm.rotation = binary.BigEndian.Uint16(msg[18:])

	// .setCrane_elevating_capacity( getShort(msg[22], msg[23]) );
	alarm.elevatingCapacity =  binary.BigEndian.Uint16(msg[20:])

	// .setCrane_wind_speed( getShort(msg[24], msg[25]) );
	alarm.windSpeed = binary.BigEndian.Uint16(msg[22:])
		
	// .setCrane_tilt_angle( getShort(msg[26], msg[27]) );
	alarm.tileAngle = binary.BigEndian.Uint16(msg[24:])

	// .setCrane_weight_percentage( msg[28] );
	alarm.weightPercentage = msg[26]

	// .setCrane_torque_percentage( msg[29] );
	alarm.torquePercentage = msg[27]

	// .setCrane_wind_speed_percentage( msg[30] );
	alarm.windSpeedPercentage = msg[28]

	// .setCrane_tilt_percentage( msg[31] );
	alarm.tiltPercentage = msg[29]

	// .setCrane_braking_state( msg[32] );
	alarm.brakingState = msg[30]
	
	alert(alarm) 
	
	return nil
}

func alert(alarm Alarm)  {
	log.Printf("alarm:%+v\n", alarm)
}
