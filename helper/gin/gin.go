package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	CORS   *CORSConfig
	Logger *LoggerConfig
}

type CORSConfig struct {
	AllowHeaders     []string
	AllowOrigins     []string
	AllowOriginRegex string
	AllowOriginFunc  func(origin string) bool
}

type LoggerConfig struct {
	SkipPaths []string
}

func New(cfg *Config) *gin.Engine {
	r := gin.New()

	// Register middlewares
	mw := []gin.HandlerFunc{
		GinContext(),
	}
	if cfg.Logger != nil {
		mw = append(mw, Logger(cfg.Logger.SkipPaths...))
	}
	if cfg.CORS != nil {
		corsConfig := cors.DefaultConfig()
		corsConfig.AddAllowHeaders("Accept", "Accept-Language", "Authorization", "Referer", "Vary", "User-Agent")
		corsConfig.AddAllowHeaders(cfg.CORS.AllowHeaders...)
		corsConfig.AllowCredentials = true
		// Note: DO NOT use AllowAllOrigins because it conflicts with AllowCredentials
		if cfg.CORS.AllowOriginFunc != nil {
			corsConfig.AllowOriginFunc = cfg.CORS.AllowOriginFunc
		} else if len(cfg.CORS.AllowOrigins) > 0 {
			corsConfig.AllowOriginFunc = func(origin string) bool {
				return IsOriginAllowed(cfg.CORS.AllowOrigins, origin)
			}
		} else if cfg.CORS.AllowOriginRegex != "" {
			corsConfig.AllowOriginFunc = func(origin string) bool {
				return IsOriginAllowedRegex(cfg.CORS.AllowOriginRegex, origin)
			}
		} else {
			corsConfig.AllowOriginFunc = func(origin string) bool {
				return true
			}
		}
		mw = append(mw, cors.New(corsConfig))
	}
	r.Use(mw...)

	// Register health check API
	r.GET("healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	return r
}

func IsOriginAllowed(allowOrigins []string, origin string) bool {
	if pie.Contains(allowOrigins, "*") {
		return true
	}
	parsedOrigin, err := url.Parse(origin)
	if err != nil {
		return false
	}
	if parsedOrigin.Hostname() == "" {
		return false
	}
	for _, allowed := range allowOrigins {
		p, _ := url.Parse(allowed)
		if p.Hostname() == parsedOrigin.Hostname() && p.Scheme == parsedOrigin.Scheme {
			return true
		} else if strings.Index(p.Hostname(), "*.") == 0 {
			hostname := strings.TrimLeft(p.Hostname(), "*")
			hostname = strings.ReplaceAll(hostname, ".", "\\.")
			pattern := fmt.Sprintf("^%s://[a-zA-Z][a-zA-Z0-9-.]+%s(:[0-9]{1,5})?", p.Scheme, hostname)
			match, _ := regexp.MatchString(pattern, origin)
			if match {
				return true
			}
		}
	}
	return false
}

func IsOriginAllowedRegex(pattern string, origin string) bool {
	match, _ := regexp.MatchString(pattern, origin)
	return match
}
