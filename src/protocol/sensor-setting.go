package protocol

import (
	// "errors"
	"encoding/binary"
	"log"
)

const (
	MsgSensorSetting     byte = 0x33 //传感器标定信息 0x33 = 51
    MsgAlramSetting    byte = 0x35 //报警设置信息
	MsgThresholdSetting    byte = 0x37 //限位设置信息 0x37
)

func init() {
	if _, ok := msgHanlders[MsgSensorSetting]; ok == true {
        log.Printf("can't add function sensorSettingHandle, the key %2X has exits\n", MsgSensorSetting)   
    } else {
        msgHanlders[MsgSensorSetting] = sensorSettingHandle
	}
	
	if _, ok := msgHanlders[MsgAlramSetting]; ok == true {
        log.Printf("can't add function alarmSettingHandle, the key %2X has exits\n", MsgAlramSetting)   
	} else {
			msgHanlders[MsgAlramSetting] = alarmSettingHandle
	}
	if _, ok := msgHanlders[MsgThresholdSetting]; ok == true {
		log.Printf("can't add function thresholdSettingHandle, the key %2X has exits\n", MsgThresholdSetting)   
	} else {
			msgHanlders[MsgThresholdSetting] = thresholdSettingHandle
	}	
}

type SensorSetting struct {
	towerCraneId byte
	deviceId string
	proximal_amplitude_demarcate_AD uint16
	proximal_amplitude_demarcate_actual uint16
	remote_service_demarcate_AD uint16
	remote_service_demarcate_actual uint16
	proximal_height_demarcate_AD uint16
	proximal_height_demarcate_actual uint16
	remoteHeightDemarcateAD uint16
	remote_height_demarcate_actual uint16
	noLoad_weight_AD uint16
	noLoad_weight_actual uint16
	load_weight_AD uint16
	load_weight_actual uint16
	rotation_starting_point_AD uint16
	rotation_starting_point_actual uint16
	rotationEndinigPointAD uint16
	rotation_endinig_point_actual uint16

	windSpeedCalibration uint16

	tilt_the_calibration uint16
}
func sensorSettingHandle(msg []byte) (resp []byte) {
	if len(msg) < 53 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}
	setting := SensorSetting {}
	//塔吊编号
	setting.towerCraneId = msg[1];
	//设备编号 
	setting.deviceId = BytetoH(msg[2:5])
	
	// .setProximal_amplitude_demarcate_AD( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF) );
	setting.proximal_amplitude_demarcate_AD = binary.BigEndian.Uint16(msg[5:])

	// .setProximal_amplitude_demarcate_actual( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.proximal_amplitude_demarcate_actual = binary.BigEndian.Uint16(msg[7:])

	// .setRemote_service_demarcate_AD( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF) );
	setting.remote_service_demarcate_AD = binary.BigEndian.Uint16(msg[9:])

	// .setRemote_service_demarcate_actual( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.remote_service_demarcate_actual = binary.BigEndian.Uint16(msg[11:])

	// .setProximal_height_demarcate_AD( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF) );
	setting.proximal_height_demarcate_AD = binary.BigEndian.Uint16(msg[13:])

	// .setProximal_height_demarcate_actual( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.proximal_height_demarcate_actual = binary.BigEndian.Uint16(msg[15:])

	// .setRemote_height_demarcate_AD( new BigDecimal(getShort(msg[19], msg[20])&0xFFFF) );
	setting.remoteHeightDemarcateAD = binary.BigEndian.Uint16(msg[17:])

	// .setRemote_height_demarcate_actual( new BigDecimal(getShort(msg[21], msg[22])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.remote_height_demarcate_actual = binary.BigEndian.Uint16(msg[19:])

	// .setNoLoad_weight_AD( new BigDecimal(getShort(msg[23], msg[24])&0xFFFF) );
	setting.noLoad_weight_AD = binary.BigEndian.Uint16(msg[21:])

	// .setNoLoad_weight_actual( new BigDecimal(getShort(msg[25], msg[26])&0xFFFF) );
	setting.noLoad_weight_actual = binary.BigEndian.Uint16(msg[23:])

	// .setLoad_weight_AD( new BigDecimal(getShort(msg[27], msg[28])&0xFFFF) );
	setting.load_weight_AD = binary.BigEndian.Uint16(msg[25:])

	// 原始代码有错，把实际重量放到空载重量中去了
	// .setNoLoad_weight_actual( new BigDecimal(getShort(msg[29], msg[30])&0xFFFF) );
	setting.load_weight_actual = binary.BigEndian.Uint16(msg[27:])
	
	// .setRotation_starting_point_AD( new BigDecimal(getShort(msg[31], msg[32])&0xFFFF) );
	setting.rotation_starting_point_AD = binary.BigEndian.Uint16(msg[29:])

	// .setRotation_starting_point_actual( new BigDecimal(getShort(msg[33], msg[34])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.rotation_starting_point_actual = binary.BigEndian.Uint16(msg[31:])

	// .setRotation_endinig_point_AD( new BigDecimal(getShort(msg[35], msg[36])&0xFFFF) );
	setting.rotationEndinigPointAD = binary.BigEndian.Uint16(msg[33:])
	
	// .setRotation_endinig_point_actual( new BigDecimal(getShort(msg[37], msg[38])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.rotation_endinig_point_actual = binary.BigEndian.Uint16(msg[35:])
	
	// .setWind_speed_calibration( new BigDecimal(getShort(msg[39], msg[40])&0xFFFF) );
	//风速校准系数（2byte）
	setting.windSpeedCalibration = binary.BigEndian.Uint16(msg[37:])
	
	// .setTilt_the_calibration( new BigDecimal(getShort(msg[47], msg[48])&0xFFFF) );
	//倾斜校准系数（2byte）
	setting.tilt_the_calibration = binary.BigEndian.Uint16(msg[45:])
	
	reportSensorSetting(setting)

	return nil
}

