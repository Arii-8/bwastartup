package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)           // Menyimpan user (save user)
	FindByEmail(email string) (User, error) // Cari user berdaarkan email
	FindByID(ID int) (User, error)          // Cari user berdasarkan id
	Update(user User) (User, error)         // Update
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Fungsi save user
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Fungsi find email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error // menemukan email string dengan format 'email' berdasarkan single user

	// check jika error
	if err != nil {
		return user, err
	}
	// jika sukses
	return user, nil
}

// Fungsi find ID
func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error // menemukan user dengan berdasarkan 'id' berdasarkan single user

	// check jika error
	if err != nil {
		return user, err
	}
	// jika sukses
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	// check jika error
	if err != nil {
		return user, err
	}

	// jika sukses
	return user, nil
}
