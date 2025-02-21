package adapters_test

// func TestBookmarkThread(t *testing.T) {
// 	setup := func() (*MockThreadUseCase, *adapters.HttpThreadHandler, *fiber.App) {
// 		mockThreadUseCase := new(MockThreadUseCase)
// 		httpThreadHandler := adapters.NewHttpThreadHandler(mockThreadUseCase)
// 		app := fiber.New()

// 		app.Post("/thread/:id/bookmark", httpThreadHandler.BookMark)

// 		return mockThreadUseCase, httpThreadHandler, app
// 	}

// 	t.Run("Success", func(t *testing.T) {
// 		mockThreadUseCase, _, app := setup()

// 		mockBookmark := entities.Bookmark{
// 			ThreadID: 1,
// 			UserID:   1,
// 		}

// 		mockThreadUseCase.On("AddBookmark", uint(1), "test-token").Return(mockBookmark, nil)

// 		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)
// 		req.Header.Set("token", "test-token")

// 		resp, _ := app.Test(req)

// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
// 		mockThreadUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Invalid ID", func(t *testing.T) {
// 		_, _, app := setup()

// 		req := httptest.NewRequest(fiber.MethodPost, "/thread/invalid/bookmark", nil)
// 		req.Header.Set("token", "test-token")

// 		resp, _ := app.Test(req)

// 		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
// 	})

// 	t.Run("Missing Token", func(t *testing.T) {
// 		_, _, app := setup()

// 		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)

// 		resp, _ := app.Test(req)

// 		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
// 	})

// 	t.Run("Error Adding Bookmark", func(t *testing.T) {
// 		mockThreadUseCase, _, app := setup()

// 		mockThreadUseCase.On("AddBookmark", uint(1), "test-token").Return(entities.Bookmark{}, errors.New("invalid thread ID"))

// 		req := httptest.NewRequest(fiber.MethodPost, "/thread/1/bookmark", nil)
// 		req.Header.Set("token", "test-token")

// 		resp, _ := app.Test(req)

// 		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
// 		mockThreadUseCase.AssertExpectations(t)
// 	})
// }
