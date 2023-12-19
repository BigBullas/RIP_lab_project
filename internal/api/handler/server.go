// package api

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type CardLaunchVehicle struct {
// 	Id                  int32
// 	ImgURL              string
// 	Title               string
// 	Description         string
// 	LoadCapacity        float64
// 	FlightDate          time.Time
// 	Price               float64
// 	DetailedDescription string
// }

// func StartServer() {
// 	cards := []CardLaunchVehicle{
		// {
		// 	Id:           1,
		// 	ImgURL:       "https://ntv-static.cdnvideo.ru/home/news/2023/20230205/sputn_io.jpg",
		// 	Title:        "«Электро-Л» № 4",
		// 	Description:  "Гидрометеорологическй космический аппарат",
		// 	LoadCapacity: 1.8,
		// 	FlightDate:   time.Date(2023, time.February, 5, 12, 12, 52, 0, time.UTC),
		// 	Price:        16.5,
		// 	DetailedDescription: "Спутники «Электро-Л» создаются в рамках Федеральной космической программы России и входят" +
		// 		" в геостационарную гидрометеорологическую космическую систему «Электро» разработки НПО Лавочкина. Они предназначены" +
		// 		" для обеспечения оперативной и независимой гидрометеорологической информацией подразделений Федеральной службы по" +
		// 		" гидрометеорологии и мониторингу окружающей среды (Росгидромет) и других ведомств. " +
		// 		"Сейчас в системе «Электро», функционирующей на околоземной орбите с 2011 года, используются по целевому назначению " +
		// 		"два спутника — «Электро-Л» № 2 (запущен 11 декабря 2015 года) в точке стояния 14,5° западной долготы и «Электро-Л» №" +
		// 		" 3 (запущен 24 декабря 2019 года) в точке стояния 76° восточной долготы. Аппарату «Электро-Л» № 4 предстоит работать" +
		// 		" в точке стояния 165,8° восточной долготы. \n" +
		// 		"Уникальные особенности спутников «Электро-Л» позволяют получать независимые метеоданные с орбиты Земли каждые 15—30" +
		// 		" минут. Благодаря круглосуточной передаче с космических аппаратов высококачественных многоспектральных снимков повышается" +
		// 		" не только качество и оперативность прогнозов погоды, но и решаются глобальные вопросы мониторинга климата и изменений," +
		// 		" выдаются штормовые и экстренные телеграммы при выявлении чрезвычайных ситуаций. \n" +
		// 		"Также, спутники ретранслируют сигналы от аварийных радиобуев международной спутниковой поисково-спасательной системы" +
		// 		" КОСПАС-САРСАТ. Это помогает поисково-спасательным службам эффективнее реагировать на сигналы бедствия для спасения человеческих жизней.",
		// },
		// {
		// 	Id:           2,
		// 	ImgURL:       "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b2/ISS-59_Progress_MS-11_approaches_the_ISS.jpg/1200px-ISS-59_Progress_MS-11_approaches_the_ISS.jpg",
		// 	Title:        "«Прогресс МС-22»",
		// 	Description:  "Грузовой корабль",
		// 	LoadCapacity: 2.5,
		// 	FlightDate:   time.Date(2023, time.February, 9, 9, 15, 0, 0, time.UTC),
		// 	Price:        26.5,
		// 	DetailedDescription: "«Прогресс МС-22» — космический транспортный грузовой корабль серии «Прогресс», запуск которого состоялся 6" +
		// 		" февраля 2023 года со стартового комплекса «Восток» (Площадка 31) космодрома «Байконур» по программе 83-й миссии снабжения " +
		// 		"Международной космической станции[1]. Это 175-й полёт космического корабля серии «Прогресс».",
		// },
		// {
		// 	Id:           3,
		// 	ImgURL:       "https://news.cgtn.com/news/2023-02-26/Russia-s-replacement-Soyuz-spacecraft-arrives-at-space-station-1hJfEkcHidG/img/b5c3147c02e44da8941bf851fdebfcc7/b5c3147c02e44da8941bf851fdebfcc7-1280.png",
		// 	Title:        "«Союз МС-23»",
		// 	Description:  "Беспилотный корабль",
		// 	LoadCapacity: 7.2,
		// 	FlightDate:   time.Date(2023, time.February, 24, 3, 24, 0, 0, time.UTC),
		// 	Price:        35.5,
		// 	DetailedDescription: "'Союз МС' ('МС' - 'модернизированные системы') принадлежит к семейству космических" +
		// 		" кораблей 'Союз' (первый пилотируемый запуск состоялся в 1967 году). Головным разработчиком и изготовителем" +
		// 		" корабля является Ракетно-космическая корпорация 'Энергия' им. С. П. Королева (РКК 'Энергия'; город Королев," +
		// 		" Московская область). Эскизный проект 'Союза МС', разработанный по заданию Федерального космического агентства" +
		// 		" (ныне госкорпорация 'Роскосмос'), был одобрен на заседании научно-технического совета РКК 'Энергия' в августе" +
		// 		" 2011 года. Корабль создан на базе предыдущей модификации 'Союз ТМА-М' (запуски проводились в 2010-2016 годах)" +
		// 		" путем глубокой модернизации. 'Союз МС' предназначен для доставки экипажей на МКС и возвращения их с орбиты" +
		// 		" обратно на Землю. Он выполняет роль корабля-спасателя в случаях вынужденной или аварийной эвакуации экипажа" +
		// 		" (при возникновении опасной ситуации на станции, заболевания или травмы космонавтов). Кроме того, 'Союз МС'" +
		// 		" используется для доставки на МКС и возвращения с орбиты небольших грузов (научно-исследовательской аппаратуры" +
		// 		" и результатов экспериментов, личных вещей космонавтов и др.), а также для удаления со станции отходов в бытовом" +
		// 		" отсеке, который сгорает плотных слоях атмосферы при спуске корабля.",
		// },
		// {
		// 	Id:           4,
		// 	ImgURL:       "https://finobzor.ru/uploads/posts/2016-09/org_vrke626.jpg",
		// 	Title:        "«Луч-5Х»",
		// 	Description:  "Многофункциональная космическая система ретрансляции",
		// 	LoadCapacity: 37,
		// 	FlightDate:   time.Date(2023, time.March, 13, 2, 13, 0, 0, time.UTC),
		// 	Price:        65,
		// 	DetailedDescription: "'Олимп-К', также обозначаемый как 'Луч', является российским геостационарным спутником," +
		// 		" созданным для Министерства обороны России и российского разведывательного управления ФСБ. Цели миссий не" +
		// 		" опубликованы. Согласно сообщению 'Коммерсанта', спутник будет выполнять двойную роль: одна из них - радиотехническая" +
		// 		" разведка (SIGINT), а другая обеспечивает защищенную связь для правительственных нужд. Обозначение 'Луч' указывает" +
		// 		" на роль ретранслятора данных. Следовательно, обозначение 'Олимп-К' может относиться к полезной нагрузке" +
		// 		" SIGINT, в то время как обозначение 'Луч' относится к полезной нагрузке ретрансляции данных. Другой источник" +
		// 		" сообщает, что спутник должен обеспечивать сигналы навигационной коррекции для системы ГЛОНАСС. Сообщалось также" +
		// 		" о бортовом лазерном устройстве связи.",
		// },
		// {
		// 	Id:           5,
		// 	ImgURL:       "https://avatars.dzeninfra.ru/get-zen_doc/9428044/pub_641e3138e540d5493c71189b_641e520f617db875869202c0/scale_1200",
		// 	Title:        "«Барс-М №4»",
		// 	Description:  "Электронно-оптический спутник",
		// 	LoadCapacity: 4,
		// 	FlightDate:   time.Date(2023, time.March, 23, 9, 40, 0, 0, time.UTC),
		// 	Price:        17.4,
		// 	DetailedDescription: "Спутник 'Барс-М' - это новый электронно-оптический спутник наблюдения за территорией, который" +
		// 		" заменит серию 'Янтарь-1KFT' (Kometa) с возвратом пленки и отмененную серию 'Барс'." +
		// 		" 'Барс-М' является вторым воплощением проекта 'Барс', который был начат в середине 1990-х годов для разработки преемника спутников" +
		// 		" наблюдения за территорией класса Komtea. Первоначальный проект Bars был остановлен в начале 2000-х годов. В" +
		// 		" 2007 году 'ЦСКБ-Прогресс' заключило контракт на поставку 'Барс-М', для которого, как сообщается, сервисный" +
		// 		" модуль на базе 'Янтаря' был заменен новым усовершенствованным сервисным модулем." +
		// 		" Спутники 'Барс-М' оснащены электронно-оптической фотосистемой 'Карат', разработанной и изготовленной Ленинградским" +
		// 		" оптико-механическим объединением (ЛОМО), и двойным лазерным высотомером для получения топографических изображений," +
		// 		" стереоизображений, данных высотомера и изображений высокого разрешения с наземным разрешением около 1 метра.",
		// },
// 	}
// 	log.Println("Server start up")

// 	r := gin.Default()
// 	// в будущем убрать
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	r.LoadHTMLGlob("templates/*")

// 	r.GET("/home", func(c *gin.Context) {
// 		queryString := c.Request.URL.Query() // queryString - это тип url.Values, который содержит все query параметры

// 		strSearch := queryString.Get("spacecraft") // Получение значения конкретного параметра по его имени
// 		if strSearch == "" {
// 			c.HTML(http.StatusOK, "main_page.gohtml", gin.H{
// 				"cards":      cards,
// 				"spacecraft": strSearch,
// 			})
// 		} else {
// 			var selectedCards []CardLaunchVehicle
// 			flag := false
// 			for i := 0; i < len(cards); i++ {
// 				if strings.Contains(strings.ToLower(cards[i].Title), strings.ToLower(strSearch)) {
// 					flag = true
// 					selectedCards = append(selectedCards, cards[i])
// 				}
// 			}

// 			if flag {
// 				c.HTML(http.StatusOK, "main_page.gohtml", gin.H{
// 					"cards":      selectedCards,
// 					"spacecraft": strSearch,
// 				})
// 			} else {
// 				c.HTML(http.StatusOK, "main_page.gohtml", gin.H{
// 					"cards":      "",
// 					"spacecraft": strSearch,
// 				})
// 			}
// 		}

// 	})

// 	r.GET("/card", func(c *gin.Context) {
// 		queryString := c.Request.URL.Query() // queryString - это тип url.Values, который содержит все query параметры

// 		strCardId := queryString.Get("cardId") // Получение значения конкретного параметра по его имени
// 		cardId, err := strconv.Atoi(strCardId)
// 		if err != nil {
// 			fmt.Println("Ошибка при преобразовании строки в число:", err)
// 			return
// 		}
// 		c.HTML(http.StatusOK, "spacecraft_card.gohtml", gin.H{
// 			"card": cards[cardId-1],
// 		})
// 	})

// 	r.Static("/style", "./style")
// 	r.Static("/image", "./resources") // не нужно

// 	err := r.Run()
// 	if err != nil {
// 		log.Println("Server error!!!")
// 		return
// 	}

// 	log.Println("Server down")
// }

package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"RIP_lab1/internal/api"
)

