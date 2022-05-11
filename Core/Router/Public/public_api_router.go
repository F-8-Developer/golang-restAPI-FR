package Public

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"golang-restAPI-FR/Core/Structs"
	"golang-restAPI-FR/Core/Validator"
	cv "golang-restAPI-FR/Core/Validator/CustomValidation"
	"golang-restAPI-FR/Core/Api"
)

var (
	frq_req Structs.FriendRequestRequest
	frl_req Structs.FriendListRequest
	frb_req Structs.FriendBetweenRequest
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := Validator.InitValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterValidation("email", cv.Email)
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	// end set validator

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to restAPI",
			"userInfo": "Hello World!!!",
			"readme": "https://github.com/F-8-Developer/golang-restAPI-FR/blob/main/README.md",
		})
	})

	// Friend Group
	friend := router.Group("/friend")
	// ============
	
	friend.POST("/request", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frq_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		if err := validate.Struct(frq_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		response := Api.FriendRequest(frq_req)
		c.JSON(200,&response)
		frq_req = Structs.FriendRequestRequest{}
	})

	friend.POST("/accept", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frq_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		if err := validate.Struct(frq_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		response := Api.UpdateFriendRequest(frq_req, "pending", "accepted")
		c.JSON(200,&response)
		frq_req = Structs.FriendRequestRequest{}
	})

	friend.POST("/reject", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frq_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		if err := validate.Struct(frq_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		response := Api.UpdateFriendRequest(frq_req, "pending", "rejected")
		c.JSON(200,&response)
		frq_req = Structs.FriendRequestRequest{}
	})

	friend.POST("/block", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frq_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		if err := validate.Struct(frq_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frq_req = Structs.FriendRequestRequest{}
			return
		}

		response := Api.BlockFriend(frq_req)
		c.JSON(200,&response)
		frq_req = Structs.FriendRequestRequest{}
	})

	friend.POST("/list-request", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frl_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frl_req = Structs.FriendListRequest{}
			return
		}

		if err := validate.Struct(frl_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frl_req = Structs.FriendListRequest{}
			return
		}

		response := Api.ListFriendRequest(frl_req)
		c.JSON(200,&response)
		frl_req = Structs.FriendListRequest{}
	})

	friend.POST("/list-friends", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frl_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frl_req = Structs.FriendListRequest{}
			return
		}

		if err := validate.Struct(frl_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frl_req = Structs.FriendListRequest{}
			return
		}

		response := Api.ListFriendRequest(frl_req)
		c.JSON(200,&response)
		frl_req = Structs.FriendListRequest{}
	})

	friend.POST("/list-friends-between", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&frb_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": err.Error(),
			})
			frb_req = Structs.FriendBetweenRequest{}
			return
		}

		if err := validate.Struct(frb_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors": errs.Errors,
			})
			frb_req = Structs.FriendBetweenRequest{}
			return
		}

		// validate 2 email
		if len(frb_req.Friends) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error": "must between 2 email user",
			})
			frb_req = Structs.FriendBetweenRequest{}
			return
		}

		response := Api.ListFriendBetween(frb_req)
		c.JSON(200,&response)
		frb_req = Structs.FriendBetweenRequest{}
	})
	// End Friend Group
}