package usecases

import (
	"fmt"
	"testing"

	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type mockAdminRepo struct {
	mock.Mock
	createAdminFunc     func(admin entities.Admin) (entities.Admin, error)
	getAdminsFunc       func() ([]entities.Admin, error)
	getAdminFunc        func(id int) (entities.Admin, error)
	updateAdminFunc     func(id int, admin entities.Admin) (entities.Admin, error)
	deleteAdminFunc     func(id int) (entities.Admin, error)
	getAdminByEmailFunc func(email string) (entities.Admin, error)
}

func (m *mockAdminRepo) CreateAdmin(admin entities.Admin) (entities.Admin, error) {
	return m.createAdminFunc(admin)
}

func (m *mockAdminRepo) GetAdmins() ([]entities.Admin, error) {
	return m.getAdminsFunc()
}

func (m *mockAdminRepo) GetAdmin(id int) (entities.Admin, error) {
	return m.getAdminFunc(id)
}

func (m *mockAdminRepo) UpdateAdmin(id int, admin entities.Admin) (entities.Admin, error) {
	return m.updateAdminFunc(id, admin)
}

func (m *mockAdminRepo) DeleteAdmin(id int) (entities.Admin, error) {
	return m.deleteAdminFunc(id)
}

// GetAdminByEmail implements repositories.AdminRepository.
func (m *mockAdminRepo) GetAdminByEmail(email string) (entities.Admin, error) {
	return m.getAdminByEmailFunc(email)
}

func TestCreateAdmin(t *testing.T) {
	// Success case
	// t.Run("success", func(t *testing.T) {
	// 	app := fiber.New()

	// 	c := app.AcquireCtx(&fasthttp.RequestCtx{})

	// 	repo := &mockAdminRepo{
	// 		createAdminFunc: func(admin entities.Admin) (entities.Admin, error) {
	// 			// Simulate successful save
	// 			return entities.Admin{
	// 				FullName: "aut",
	// 				Email:    "aut@gmail.com",
	// 				Password: "atchima1234",
	// 			}, nil
	// 		},
	// 	}

	// 	service := NewAdminUseCase(repo)

	// 	adminInput := entities.Admin{
	// 		FullName: "aut",
	// 		Email:    "aut@gmail.com",
	// 		Password: "atchima1234",
	// 	}
	// 	mockImage := multipart.FileHeader{

	// 	}
	// 	result, err := service.CreateAdmin(
	// 		adminInput,
	// 		*mockImage,
	// 		c,
	// 	)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, adminInput.FullName, result.FullName)
	// 	assert.Equal(t, adminInput.Email, result.Email)
	// 	assert.NotEmpty(t, result.Password)
	// })

	// t.Run("failure", func(t *testing.T) {
	// 	app := fiber.New()
	// 	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	// 	repo := &mockAdminRepo{
	// 		createAdminFunc: func(admin entities.Admin) (entities.Admin, error) {
	// 			return entities.Admin{}, fmt.Errorf("Failed to create admin")
	// 		},
	// 	}

	// 	service := NewAdminUseCase(repo)
	// 	mockImage := createMockFile(t, "image.jpg")

	// 	_, err := service.CreateAdmin(
	// 		entities.Admin{
	// 			FullName: "aut",
	// 			Email:    "aut@gmail.com",
	// 			Password: "atchima1234",
	// 		},
	// 		*mockImage,
	// 		c,
	// 	)
	// 	assert.Error(t, err)
	// })
	panic("unimplemented")
}

func TestGetAdmins(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		outputData := []entities.Admin{
			{FullName: "aut", Email: "aut@gmail.com"},
			{FullName: "bee", Email: "bee@gmail.com"},
		}
		repo := &mockAdminRepo{
			getAdminsFunc: func() ([]entities.Admin, error) {
				return outputData, nil
			},
		}
		service := NewAdminUseCase(repo)

		result, err := service.GetAdmins()

		assert.NoError(t, err)
		assert.EqualValues(t, outputData, result)
	})

	t.Run("not found or error", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminsFunc: func() ([]entities.Admin, error) {
				return nil, fmt.Errorf("not found or error")
			},
		}
		service := NewAdminUseCase(repo)

		result, err := service.GetAdmins()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "not found or error", err.Error())
	})
}

