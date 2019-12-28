package protocol

import (
	// "errors"
	"time"
	"encoding/binary"
	"log"
)

const (
	MsgRealtimeData     byte = 0x0c //工况数据上传（数据上传）
)

func init() {
	if _, ok := msgHanlders[MsgRealtimeData]; ok == true {
        log.Printf("can't add function realtimeHandle, the key %2X has exits\n", MsgRealtimeData)   
	} else {
		msgHanlders[MsgRealtimeData] = realtimeHandle
	}	
}

type SensorData struct {
	towerCraneId byte
	deviceId string
	protocolVersion byte
	createTime time.Time
	manufacturer_and_device_type byte
	craneHeight uint16
	craneRange uint16
	rotation uint16
	elevatingCapacity uint16
	windSpeed uint16
	tiltAngle uint16
	weightPercentage byte
	torquePercentage byte
	windSpeedPercentage byte	
	tilePercentage byte	
	alarmReason byte
	brakingState byte
}


func realtimeHandle(msg []byte) (resp []byte) {
	if len(msg) < 31 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}

	data := SensorData {} 
	//塔吊编号
	data.towerCraneId = msg[1];
	//设备编号 
	data.deviceId = BytetoH(msg[2:5])

	// .setCrane_protocol_version(msg[7]);//版本号
	data.protocolVersion = msg[5]

	//因为设备偶尔会发送2080年的数据，为了不影响查询，当遇到年份大于50的数据时，用当前时间替换
	if msg[8]>50 {
		// .setCrane_time(getTimestamp(new Date()));
		data.createTime = time.Now()
	}else {
		//yyMMddHHmmss格式时间转换成yyyy-MM-dd HH:mm:ss形式
		//.setCrane_time(getTimeStampBy6ByteArrayTime( new byte[] {msg[8], msg[9], msg[10], msg[11], msg[12], msg[13]} ) );
		data.createTime = BytesToTime(msg[6:12])
	}
	// .setCrane_manufacturer_and_device_type( msg[14] );//
	data.manufacturer_and_device_type = msg[12]
	//高度数据 = 设备信息中的大臂高度-原来的高度数据
	//PS：boom_height的值在基本信息中
	// .setCrane_height( (short) ( boom_height-getShort(msg[15], msg[16])&0xFFFF) );} 
	data.craneHeight = binary.BigEndian.Uint16(msg[13:])

	//幅度
	// .setCrane_range( getShort(msg[17], msg[18]) );
	data.craneRange = binary.BigEndian.Uint16(msg[15:])

	// .setCrane_rotation( getShort(msg[19], msg[20]) );
	//回转
	data.rotation = binary.BigEndian.Uint16(msg[17:])

	// .setCrane_elevating_capacity( getShort(msg[21], msg[22]) );
	//起重量数据
	data.elevatingCapacity = binary.BigEndian.Uint16(msg[19:])
	
	// .setCrane_wind_speed( getShort(msg[23], msg[24]) );
	//风速
	data.windSpeed = binary.BigEndian.Uint16(msg[21:])

	// .setCrane_tilt_angle( getShort(msg[25], msg[26]) );
	//倾角数据
	data.tiltAngle = binary.BigEndian.Uint16(msg[23:])

	// .setCrane_weight_percentage(msg[27]);
	//重量百分比
	data.weightPercentage = msg[25]

	// .setCrane_torque_percentage(msg[28]);
	//力矩百分比
	data.torquePercentage = msg[26]

	// .setCrane_wind_speed_percentage(msg[29]);//风速百分比
	data.windSpeedPercentage = msg[27]
	
	// .setCrane_tilt_percentage(msg[30]);
	//倾斜百分比
	data.tilePercentage = msg[28]
	
	// .setCrane_alarm_reason( msg[31] );
	//警报原因
	data.alarmReason = msg[29]
	
	// .setCrane_braking_state( msg[32] );//制动状态
	data.brakingState = msg[30]
	
	reportData(data)
	return nil
}

func reportData(data SensorData)  {
	log.Printf("data(%s):%+v\n", data.createTime.String(), data)
}