func reportSensorSetting(setting SensorSetting)  {
	log.Printf("sensor setting:%+v\n", setting)
	
}


type AlarmSetting struct {
	//塔吊编号
	towerCraneId byte
	deviceId string
	horizontal_distance uint16
	perpendicular_distance uint16
	weight_distance uint16
	torque uint16
	windSpeed uint16
	alarmTilt uint16
	warningHorizontalDistance uint16
	warningPerpendicularDistance uint16
	warningWeightPercentage uint16
	warningTorque uint16
	warningWindSpeed uint16
	warningTilt uint16

}
func alarmSettingHandle(msg []byte) (resp []byte) {
	if len(msg) < 29 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}
	setting := AlarmSetting{}
	//塔吊编号
	setting.towerCraneId = msg[1];
	//设备编号 
	setting.deviceId = BytetoH(msg[2:5])

	// .setAlarm_horizontal_distance( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.horizontal_distance = binary.BigEndian.Uint16(msg[5:])
	
	// .setAlarm_perpendicular_distance( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.perpendicular_distance = binary.BigEndian.Uint16(msg[7:])
	
	// .setAlarm_weight_distance( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF) );
	setting.weight_distance = binary.BigEndian.Uint16(msg[9:])

	// .setAlarm_torque( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF) );
	setting.torque = binary.BigEndian.Uint16(msg[11:])

	// .setAlarm_wind_speed( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.windSpeed = binary.BigEndian.Uint16(msg[13:])

	// .setAlarm_tilt( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF) );
	setting.alarmTilt = binary.BigEndian.Uint16(msg[15:])
	
	// .setWarning_horizontal_distance( new BigDecimal(getShort(msg[19], msg[20])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.warningHorizontalDistance = binary.BigEndian.Uint16(msg[17:])
	
	// .setWarning_perpendicular_distance( new BigDecimal(getShort(msg[21], msg[22])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.warningPerpendicularDistance = binary.BigEndian.Uint16(msg[19:])

	// .setWarning_weight_percentage( new BigDecimal(getShort(msg[23], msg[24])&0xFFFF) );
	setting.warningWeightPercentage = binary.BigEndian.Uint16(msg[21:])
	
	// .setWarning_torque( new BigDecimal(getShort(msg[25], msg[26])&0xFFFF) );
	setting.warningTorque = binary.BigEndian.Uint16(msg[23:])

	// .setWarning_wind_speed( new BigDecimal(getShort(msg[27], msg[28])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );
	setting.warningWindSpeed = binary.BigEndian.Uint16(msg[25:])
	
	// .setWarning_tilt( new BigDecimal(getShort(msg[29], msg[30])&0xFFFF) );
	setting.warningTilt = binary.BigEndian.Uint16(msg[27:])

	reportAlarmSetting(setting)
	return nil
}

func reportAlarmSetting(setting AlarmSetting) {
	log.Printf("alarm setting:%+v\n", setting)
}

type Threshold struct{
	towerCraneId byte
	deviceId string
	range_limit_starting uint16
	range_limit_terminal uint16
	height_limit_starting uint16
	height_limit_terminal_value uint16
	rotation_limit_starting_value uint16
	rotation_limit_terminal uint16
}
func thresholdSettingHandle(msg []byte) (resp []byte) {
	if len(msg) < 17 {
		log.Printf("incorrect message: %s\n", BytetoH(msg))
		return nil
	}
	setting := Threshold{}
	//塔吊编号
	setting.towerCraneId = msg[1];
	//设备编号 
	setting.deviceId = BytetoH(msg[2:5])

	// .setRange_limit_starting_value( new BigDecimal(getShort(msg[7], msg[8])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//幅度限位起点值
	setting.range_limit_starting = binary.BigEndian.Uint16(msg[5:])
	
	// .setRange_limit_terminal_value( new BigDecimal(getShort(msg[9], msg[10])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//幅度限位终点值
	setting.range_limit_terminal = binary.BigEndian.Uint16(msg[7:])

	// .setHeight_limit_starting_value( new BigDecimal(getShort(msg[11], msg[12])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//高度限位起点值
	setting.height_limit_starting = binary.BigEndian.Uint16(msg[9:])

	// .setHeight_limit_terminal_value( new BigDecimal(getShort(msg[13], msg[14])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//高度限位终点值
	setting.height_limit_terminal_value = binary.BigEndian.Uint16(msg[11:])
	
	// .setRotation_limit_starting_value( new BigDecimal(getShort(msg[15], msg[16])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//回转限位起点值
	setting.rotation_limit_starting_value = binary.BigEndian.Uint16(msg[13:])

	// .setRotation_limit_terminal_value( new BigDecimal(getShort(msg[17], msg[18])&0xFFFF).setScale(1, BigDecimal.ROUND_UNNECESSARY).divide(BigDecimalSet.b10) );//回转限位终点值
	//回转限位终点值
	setting.rotation_limit_terminal = binary.BigEndian.Uint16(msg[15:])
	
	reportThreshold(setting)
	
	return nil
}

func reportThreshold(setting Threshold)  {
	log.Printf("threshold setting:%+v\n", setting)
}