func TestGetAdmin(t *testing.T) {
	expectData := entities.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
		Image:    "autimage",
	}

	t.Run("success", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return expectData, nil
			},
		}
		service := NewAdminUseCase(repo)

		result, err := service.GetAdmin(1)

		assert.NoError(t, err)
		assert.EqualValues(t, expectData, result)
	})

	t.Run("not found or error", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.GetAdmin(2)

		assert.Error(t, err)
		assert.Equal(t, "not found or error", err.Error())
	})
}

func TestUpdateAdmin(t *testing.T) {
	expectData := entities.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
		Image:    "autimage",
	}

	t.Setenv("JWT_SECRET_KEY", "test secret key")
	t.Run("can't extract token", func(t *testing.T) {
		repo := &mockAdminRepo{}
		service := NewAdminUseCase(repo)

		token := "whattoken"
		_, err := service.UpdateAdmin(token, expectData, nil, nil)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to extract token")
	})

	t.Run("get admin not found or error", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.GetAdmin(2)

		assert.Error(t, err)
		assert.Equal(t, "not found or error", err.Error())
	})

	t.Run("success: skip image upload if not provide image", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found or error")
			},
			updateAdminFunc: func(id int, admin entities.Admin) (entities.Admin, error) {
				if id == 1 {
					return admin, nil
				}
				return entities.Admin{}, fmt.Errorf("failed to update admin")
			},
		}
		service := NewAdminUseCase(repo)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzcyMTM1MTF9.3GAecbqrWXrcY2-cw-G9_8rdRtFfu9_O_3Nr3lvwK5c"
		result, err := service.UpdateAdmin(token, expectData, nil, nil)

		assert.NoError(t, err)
		assert.EqualValues(t, expectData, result)
	})

	///////success: upload image if never upload///////

	///////success: update image if ever upload///////

	t.Run("failed to update", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found or error")
			},
			updateAdminFunc: func(id int, admin entities.Admin) (entities.Admin, error) {
				if id == 2 {
					return admin, nil
				}
				return entities.Admin{}, fmt.Errorf("failed to update admin")
			},
		}
		service := NewAdminUseCase(repo)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzcyMTM1MTF9.3GAecbqrWXrcY2-cw-G9_8rdRtFfu9_O_3Nr3lvwK5c"
		_, err := service.UpdateAdmin(token, expectData, nil, nil)

		assert.Error(t, err)
		assert.Equal(t, "failed to update admin", err.Error())
	})
}

