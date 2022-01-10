package models

type Constants struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type message struct {
	Th string `json:"th"`
	En string `json:"en"`
}

func Invalid_syntax() Constants {
	return Constants{
		false,
		"Syntax ไม่ถูกต้อง",
	}
}

func ID_not_found() Constants {
	return Constants{
		false,
		"ไม่เจอ ID นี้",
	}
}

func Password_Incorrect() Constants {
	return Constants{
		false,
		"พาสเวิร์ดไม่ถูกต้อง",
	}
}

func Authen_invalid() Constants {
	return Constants{
		false,
		"ชื่อผู้ใช้งานหรือพาสเวิร์ดไม่ถูกต้อง",
	}
}

func Invalid_token() Constants {
	return Constants{
		false,
		"โทเค็นไม่ถูกต้อง",
	}
}

func Token_expired() Constants {
	return Constants{
		false,
		"โทเค็นหมดอายุ",
	}
}

func Token_not_match() Constants {
	return Constants{
		false,
		"โทเค็นไม่ตรงกัน",
	}
}

func Token_not_found() Constants {
	return Constants{
		false,
		"ไม่เจอโทเค็นนี้",
	}
}

func Logout_success() Constants {
	return Constants{
		true,
		"ออกจากระบบสำเร็จ",
	}
}

func Get_data_error() Constants {
	return Constants{
		false,
		"ส่งข้อมูลผิดพลาด",
	}
}

func Cancel_success() Constants {
	return Constants{
		true,
		"ยกเลิกสำเร็จ",
	}
}

func Cancel_error() Constants {
	return Constants{
		false,
		"ยกเลิกไม่สำเร็จ",
	}
}

func Username_Pass() Constants {
	return Constants{
		true,
		"ชื่อผู้ใช้นี้สามารถใช้งานได้",
	}
}

func Username_Not_Pass() Constants {
	return Constants{
		false,
		"ชื่อผู้ใช้นี้ถูกใช้งานไปแล้ว",
	}
}

func Username_Not_Found() Constants {
	return Constants{
		false,
		"ไม่เจอชื่อผู้ใช้นี้",
	}
}

func PhoneNo_Not_Pass() Constants {
	return Constants{
		false,
		"เบอร์โทรศัพท์นี้ถูกใช้งานไปแล้ว",
	}
}

func Get_data_success() Constants {
	return Constants{
		true,
		"รับข้อมูลถูกต้อง",
	}
}

func Delete_picture_success() Constants {
	return Constants{
		true,
		"ลบรูปภาพสำเร็จ",
	}
}

func Save_picture_error() Constants {
	return Constants{
		false,
		"บันทึกรูปภาพผิดพลาด",
	}
}

func Insert_error() Constants {
	return Constants{
		false,
		"บันทึกข้อมูลผิดพลาด",
	}
}

func Insert_success() Constants {
	return Constants{
		true,
		"บันทึกข้อมูลสำเร็จ",
	}
}

func Update_error() Constants {
	return Constants{
		false,
		"บันทึกข้อมูลผิดพลาด",
	}
}
func Have_no_shift() Constants {
	return Constants{
		false,
		"ไม่มีกะในวันดังกล่าว",
	}
}
func File_error() Constants {
	return Constants{
		false,
		"บันทึกไฟล์ผิดพลาด",
	}
}

func Image_Invalid() Constants {
	return Constants{
		false,
		"รูปภาพไม่ถูกต้อง",
	}
}

func EmailAlreadyUsed() Constants {
	return Constants{
		false,
		"อีเมลล์นี้ถูกใช้ไปแล้ว",
	}
}

func Data_not_found() Constants {
	return Constants{
		false,
		"ไม่พบข้อมูล",
	}
}

func Update_success() Constants {
	return Constants{
		true,
		"บันทึกข้อมูลสำเร็จ",
	}
}

func Delete_file_error() Constants {
	return Constants{
		false,
		"ลบไฟล์ล้มเหลว",
	}
}
func Delete_file_success() Constants {
	return Constants{
		true,
		"ลบไฟล์สำเร็จ",
	}
}

func Edit_error() Constants {
	return Constants{
		false,
		"แก้ไขข้อมูลผิดพลาด",
	}
}

func Edit_success() Constants {
	return Constants{
		true,
		"แก้ไขข้อมูลสำเร็จ",
	}
}

func Leave_Day() Constants {
	return Constants{
		false,
		"ไม่สามารถแก้ไขได้วันนั้นเป็นวันหยุด",
	}
}

func Delete_error() Constants {
	return Constants{
		false,
		"ลบข้อมูลผิดพลาด",
	}
}

func Delete_success() Constants {
	return Constants{
		true,
		"ลบข้อมูลสำเร็จ",
	}
}

func Create_token_error() Constants {
	return Constants{
		true,
		"สร้างโทเค็นผิดพลาด",
	}
}

func Change_password_success() Constants {
	return Constants{
		true,
		"เปลี่ยนพาสเวิร์ดสำเร็จ",
	}
}

func Change_password_error() Constants {
	return Constants{
		true,
		"เปลี่ยนพาสเวิร์ดผิดพลาด",
	}
}

func Password_not_match() Constants {
	return Constants{
		true,
		"พาสเวิร์ดไม่ตรงกัน",
	}
}

func Invalid_key() Constants {
	return Constants{
		true,
		"รหัสผ่านไม่ถูกต้อง",
	}
}
