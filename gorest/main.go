// main function of the application
package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/controller"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/lib/middleware"

	"github.com/gin-gonic/gin"
)

var configure = config.Config()

func main() {
	if configure.Database.RDBMS.Activate == config.Activated {
		// Initialize RDBMS client
		if err := database.InitDB().Error; err != nil {
			fmt.Println(err)
			return
		}
	}

	if configure.Database.REDIS.Activate == config.Activated {
		// Initialize REDIS client
		if _, err := database.InitRedis(); err != nil {
			fmt.Println(err)
			return
		}
	}

	if configure.Database.MongoDB.Activate == config.Activated {
		// Initialize MONGO client
		if _, err := database.InitMongo(); err != nil {
			fmt.Println(err)
			return
		}
	}

	router, err := SetupRouter()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = router.Run(":" + configure.Server.ServerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SetupRouter ...
func SetupRouter() (*gin.Engine, error) {
	if configure.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode
	}

	// Write log file
	// Console color is not required to write the logs to the file
	//	gin.DisableConsoleColor()

	// Create a log file with start time
	// dt := time.Now()
	// t := dt.Format(time.RFC3339)
	// file, err := os.Create("./logs/start:" + t + ".log")
	// if err != nil {
	//	 return nil, err
	// }
	// gin.DefaultWriter = io.MultiWriter(file)

	// If it is required to write the logs to the file and the console
	// at the same time
	// gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// Creates a router without any middleware by default
	// router := gin.New()

	// Logger middleware: gin.DefaultWriter = os.Stdout
	// router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500
	// if there is one
	// router.Use(gin.Recovery())

	// gin.Default() = gin.New() + gin.Logger() + gin.Recovery()
	router := gin.Default()

	// Which proxy to trust:
	// disable this feature as it still fails
	// to provide the real client IP in
	// different scenarios
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return router, err
	}

	// when using Cloudflare's CDN:
	// router.TrustedPlatform = gin.PlatformCloudflare
	//
	// when running on Google App Engine:
	// router.TrustedPlatform = gin.PlatformGoogleAppEngine
	//
	/*
		when using apache or nginx reverse proxy
		without Cloudflare's CDN or Google App Engine

		config for nginx:
		=================
		proxy_set_header X-Real-IP       $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	*/
	// router.TrustedPlatform = "X-Real-Ip"
	//
	// set TrustedPlatform to get the real client IP
	trustedPlatform := configure.Security.TrustedPlatform
	if trustedPlatform == "cf" {
		trustedPlatform = gin.PlatformCloudflare
	}
	if trustedPlatform == "google" {
		trustedPlatform = gin.PlatformGoogleAppEngine
	}
	router.TrustedPlatform = trustedPlatform

	// CORS
	// if configure.Security.MustCORS == config.Activated {
	// 	router.Use(middleware.CORS(
	// 		configure.Security.CORS.Origin,
	// 		configure.Security.CORS.Credentials,
	// 		configure.Security.CORS.Headers,
	// 		configure.Security.CORS.Methods,
	// 		configure.Security.CORS.MaxAge,
	// 	))
	// }
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	// Sentry.io
	if configure.Logger.Activate == config.Activated {
		router.Use(middleware.SentryCapture(configure.Logger.SentryDsn))
	}

	// WAF
	if configure.Security.MustFW == config.Activated {
		router.Use(middleware.Firewall(
			configure.Security.Firewall.ListType,
			configure.Security.Firewall.IP,
		))
	}

	// Render HTML
	if configure.ViewConfig.Activate == config.Activated {
		router.Use(middleware.Pongo2(configure.ViewConfig.Directory))
	}

	// API:v1.0
	v1 := router.Group("/api/v1/")
	{
		v1.GET("greetings", controller.Greetings)

		// RDBMS
		if configure.Database.RDBMS.Activate == config.Activated {
			// Register - no JWT required
			v1.POST("register", controller.CreateUserAuth)

			// Verify email
			if configure.Security.VerifyEmail {
				if configure.Database.REDIS.Activate == config.Activated {
					v1.POST("verify", controller.VerifyEmail)
					v1.POST("resend-verification-email", controller.CreateVerificationEmail)
				}
			}

			// Login - app issues JWT
			v1.POST("login", controller.Login)

			// test
			rTests := v1.Group("me")
			rTests.Use(middleware.JWT())
			rTests.POST("", controller.GetCurrUserAuth)

			// Refresh - app issues new JWT
			rJWT := v1.Group("refresh")
			rJWT.Use(middleware.RefreshJWT())
			rJWT.POST("", controller.Refresh)

			// Double authentication
			if configure.Security.Must2FA == config.Activated {
				r2FA := v1.Group("2fa")
				r2FA.Use(middleware.JWT())
				r2FA.POST("setup", controller.Setup2FA)
				r2FA.POST("activate", controller.Activate2FA)
				r2FA.POST("validate", controller.Validate2FA)
				if configure.Security.Must2FA == config.Activated {
					r2FA.Use(middleware.TwoFA(
						configure.Security.TwoFA.Status.On,
						configure.Security.TwoFA.Status.Off,
						configure.Security.TwoFA.Status.Verified,
					))
				}
				// disable 2FA
				r2FA.POST("deactivate", controller.Deactivate2FA)
			}

			// Update/reset password
			rPass := v1.Group("password")
			// Reset forgotten password
			if configure.EmailConf.Activate == config.Activated {
				// send password recovery email
				rPass.POST("forgot", controller.PasswordForgot)
				// recover account and set new password
				rPass.POST("reset", controller.PasswordRecover)
			}
			rPass.Use(middleware.JWT())
			if configure.Security.Must2FA == config.Activated {
				rPass.Use(middleware.TwoFA(
					configure.Security.TwoFA.Status.On,
					configure.Security.TwoFA.Status.Off,
					configure.Security.TwoFA.Status.Verified,
				))
			}
			// change password while logged in
			rPass.POST("edit", controller.PasswordUpdate)

			// User
			rUsers := v1.Group("users")
			rUsers.GET("", controller.GetUsers)    // Non-protected
			rUsers.GET("/:id", controller.GetUser) // Non-protected
			rUsers.Use(middleware.JWT())
			if configure.Security.Must2FA == config.Activated {
				rUsers.Use(middleware.TwoFA(
					configure.Security.TwoFA.Status.On,
					configure.Security.TwoFA.Status.Off,
					configure.Security.TwoFA.Status.Verified,
				))
			}
			rUsers.POST("", controller.CreateUser)      // Protected
			rUsers.PUT("", controller.UpdateUser)       // Protected
			rUsers.PUT("/hobbies", controller.AddHobby) // Protected

			// Post
			rPosts := v1.Group("posts")
			rPosts.GET("", controller.GetPosts)    // Non-protected
			rPosts.GET("/:id", controller.GetPost) // Non-protected
			rPosts.Use(middleware.JWT())
			if configure.Security.Must2FA == config.Activated {
				rPosts.Use(middleware.TwoFA(
					configure.Security.TwoFA.Status.On,
					configure.Security.TwoFA.Status.Off,
					configure.Security.TwoFA.Status.Verified,
				))
			}
			rPosts.POST("", controller.CreatePost)       // Protected
			rPosts.PUT("/:id", controller.UpdatePost)    // Protected
			rPosts.DELETE("/:id", controller.DeletePost) // Protected

			// Hobby
			rHobbies := v1.Group("hobbies")
			rHobbies.GET("", controller.GetHobbies) // Non-protected
		}

		// REDIS Playground
		if configure.Database.REDIS.Activate == config.Activated {
			rPlayground := v1.Group("playground")
			rPlayground.GET("/redis_read", controller.RedisRead)        // Non-protected
			rPlayground.POST("/redis_create", controller.RedisCreate)   // Non-protected
			rPlayground.DELETE("/redis_delete", controller.RedisDelete) // Non-protected

			rPlayground.GET("/redis_read_hash", controller.RedisReadHash)        // Non-protected
			rPlayground.POST("/redis_create_hash", controller.RedisCreateHash)   // Non-protected
			rPlayground.DELETE("/redis_delete_hash", controller.RedisDeleteHash) // Non-protected
		}

		// Mongo Playground
		if configure.Database.MongoDB.Activate == config.Activated {
			rPlaygroundMongo := v1.Group("playground-mongo")
			rPlaygroundMongo.POST("/mongo_create_one", controller.MongoCreateOne)                 // Non-protected
			rPlaygroundMongo.GET("/mongo_get_all", controller.MongoGetAll)                        // Non-protected
			rPlaygroundMongo.GET("/mongo_get_by_id/:id", controller.MongoGetByID)                 // Non-protected
			rPlaygroundMongo.POST("/mongo_get_by_filter", controller.MongoGetByFilter)            // Non-protected
			rPlaygroundMongo.PUT("/mongo_update_by_id", controller.MongoUpdateByID)               // Non-protected
			rPlaygroundMongo.DELETE("/mongo_delete_field_by_id", controller.MongoDeleteFieldByID) // Non-protected
			rPlaygroundMongo.DELETE("/mongo_delete_doc_by_id/:id", controller.MongoDeleteByID)    // Non-protected
		}

		// Basic Auth demo
		if configure.Security.MustBasicAuth == config.Activated {
			user := configure.Security.BasicAuth.Username
			pass := configure.Security.BasicAuth.Password
			rBasicAuth := v1.Group("access_resources")
			rBasicAuth.Use(gin.BasicAuth(gin.Accounts{
				user: pass,
			}))
			rBasicAuth.GET("", controller.AccessResource) // Protected
		}
	}

	return router, nil
}
