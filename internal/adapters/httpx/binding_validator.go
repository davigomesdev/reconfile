package httpx

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/davigomesdev/reconfile/internal/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	labelCache sync.Map
	messages   = map[string]string{
		"required": "%s é obrigatório",
		"uuid":     "%s inválido",
		"email":    "%s deve ser um e-mail válido",
		"gte":      "%s deve ser maior ou igual a %s",
		"lte":      "%s deve ser menor ou igual a %s",
		"min":      "%s deve ter no mínimo %s caracteres",
		"max":      "%s deve ter no máximo %s caracteres",
		"oneof":    "%s deve ser um dos seguintes valores: %s",
		"numeric":  "%s deve ser um número válido",
		"fqdn":     "%s deve ser um nome de domínio válido",
	}
)

func ParamValidate[T any](ctx *gin.Context) (*T, bool) {
	return bindAndValidate[T](ctx, ctx.ShouldBindUri, "Parâmetros inválidos.")
}

func QueryValidate[T any](ctx *gin.Context) (*T, bool) {
	return bindAndValidate[T](ctx, ctx.ShouldBindQuery, "Parâmetros de pesquisa inválidos.")
}

func BodyValidate[T any](ctx *gin.Context) (*T, bool) {
	return bindAndValidate[T](ctx, ctx.ShouldBindJSON, "Dados da requisição inválido.")
}

func FormValidate[T any](ctx *gin.Context) (*T, bool) {
	return bindAndValidate[T](ctx, ctx.ShouldBind, "Dados da requisição inválido.")
}

func bindAndValidate[T any](ctx *gin.Context, bindFunc func(obj any) error, defaultMsg string) (*T, bool) {
	var dto T

	if err := bindFunc(&dto); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errs := make([]string, len(validationErrors))
			for i, verr := range validationErrors {
				errs[i] = buildErrorMessage(dto, verr)
			}
			ctx.Error(errors.NewUnprocessableEntityError(errs))
		} else {
			ctx.Error(errors.NewUnprocessableEntityError([]string{defaultMsg}))
		}
		return nil, false
	}

	return &dto, true
}

func buildErrorMessage(data any, verr validator.FieldError) string {
	field := verr.StructField()
	label := resolveFieldLabel(data, field)

	if tmpl, exists := messages[verr.Tag()]; exists {
		if param := verr.Param(); param != "" && tmplHasParam(verr.Tag()) {
			return fmt.Sprintf(tmpl, label, param)
		}
		return fmt.Sprintf(tmpl, label)
	}

	return fmt.Sprintf("Erro de validação no campo '%s'", label)
}

func tmplHasParam(tag string) bool {
	switch tag {
	case "gte", "lte", "min", "max", "oneof":
		return true
	default:
		return false
	}
}

func resolveFieldLabel(data any, fieldName string) string {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	labels := getOrCacheFieldLabels(t)
	if label, ok := labels[fieldName]; ok {
		return label
	}
	return fieldName
}

func getOrCacheFieldLabels(t reflect.Type) map[string]string {
	if cached, ok := labelCache.Load(t); ok {
		return cached.(map[string]string)
	}

	labels := make(map[string]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		labels[field.Name] = label
	}

	labelCache.Store(t, labels)
	return labels
}
