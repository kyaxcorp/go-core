package helper

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/kyaxcorp/go-core/core/helpers/_context"
	"github.com/kyaxcorp/go-core/core/helpers/errors2/define"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/authentication"
	"github.com/kyaxcorp/go-core/core/listeners/http/middlewares/connection"
)

type ResolverHelper struct {
	Ctx                       context.Context
	DisableGetRequestedFields bool
	httpCtx                   *gin.Context
	connDetails               *connection.ConnDetails
	authDetails               *authentication.AuthDetails
	requestedFields           []string
}

func NewResolverHelper(rh *ResolverHelper) *ResolverHelper {
	if rh == nil {
		rh = &ResolverHelper{}
	}

	// Check if context is present...
	if rh.Ctx == nil {
		rh.Ctx = _context.GetDefaultContext()
	}

	// call it once...
	_, _err := rh.GetGinContext()
	if _err != nil {
		panic(define.Err(0, "failed to get gin context -> ", _err.Error()))
		//return nil, define.Err(0, "failed to get gin context -> ", _err.Error())
	}
	rh.GetAuthDetails()
	rh.GetConnectionDetails()
	if !rh.DisableGetRequestedFields {
		rh.GetRequestedFields()
	}
	return rh
}

func (r *ResolverHelper) GetConnectionDetails() *connection.ConnDetails {
	if r.connDetails != nil {
		return r.connDetails
	}
	r.connDetails = connection.GetConnectionDetailsFromCtx(r.httpCtx)
	return r.connDetails
}

func (r *ResolverHelper) GetRequestedFields() []string {
	if r.requestedFields != nil {
		return r.requestedFields
	}
	r.requestedFields = GetPreloads(r.Ctx)
	return r.requestedFields
}

func (r *ResolverHelper) GetAuthDetails() *authentication.AuthDetails {
	if r.authDetails != nil {
		return r.authDetails
	}
	r.authDetails = authentication.GetAuthDetailsFromCtx(r.httpCtx)
	return r.authDetails
}

func (r *ResolverHelper) GetGinContext() (*gin.Context, error) {
	if r.httpCtx != nil {
		return r.httpCtx, nil
	}
	httpCtx, _err := GinContextFromContext(r.Ctx)
	if _err != nil {
		return nil, _err
	}
	r.httpCtx = httpCtx
	return httpCtx, nil
}

//

// ========================================================================\\

//

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