func TestDeleteAdmin(t *testing.T) {
	expectData := entities.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
		Image:    "autimage",
	}

	t.Setenv("SUPA_API_URL", "https://azorxaeszmxrlsoeshfg.supabase.co/storage/v1")
	t.Setenv("SUPA_API_KEY", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImF6b3J4YWVzem14cmxzb2VzaGZnIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzU2MzQ1NTYsImV4cCI6MjA1MTIxMDU1Nn0.UxWD_NlyhpHtpkjeeOVrERu_WO0nmcKl_tH3_uMLrTo")
	t.Setenv("SUPA_BUCKET_NAME", "skincare")
	t.Run("failed to get admin", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.DeleteAdmin(2)

		assert.Error(t, err)
		assert.Equal(t, "not found this admin or error", err.Error())
	})

	////////Make add image to supabase error
	t.Run("failed to remove admin image from supabase", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.DeleteAdmin(1)

		assert.Error(t, err)
	})

	t.Run("failed to delete admin", func(t *testing.T) {
		repo := &mockAdminRepo{
		    getAdminFunc: func(id int) (entities.Admin, error) {
			  if id == 1 {
				return expectData, nil // คืนข้อมูล admin ถ้า ID ตรง
			  }
			  return entities.Admin{}, fmt.Errorf("not found this admin")
		    },
		    deleteAdminFunc: func(id int) (entities.Admin, error) {
			  // จำลองกรณีลบ admin ไม่สำเร็จ
			  if id == 1 {
				return entities.Admin{}, fmt.Errorf("failed to delete admin") // คืนข้อผิดพลาด
			  }
			  return entities.Admin{}, fmt.Errorf("admin not found")
		    },
		}
	  
		service := NewAdminUseCase(repo)
	  
		result, err := service.DeleteAdmin(1)
	  
		// ตรวจสอบว่าเกิดข้อผิดพลาด
		assert.Error(t, err)
		assert.Equal(t, "failed to delete admin", err.Error()) // ควรจะได้ข้อความ error ตามที่คาดไว้
		assert.Equal(t, entities.Admin{}, result) // ผลลัพธ์ต้องเป็น admin ว่างๆ (ไม่ได้ลบ)
	})
	  

	t.Run("success", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
			deleteAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil // Mock successful deletion
				}
				return entities.Admin{}, fmt.Errorf("failed to delete admin")
			},
		}
		service := NewAdminUseCase(repo)

		// Test success case
		result, err := service.DeleteAdmin(1)

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
	})
}

func TestLogin(t *testing.T) {
	expectData := entities.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "$2b$12$8/XDIvr.9mE0x2bWozPzYem0QeAERA1qrunUyh5/5RhOSXDG2/6fu", //aut1234
		Image:    "autimage",
	}

	t.Setenv("JWT_SECRET_KEY", "test-key")

	t.Run("failed to get admin", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminByEmailFunc: func(email string) (entities.Admin, error) {
				if email != expectData.Email {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.LogIn(expectData.Email, expectData.Password)

		assert.Error(t, err)
		assert.Equal(t, "not found this admin or error", err.Error())
	})
	  
	t.Run("incorrect password or failed compare hash password", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminByEmailFunc: func(email string) (entities.Admin, error) {
				if email == expectData.Email {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.LogIn(expectData.Email, "aut12")

		assert.Error(t, err)
	})

	//////Make it fail in generate token
	t.Run("failed to generate token", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminByEmailFunc: func(email string) (entities.Admin, error) {
				if email == expectData.Email {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		_, err := service.LogIn(expectData.Email, "aut1234")

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminByEmailFunc: func(email string) (entities.Admin, error) {
				if email == expectData.Email {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)
	
		_, err := service.LogIn(expectData.Email, "aut1234")
	
		assert.NoError(t, err)
	})
}

func TestGetAdminByToken(t *testing.T) {
	expectData := entities.Admin{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "$2b$12$8/XDIvr.9mE0x2bWozPzYem0QeAERA1qrunUyh5/5RhOSXDG2/6fu", //aut1234
		Image:    "autimage",
	}

	t.Setenv("JWT_SECRET_KEY", "test secret key")
	t.Run("failed to extract token", func(t *testing.T) {
		repo := &mockAdminRepo{}
		service := NewAdminUseCase(repo)

		_, err := service.GetAdminByToken("wrong-token")

		assert.Error(t, err)
	})

	t.Run("failed to get admin", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id != 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzcyMTM1MTF9.3GAecbqrWXrcY2-cw-G9_8rdRtFfu9_O_3Nr3lvwK5c"
		_, err := service.GetAdminByToken(token)

		assert.Error(t, err)
		assert.Equal(t, "not found this admin or error", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		repo := &mockAdminRepo{
			getAdminFunc: func(id int) (entities.Admin, error) {
				if id == 1 {
					return expectData, nil
				}
				return entities.Admin{}, fmt.Errorf("not found this admin or error")
			},
		}
		service := NewAdminUseCase(repo)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MzcyMTM1MTF9.3GAecbqrWXrcY2-cw-G9_8rdRtFfu9_O_3Nr3lvwK5c"
		result, err := service.GetAdminByToken(token)

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
	})
}