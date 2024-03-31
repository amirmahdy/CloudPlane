package api

import (
	db "cloudplane/db/model"
	"cloudplane/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createProfileRequest struct {
	Region      string `json:"region" binding:"required,alphanum"`
	AccessID    string `json:"access_id" binding:"required,min=6"`
	SecretKey   string `json:"secret_key" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type createProfileResponse struct {
	ProfID string `json:"prof_id"`
}

// @Summary Profile
// @Schemes
// @Description profile create
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createProfileRequest body createProfileRequest true "Create Profile Param"
// @Success 200 {object} createProfileResponse
// @Router /profile/create [post]
func (server *Server) createProfile(ctx *gin.Context) {
	var req createProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authPayload := ctx.MustGet("authorization_payload").(*token.Payload)
	profParam := db.CreateProfileTXParam{
		Region:      req.Region,
		AccessID:    req.AccessID,
		SecretKey:   req.SecretKey,
		Description: req.Description,
		Username:    authPayload.UserID,
	}

	profRes, err := server.store.CreateProfileTX(profParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ur := createProfileResponse{
		ProfID: profRes.ProfID.String(),
	}
	ctx.JSON(http.StatusOK, ur)
}
