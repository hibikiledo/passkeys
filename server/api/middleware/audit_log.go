package middleware

import (
	"github.com/labstack/echo/v4"
	auditlog "github.com/teamhanko/passkey-server/audit_log"
	"github.com/teamhanko/passkey-server/persistence"
	"github.com/teamhanko/passkey-server/persistence/models"
)

func AuditLogger(persister persistence.Persister) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			tenant := ctx.Get("tenant").(*models.Tenant)
			auditLogConfig := tenant.Config.AuditLogConfig

			auditLogger := auditlog.NewLogger(persister, auditLogConfig, ctx, tenant)
			ctx.Set("audit_logger", auditLogger)

			return next(ctx)
		}
	}
}
