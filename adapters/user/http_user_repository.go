package adapters

import (
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/Narutchai01/Project_S-BE/presentation"
	"github.com/Narutchai01/Project_S-BE/usecases"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	userUcase usecases.UserUsecases
}

func NewHttpUserHandler(userUcase usecases.UserUsecases) *HttpUserHandler {
	return &HttpUserHandler{userUcase}
}

func validateUser(user entities.User) error {
	if user.Email == "" || user.Password == "" || user.FullName == "" || user.SensitiveSkin == nil || user.Birthday == nil {
		return fiber.ErrBadRequest
	}
	return nil

}

// Register godoc
//
//	@Summary		Register new user
//	@Description	Register new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entities.User	true	"User information"
//	@Success		201		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/user/register [post]
func (handler *HttpUserHandler) Register(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	if err := validateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}

	result, err := handler.userUcase.Register(user, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(presentation.UserResponse(result))
}

// LogIn godoc
//
//	@Summary		Log in
//	@Description	Log in
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		object{email=string,password=string}	true	"Admin Object"
//
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/user/login [post]
func (handler *HttpUserHandler) LogIn(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.userUcase.LogIn(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.TokenResponse(result))
}

func (handler *HttpUserHandler) ForgetPassword(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.userUcase.ChangePassword(int(user.ID), user.Password, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.UserResponse(result))
}

// GoogleSignIn godoc
//
//	@Summary		Google sign in
//	@Description	Google sign in
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		object{email=string,fullname=string,image=string,sensitive_skin=boolean}	true	"Admin Object"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/user/goolge-signin [post]
func (handler *HttpUserHandler) GoogleSignIn(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	if user.Email == "" || user.FullName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(fiber.ErrBadRequest))
	}
	result, err := handler.userUcase.GoogleSignIn(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.TokenResponse(result))
}

// GetUser godoc
//
//	@Summary		Get user by id
//	@Description	Get user by id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string	true	"Token"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/user/me [get]
func (handler *HttpUserHandler) GetUser(c *fiber.Ctx) error {

	token := c.Get("token")

	result, err := handler.userUcase.GetUser(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.UserResponse(result))
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	header		string			true	"Token"
//	@Param			user	body		entities.User	true	"User information"
//	@Success		200		{object}	presentation.Responses
//	@Failure		400		{object}	presentation.Responses
//	@Failure		404		{object}	presentation.Responses
//	@Router			/user/ [put]
func (handler *HttpUserHandler) UpdateUser(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	token := c.Get("token")

	file, _ := c.FormFile("file")

	result, err := handler.userUcase.UpdateUser(user, token, file, c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler *HttpUserHandler) Follower(c *fiber.Ctx) error {
	token := c.Get("token")
	folllow_id, err := c.ParamsInt("follow_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	result, err := handler.userUcase.Follower(uint(folllow_id), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ToFollowerResponse(result))
}

func (handler *HttpUserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ErrorResponse(err))
	}

	token := c.Get("token")

	result, err := handler.userUcase.GetUserByID(uint(id), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.UserResponse(result))
}
