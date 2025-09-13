package handlers

import (
	"encoding/json"
	"net/http"

	usecases_letter_soup "lettersoup/src/application/modules/letter_soup/use_cases"
	dtos "lettersoup/src/application/shared/DTOs"
	"lettersoup/src/domain/models"
	database "lettersoup/src/infrastructure/database/lettersoup"
	"lettersoup/src/infrastructure/providers"
	"lettersoup/src/infrastructure/repositories"
)

// CreateLetterSoup
// @Summary This endpoint creates a letter soup puzzle
// @Description This endpoint creates a letter soup puzzle based on the provided words and grid size
// @Tags LetterSoup
// @Accept json
// @Produce json
// @Param request body dtos.LetterSoupCreate true "Datos para crear la sopa de letras"
// @Param Accept-Language header string false "Locale for response messages" Enums(en-US, es-ES) default(en-US)
// @Success 201 {object} models.LetterSoup "Sopa de letras creada"
// @Failure 400 {object} map[string]string "Error de validación"
// @Router /api/letter-soup [post]
// @Security Bearer
func CreateLetterSoup(ctx HandlerContext) {
	var letterSoupCreate dtos.LetterSoupCreate

	if err := json.NewDecoder(*ctx.Body).Decode(&letterSoupCreate); err != nil {
		http.Error(ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	ucResult := usecases_letter_soup.NewCreateLetterSoupUseCase(
		providers.Logger,
		repositories.NewLetterSoupRepository(database.DB, providers.Logger),
	).Execute(ctx.c, ctx.Locale, letterSoupCreate)

	headers := map[HTTPHeaderTypeEnum]string{
		CONTENT_TYPE: string(APPLICATION_JSON),
	}
	NewRequestResolver[models.LetterSoup]().ResolveDTO(ctx.ResponseWriter, ucResult, headers)
}
