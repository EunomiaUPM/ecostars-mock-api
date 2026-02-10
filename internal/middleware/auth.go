package middleware

import (
	"context"
	"ecostars-fake-api/internal/config"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	provider, err := oidc.NewProvider(context.Background(), cfg.KeycloakURL+"/realms/"+cfg.KeycloakRealm)
	if err != nil {
		// In production, you might want to handle this more gracefully,
		// e.g. retrying or failing startup. For now, panic if auth provider is unreachable
		// or just log error and maybe fail the request if it happens during request time?
		// Better to panic on startup or lazy load.
		// Since this is a simple factory, we'll just log and let the verifier fail later?
		// No, NewProvider makes network calls. It's better to do this initialization outside the middleware closure
		// or handle the error properly.
		// For simplicity in this task, we will assume the provider is available or lazy load it?
		// Initializing inside the middleware for every request is bad for performance.
		// But initialization at startup might fail if Keycloak isn't up yet.
		// Let's stick to simple implementation: panic if not reachable at startup, or better:
		// allow the middleware to try connecting.
		// Let's assume we initialize it once.
	}

	var verifier *oidc.IDTokenVerifier
	if provider != nil {
		verifier = provider.Verifier(&oidc.Config{ClientID: cfg.KeycloakClientID})
	}

	return func(c *gin.Context) {
		// If provider setup failed (e.g. Keycloak down during startup), try to re-init?
		// Or just fail.
		if provider == nil {
			var err error
			provider, err = oidc.NewProvider(context.Background(), cfg.KeycloakURL+"/realms/"+cfg.KeycloakRealm)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "Authentication provider unavailable"})
				return
			}
			verifier = provider.Verifier(&oidc.Config{ClientID: cfg.KeycloakClientID})
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		idToken, err := verifier.Verify(context.Background(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		// Extract claims if needed
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to parse claims"})
			return
		}

		c.Set("claims", claims)
		c.Set("user_id", idToken.Subject)

		c.Next()
	}
}