type Handler struct {
	repo api.Repo
}

func NewHandler(repo api.Repo) *Handler {
	return &Handler{repo: repo}
}	

func (h *Handler) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string

		if id != "" {
			log.Printf("id recived %s\n", id)
			intID, err := strconv.Atoi(id) // пытаемся привести это к числу
			if err != nil {                // если не получилось
				log.Printf("cant convert id %v", err)
				c.Error(err)
				return
			}
			// получаем данные по товару
			product, err := h.repo.GetProductByID(uint(intID))
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "try with id",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"title": "Main website",
			"test":  []string{"a", "b"},
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

// func (h *Handler) StartServer() {
// 	log.Println("Server start up")

// 	r := gin.Default()
// 	r.GET("/ping", h.Ping)

// 	// loads all html in templates dir
// 	r.LoadHTMLGlob("templates/*")

// 	r.GET("/home", h.GetStarList)
// 	r.GET("/star/:id", h.GetStarById)
// 	r.POST("/star/:id", h.DeleteStarById)

// 	r.Static("/image", "./resources")
// 	r.Static("/styles", "./styles")

// 	// listen and serve on 127.0.0.1:8080
// 	err := r.Run()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func (h *Handler) Ping(c *gin.Context) {
// 	c.JSON(
// 		http.StatusOK,
// 		gin.H{
// 			"message": "pong",
// 		})
// }

// func (h *Handler) GetStarList(c *gin.Context) {
// 	files := []string{
// 		"templates/list.gohtml",
// 	}

// 	data, err := h.repo.GetStarsByNameFilter(c.Query("starName"))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	render.RenderTmpl(files, gin.H{
// 		"Items":      data,
// 		"QueryParam": c.Query("starName"),
// 	}, c)
// }

// func (h *Handler) GetStarById(c *gin.Context) {
// 	files := []string{
// 		"templates/item.gohtml",
// 	}

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	item, err := h.repo.GetStarByID(id)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	render.RenderTmpl(files, item, c)
// }

// func (h *Handler) DeleteStarById(c *gin.Context) {
// 	cardId := c.Param("id")
// 	id, err := strconv.Atoi(cardId)
// 	if err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 	}

// 	err = h.repo.DeleteStarById(id)
// 	if err != nil { // если не получилось
// 		log.Printf("cant get star by id %v", err)
// 		c.Error(err)
// 		return
// 	}
// 	c.Redirect(http.StatusFound, "/home")
// }