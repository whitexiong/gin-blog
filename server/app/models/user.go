package models

import "blog/app/dao"

//  Model
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

/*
	user这个Model的增删改查操作都放在这里
*/
// CreateAuser 创建todo
func CreateUser(todo *User) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllUser() (todoList []*User, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAUser(id string) (todo *User, err error) {
	todo = new(User)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUser(todo *User) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteAUser(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&User{}).Error
	return
}
