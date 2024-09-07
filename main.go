package main

import "fmt"

func main() {
	addStudent(2, "Trần Quốc Toàn", "Nam")
	addStudent(3, "Nguyễn Vũ Nhật Đăng", "Nam")
	updateStudent(3, "Nguyễn Nhật Đăng", "Nữ")
	viewStudentGender()

}

type Students struct {
	ID     int
	Name   string
	Gender string
}

var Student = make(map[int]Students)

func addStudent(id int, name string, gender string) {
	Student[id] = Students{ID: id, Name: name, Gender: gender}
}
func viewAllStudents() {
	if len(Student) == 0 {
		fmt.Println("Không có sinh viên nào trong danh sách.")
		return
	}
	for _, student := range Student {
		fmt.Printf("ID: %d, Tên: %s, Giới tính: %s\n", student.ID, student.Name, student.Gender)
	}
}
func deleteStudent(id int) {
	_, exists := Student[id]
	if exists {
		delete(Student, id)
		fmt.Println("Xóa sinh viên thành công!")
	} else {
		fmt.Println("Sinh viên không tồn tại.")
	}
}
func updateStudent(id int, name string, gender string) {
	_, exists := Student[id]
	if exists {
		Student[id] = Students{ID: id, Name: name, Gender: gender}
		fmt.Println("Sửa thông tin sinh viên thành công!")
	} else {
		fmt.Println("Sinh viên không tồn tại.")
	}
}
func viewStudentGender() {
	for _, student := range Student {
		if student.Gender == "Nam" {
			fmt.Printf("Đây là Nam:ID: %d, Tên: %s, Giới tính: %s\n", student.ID, student.Name, student.Gender)
		} else {
			fmt.Printf("Dây là Nữ:ID: %d, Tên: %s, Giới tính: %s\n", student.ID, student.Name, student.Gender)
		}

	}
}
