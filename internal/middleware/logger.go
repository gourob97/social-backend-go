package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Pretty logger with colors and cute formatting
func PrettyLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[90m🕐 ${time_rfc3339_nano}\033[0m " +
			"${status_color}${status}\033[0m " +
			"${method_color}${method}\033[0m " +
			"\033[34m${uri}\033[0m " +
			"${latency_color}${latency_human}\033[0m " +
			"\033[36m⬇️${bytes_in}B\033[0m " +
			"\033[35m⬆️${bytes_out}B\033[0m " +
			"\033[90m👤${remote_ip}\033[0m\n",
	})
}

// Enhanced pretty logger with custom formatting
func EnhancedLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Call next handler
			err := next(c)

			// Calculate response time
			latency := time.Since(start)
			status := c.Response().Status
			method := c.Request().Method
			uri := c.Request().RequestURI
			ip := c.RealIP()
			bytesIn := c.Request().ContentLength
			bytesOut := c.Response().Size

			// Status color
			var statusColor string
			switch {
			case status >= 200 && status < 300:
				statusColor = "\033[32m" // Green
			case status >= 300 && status < 400:
				statusColor = "\033[33m" // Yellow
			case status >= 400 && status < 500:
				statusColor = "\033[91m" // Light Red
			default:
				statusColor = "\033[31m" // Red
			}

			// Method color
			var methodColor string
			switch method {
			case "GET":
				methodColor = "\033[36m" // Cyan
			case "POST":
				methodColor = "\033[35m" // Magenta
			case "PUT":
				methodColor = "\033[33m" // Yellow
			case "DELETE":
				methodColor = "\033[31m" // Red
			default:
				methodColor = "\033[37m" // White
			}

			// Latency indicator
			var latencyIcon string
			switch {
			case latency < 100*time.Millisecond:
				latencyIcon = "\033[32m⚡" // Green lightning
			case latency < 500*time.Millisecond:
				latencyIcon = "\033[33m⏱️" // Yellow clock
			default:
				latencyIcon = "\033[31m🐌" // Red snail
			}

			// Format and print log
			fmt.Printf("🕐 \033[90m%s\033[0m %s%d\033[0m %s%s\033[0m \033[34m%s\033[0m %s%v\033[0m \033[36m⬇️%dB\033[0m \033[35m⬆️%dB\033[0m \033[90m👤%s\033[0m\n",
				start.Format("15:04:05"),
				statusColor, status,
				methodColor, method,
				uri,
				latencyIcon, latency.Truncate(time.Millisecond),
				bytesIn,
				bytesOut,
				ip,
			)

			return err
		}
	}
}

// Startup banner for cute server start
func PrintStartupBanner(port string) {
	fmt.Println("\033[35m") // Magenta
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                   🚀 Social Backend Go                   ║")
	fmt.Println("║                     Ready to Rock! 🎸                   ║")
	fmt.Println("╠══════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Server: \033[36mhttp://localhost:%-8s\033[35m                 ║\n", port)
	fmt.Println("║  Status: \033[32m✅ Running\033[35m                              ║")
	fmt.Printf("║  Time:   \033[33m%s\033[35m                    ║\n", time.Now().Format("15:04:05 02-Jan-2006"))
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Print("\033[0m") // Reset color
	fmt.Println()
	fmt.Println("💡 \033[36mEndpoints:\033[0m")
	fmt.Println("   🏥 GET  /health       - Health check")
	fmt.Println("   📝 POST /api/register - Create new user")
	fmt.Println("   🔐 POST /api/login    - User login")
	fmt.Println()
	fmt.Println("🌟 \033[33mHappy coding!\033[0m")
	fmt.Println("🎯 \033[32mCurl example: curl http://localhost:" + port + "/health\033[0m")
	fmt.Println()
}
