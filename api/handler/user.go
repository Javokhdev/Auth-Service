package handler

import (
	"github.com/Javokhdev/Auth-Service/api/token"

	pb "github.com/Javokhdev/Auth-Service/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateUser 		handles the creation of a new user
// @Summary 		Create User
// @Description 	Create page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.Users    true   "Create"
// @Success 		200   {string}   pb.Users   "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/user/registr [post]
func (h *Handler) RegisterUser(ctx *gin.Context){
		arr:=pb.Users{}
		err:=ctx.BindJSON(&arr)
		if err!=nil{
			panic(err)
		}
		_, err=h.User.CreateUser(ctx, &arr)
		if err!=nil{
			panic(err)
		}
	t:=token.GenereteJWTToken(&arr)
	ctx.JSON(200, t)
}

// UpdateUser 		handles the creation of a new user
// @Summary			Update User
// @Description 	Update page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id 		path   string     true   "User ID"
// @Param   		Update  body   pb.Users    true   "Update"
// @Success 		200   {string} string    "Update Successful"
// @Failure 		401   {string} string    "Error while created"
// @Router 			/user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	arr:=pb.Users{}
	id:=ctx.Param("id")
	arr.Id=id
	err:=ctx.BindJSON(&arr)
	if err!=nil{
		panic(err)
	}
	_, err=h.User.UpdateUser(ctx, &arr)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}


// DeleteUser 		handles the creation of a new User
// @Summary			Delete User
// @Description 	Delete page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id   path     string   true   "User ID"
// @Success 		200 {string}  string   "Delete Successful"
// @Failure 		401 {string}  string   "Error while Deleted"
// @Router 			/user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	_, err:=h.User.DeleteUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, "Success!!!")
}

// GetAllUser 		handles the creation of a new User
// @Summary 		GetAll User
// @Description 	GetAll page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			query query pb.Users true    "Query parameter"
// @Success 		200 {object}  pb.GetAllUsers  "GetAll Successful"
// @Failure 		401 {string}  string  		  "Error while GetAll"
// @Router 			/user/getall  [get]
func (h *Handler) GetAllUser(ctx *gin.Context){
	user := &pb.Users{}
	user.Email = ctx.Param("email")
	user.Password = ctx.Param("password")
	user.Username = ctx.Param("username")

	res, err:=h.User.GetAllUser(ctx, user)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		GetById User
// @Description 	GetById page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id   path      string   true    "User ID"
// @Success 		200 {object}   pb.Users  "GetById Successful"
// @Failure 		401 {string}   string   "Error while GetByIdd"
// @Router 			/user/getbyid/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context){
	id:=pb.ById{Id: ctx.Param("id")}
	res, err:=h.User.GetByIdUser(ctx, &id)
	if err!=nil{
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetByIdUser 		handles the creation of a new User
// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body  pb.Users   true     "Create"
// @Success 		200 {object}  pb.Users  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context){
	user:=&pb.Users{}
	err:=ctx.ShouldBindJSON(user)
	if err!=nil{
		panic(err)
	}
	res, err:=h.User.LoginUser(ctx, user)
	if err!=nil{
		panic(err)
	}
	t:=token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}

