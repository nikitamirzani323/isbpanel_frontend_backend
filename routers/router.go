package routers

import (
	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/controllers"
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

	app.Post("/api/movie", controllers.Moviehome)
	app.Post("/api/movienotcdn", controllers.Moviehomenotcdn)
	app.Post("/api/moviebanner", controllers.Moviebanner)
	app.Post("/api/movietrouble", controllers.Movietroublehome)
	app.Post("/api/moviemini", controllers.Movieminihome)
	app.Post("/api/moviesave", controllers.Moviesave)
	app.Post("/api/moviedelete", controllers.Moviedelete)
	app.Post("/api/movieseries", controllers.Moviehomeseries)
	app.Post("/api/movieseriestrouble", controllers.Moviehomeseriestrouble)
	app.Post("/api/movieseriessave", controllers.Movieseriessave)
	app.Post("/api/movieseriesseason", controllers.Seasonhome)
	app.Post("/api/movieseriesseasonsave", controllers.Seasonsave)
	app.Post("/api/movieseriesseasondelete", controllers.Seasondelete)
	app.Post("/api/movieseriesepisode", controllers.Episodehome)
	app.Post("/api/movieseriesepisodesave", controllers.Episodesave)
	app.Post("/api/movieseriesepisodedelete", controllers.Episodedelete)
	app.Post("/api/moviecloudualbum", controllers.Moviecloud)
	app.Post("/api/moviecloudupdate", controllers.Movieupdatecloud)
	app.Post("/api/movieclouddelete", controllers.Moviedeletecloud)
	app.Post("/api/moviecloudupload", controllers.Movieuploadcloud)
	app.Post("/api/genremovie", controllers.Genrehome)
	app.Post("/api/genremoviesave", controllers.Genresave)
	app.Post("/api/genremoviedelete", controllers.Genredelete)

	app.Post("/api/slider", controllers.Sliderhome)
	app.Post("/api/slidersave", controllers.Slidersave)
	app.Post("/api/sliderdelete", controllers.Sliderdelete)

	app.Post("/api/domain", controllers.Domainhome)
	app.Post("/api/domainsave", controllers.DomainSave)

	app.Post("/api/webagen", controllers.Websiteagenhome)
	app.Post("/api/webagensave", controllers.Websiteagensave)
	app.Post("/api/game", controllers.Gamehome)
	app.Post("/api/gamesave", controllers.Gamesave)

	app.Post("/api/cloudflare", controllers.Moviecloud2)
	app.Post("/api/album", controllers.Albumhome)
	app.Post("/api/albumsave", controllers.Albumsave)

	app.Post("/api/crm", controllers.Crmhome)
	app.Post("/api/crmsales", controllers.Crmsaleshome)
	app.Post("/api/crmsalesperform", controllers.Crmsalesperform)
	app.Post("/api/crmdeposit", controllers.Crmdeposithome)
	app.Post("/api/crmisbtv", controllers.Crmisbtvhome)
	app.Post("/api/crmduniafilm", controllers.Crmduniafilm)
	app.Post("/api/crmsave", controllers.CrmSave)
	app.Post("/api/crmsavestatus", controllers.CrmSaveStatus)
	app.Post("/api/crmsalessave", controllers.CrmSalesSave)
	app.Post("/api/crmsalesdelete", controllers.CrmSalesDelete)
	app.Post("/api/crmsavesource", controllers.CrmSavesource)
	app.Post("/api/crmuploaddatabase", controllers.Crmdatabase)
	app.Post("/api/crmsavedatabase", controllers.CrmSaveDatabase)
	app.Post("/api/crmsavemaintenance", controllers.CrmSaveMaintenance)

	app.Post("/api/providerslot", controllers.Providerslothome)
	app.Post("/api/providerslotsave", controllers.ProviderslotSave)
	app.Post("/api/prediksislot", controllers.Gameslothome)
	app.Post("/api/prediksislotsave", controllers.GameslotSave)
	app.Post("/api/prediksislotdelete", controllers.GameslotDelete)
	app.Post("/api/prediksislotgenerator", controllers.GameslotGenerator)

	app.Post("/api/banner", controllers.Bannerhome)
	app.Post("/api/bannersave", controllers.BannerSave)

	app.Post("/api/departement", controllers.Departementhome)
	app.Post("/api/departementsave", controllers.DepartementSave)
	app.Post("/api/employee", controllers.Employeehome)
	app.Post("/api/employeedepart", controllers.Employeebydepart)
	app.Post("/api/employeesave", controllers.EmployeeSave)

	app.Post("/api/event", controllers.Eventhome)
	app.Post("/api/eventdetail", controllers.Eventdetailhome)
	app.Post("/api/eventdetailwinner", controllers.Eventdetailwinner)
	app.Post("/api/eventgroupdetail", controllers.Eventdetailmemberhome)
	app.Post("/api/eventsave", controllers.EventSave)
	app.Post("/api/eventdetailsave", controllers.EventDetailSave)
	app.Post("/api/eventdetailstatussave", controllers.EventDetailSaveStatus)
	app.Post("/api/member", controllers.Memberhome)
	app.Post("/api/memberselect", controllers.MemberSelecthome)
	app.Post("/api/membersave", controllers.MemberSave)

	app.Post("/api/livescore", controllers.Livescorehome)
	return app
}
