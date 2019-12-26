package protocol

import (
	// "errors"
	"encoding/binary"
)

const (
	MsgSensorSetting     byte = 0x33 //传感器标定信息 0x33 = 51
    MsgAlramSetting    byte = 0x35 //报警设置信息
	MsgThresholdSetting    byte = 0x37 //限位设置信息 0x37
)

func init() {
	if _, ok := msgHanlders[MsgSensorSetting]; ok == true {
        
    } else {
        msgHanlders[MsgSensorSetting] = sensorSettingHandle
	}
	
	if _, ok := msgHanlders[MsgAlramSetting]; ok == true {
        
		} else {
			msgHanlders[MsgAlramSetting] = alarmSettingHandle
	}
	if _, ok := msgHanlders[MsgThresholdSetting]; ok == true {
	
	} else {
			msgHanlders[MsgThresholdSetting] = thresholdSettingHandle
	}	
}

func sensorSettingHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]
	
	// .setProximal_amplitude_demarcate_AD( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF) );
	proximal_amplitude_demarcate_AD := binary.BigEndian.Uint16(msg[7:])

	// .setProximal_amplitude_demarcate_actual( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	proximal_amplitude_demarcate_actual := binary.BigEndian.Uint16(msg[9:])

	// .setRemote_service_demarcate_AD( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF) );
	remote_service_demarcate_AD := binary.BigEndian.Uint16(msg[11:])

	// .setRemote_service_demarcate_actual( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	remote_service_demarcate_actual := binary.BigEndian.Uint16(msg[13:])

	// .setProximal_height_demarcate_AD( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF) );
	proximal_height_demarcate_AD := binary.BigEndian.Uint16(msg[15:])

	// .setProximal_height_demarcate_actual( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	proximal_height_demarcate_actual := binary.BigEndian.Uint16(msg[17:])

	// .setRemote_height_demarcate_AD( new BigDecimal(getShort(msg[19], msg[20])&0xFFFF) );
	remoteHeightDemarcateAD := binary.BigEndian.Uint16(msg[19:])

	// .setRemote_height_demarcate_actual( new BigDecimal(getShort(msg[21], msg[22])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	remote_height_demarcate_actual := binary.BigEndian.Uint16(msg[21:])

	// .setNoLoad_weight_AD( new BigDecimal(getShort(msg[23], msg[24])&0xFFFF) );
	noLoad_weight_AD := binary.BigEndian.Uint16(msg[23:])

	// .setNoLoad_weight_actual( new BigDecimal(getShort(msg[25], msg[26])&0xFFFF) );
	noLoad_weight_actual := binary.BigEndian.Uint16(msg[25:])

	// .setLoad_weight_AD( new BigDecimal(getShort(msg[27], msg[28])&0xFFFF) );
	load_weight_AD := binary.BigEndian.Uint16(msg[27:])

	// 原始代码有错，把实际重量放到空载重量中去了
	// .setNoLoad_weight_actual( new BigDecimal(getShort(msg[29], msg[30])&0xFFFF) );
	load_weight_actual := binary.BigEndian.Uint16(msg[29:])
	
	// .setRotation_starting_point_AD( new BigDecimal(getShort(msg[31], msg[32])&0xFFFF) );
	rotation_starting_point_AD := binary.BigEndian.Uint16(msg[31:])

	// .setRotation_starting_point_actual( new BigDecimal(getShort(msg[33], msg[34])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	rotation_starting_point_actual := binary.BigEndian.Uint16(msg[33:])

	// .setRotation_endinig_point_AD( new BigDecimal(getShort(msg[35], msg[36])&0xFFFF) );
	rotationEndinigPointAD := binary.BigEndian.Uint16(msg[35:])
	
	// .setRotation_endinig_point_actual( new BigDecimal(getShort(msg[37], msg[38])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	rotation_endinig_point_actual := binary.BigEndian.Uint16(msg[37:])
	
	// .setWind_speed_calibration( new BigDecimal(getShort(msg[39], msg[40])&0xFFFF) );
	//风速校准系数（2byte）
	windSpeedCalibration := binary.BigEndian.Uint16(msg[39:])
	
	// .setTilt_the_calibration( new BigDecimal(getShort(msg[47], msg[48])&0xFFFF) );
	//倾斜校准系数（2byte）
	tilt_the_calibration := binary.BigEndian.Uint16(msg[47:])
	
	sensorSetting()

	return nil
}

func sensorSetting()  {
	
}



func alarmSettingHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]


	// .setAlarm_horizontal_distance( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	horizontal_distance := binary.BigEndian.Uint16(msg[7:])
	
	// .setAlarm_perpendicular_distance( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	perpendicular_distance := binary.BigEndian.Uint16(msg[9:])
	
	// .setAlarm_weight_distance( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF) );
	weight_distance := binary.BigEndian.Uint16(msg[11:])

	// .setAlarm_torque( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF) );
	torque := binary.BigEndian.Uint16(msg[13:])

	// .setAlarm_wind_speed( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	windSpeed := binary.BigEndian.Uint16(msg[15:])

	// .setAlarm_tilt( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF) );
	alarmTilt := binary.BigEndian.Uint16(msg[17:])
	
	// .setWarning_horizontal_distance( new BigDecimal(getShort(msg[19], msg[20])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	warningHorizontalDistance := binary.BigEndian.Uint16(msg[19:])
	
	// .setWarning_perpendicular_distance( new BigDecimal(getShort(msg[21], msg[22])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	warningPerpendicularDistance := binary.BigEndian.Uint16(msg[21:])

	// .setWarning_weight_percentage( new BigDecimal(getShort(msg[23], msg[24])&0xFFFF) );
	warningWeightPercentage := binary.BigEndian.Uint16(msg[23:])
	
	// .setWarning_torque( new BigDecimal(getShort(msg[25], msg[26])&0xFFFF) );
	warningTorque := binary.BigEndian.Uint16(msg[25:])

	// .setWarning_wind_speed( new BigDecimal(getShort(msg[27], msg[28])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	warningWindSpeed := binary.BigEndian.Uint16(msg[27:])
	
	// .setWarning_tilt( new BigDecimal(getShort(msg[29], msg[30])&0xFFFF) );
	warningTilt := binary.BigEndian.Uint16(msg[29:])

	alramSetting()
	return nil
}

func alramSetting() {
}

func thresholdSettingHandle(msg []byte) (resp []byte) {
	//塔吊编号
	towerCraneId := msg[0];
	//设备编号 
	deviceId := msg[1:4]

	// .setRange_limit_starting_value( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//幅度限位起点值
	range_limit_starting := binary.BigEndian.Uint16(msg[7:])
	
	// .setRange_limit_terminal_value( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//幅度限位终点值
	range_limit_terminal := binary.BigEndian.Uint16(msg[9:])

	// .setHeight_limit_starting_value( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//高度限位起点值
	height_limit_starting := binary.BigEndian.Uint16(msg[11:])

	// .setHeight_limit_terminal_value( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//高度限位终点值
	height_limit_terminal_value := binary.BigEndian.Uint16(msg[13:])
	
	// .setRotation_limit_starting_value( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//回转限位起点值
	rotation_limit_starting_value := binary.BigEndian.Uint16(msg[15:])

	// .setRotation_limit_terminal_value( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//回转限位终点值
	//回转限位终点值
	rotation_limit_terminal := binary.BigEndian.Uint16(msg[17:])
	thresholdSetting()
	return nil
}

func thresholdSetting()  {

}
