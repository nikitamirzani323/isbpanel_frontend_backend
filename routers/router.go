package routers

import (
	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/controllers"
	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", controllers.Home)
	app.Post("/api/alladmin", controllers.Adminhome)
	app.Post("/api/detailadmin", controllers.AdminDetail)
	app.Post("/api/saveadmin", controllers.AdminSave)

	app.Post("/api/alladminrule", controllers.Adminrulehome)
	app.Post("/api/saveadminrule", controllers.AdminruleSave)

	app.Post("/api/pasaran", controllers.Pasaranhome)
	app.Post("/api/pasaransave", controllers.Pasaransave)
	app.Post("/api/keluaran", controllers.Keluaranhome)
	app.Post("/api/keluaransave", controllers.Keluaransave)
	app.Post("/api/keluarandelete", controllers.Keluarandelete)

	app.Post("/api/prediksi", controllers.Prediksihome)
	app.Post("/api/prediksisave", controllers.Prediksisave)
	app.Post("/api/prediksidelete", controllers.Prediksidelete)

	app.Post("/api/tafsirmimpi", controllers.Tafsirmimpihome)
	app.Post("/api/tafsirmimpisave", controllers.Tafsirmimpisave)

	app.Post("/api/news", controllers.Newshome)
	app.Post("/api/newssave", controllers.Newssave)
	app.Post("/api/newsdelete", controllers.Newsdelete)
	app.Post("/api/categorynews", controllers.Categoryhome)
	app.Post("/api/categorynewssave", controllers.Categorysave)
	app.Post("/api/categorynewsdelete", controllers.Categorydelete)

	app.Post("/api/movie", middleware.JWTProtected(), controllers.Moviehome)
	app.Post("/api/movienotcdn", middleware.JWTProtected(), controllers.Moviehomenotcdn)
	app.Post("/api/movietrouble", middleware.JWTProtected(), controllers.Movietroublehome)
	app.Post("/api/moviemini", middleware.JWTProtected(), controllers.Movieminihome)
	app.Post("/api/moviesave", middleware.JWTProtected(), controllers.Moviesave)
	app.Post("/api/moviedelete", middleware.JWTProtected(), controllers.Moviedelete)
	app.Post("/api/movieseries", middleware.JWTProtected(), controllers.Moviehomeseries)
	app.Post("/api/movieseriestrouble", middleware.JWTProtected(), controllers.Moviehomeseriestrouble)
	app.Post("/api/movieseriessave", middleware.JWTProtected(), controllers.Movieseriessave)
	app.Post("/api/movieseriesseason", middleware.JWTProtected(), controllers.Seasonhome)
	app.Post("/api/movieseriesseasonsave", middleware.JWTProtected(), controllers.Seasonsave)
	app.Post("/api/movieseriesseasondelete", middleware.JWTProtected(), controllers.Seasondelete)
	app.Post("/api/movieseriesepisode", middleware.JWTProtected(), controllers.Episodehome)
	app.Post("/api/movieseriesepisodesave", middleware.JWTProtected(), controllers.Episodesave)
	app.Post("/api/movieseriesepisodedelete", middleware.JWTProtected(), controllers.Episodedelete)
	app.Post("/api/moviecloudualbum", middleware.JWTProtected(), controllers.Moviecloud)
	app.Post("/api/moviecloudupdate", middleware.JWTProtected(), controllers.Movieupdatecloud)
	app.Post("/api/movieclouddelete", middleware.JWTProtected(), controllers.Moviedeletecloud)
	app.Post("/api/moviecloudupload", middleware.JWTProtected(), controllers.Movieuploadcloud)
	app.Post("/api/genremovie", middleware.JWTProtected(), controllers.Genrehome)
	app.Post("/api/genremoviesave", middleware.JWTProtected(), controllers.Genresave)
	app.Post("/api/genremoviedelete", middleware.JWTProtected(), controllers.Genredelete)

	app.Post("/api/slider", middleware.JWTProtected(), controllers.Sliderhome)
	app.Post("/api/slidersave", middleware.JWTProtected(), controllers.Slidersave)
	app.Post("/api/sliderdelete", middleware.JWTProtected(), controllers.Sliderdelete)

	app.Post("/api/domain", controllers.Domainhome)
	app.Post("/api/domainsave", controllers.DomainSave)

	app.Post("/api/webagen", middleware.JWTProtected(), controllers.Websiteagenhome)
	app.Post("/api/webagensave", middleware.JWTProtected(), controllers.Websiteagensave)
	app.Post("/api/game", controllers.Gamehome)
	app.Post("/api/gamesave", middleware.JWTProtected(), controllers.Gamesave)

	app.Post("/api/cloudflare", middleware.JWTProtected(), controllers.Moviecloud2)
	app.Post("/api/album", middleware.JWTProtected(), controllers.Albumhome)
	app.Post("/api/albumsave", middleware.JWTProtected(), controllers.Albumsave)

	app.Post("/api/crm", controllers.Crmhome)
	app.Post("/api/crmisbtv", controllers.Crmisbtvhome)
	app.Post("/api/crmduniafilm", controllers.Crmduniafilm)
	app.Post("/api/crmsave", controllers.CrmSave)
	app.Post("/api/crmsavesource", controllers.CrmSavesource)
	return app
}
