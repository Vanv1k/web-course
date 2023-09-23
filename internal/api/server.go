package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"database/sql"

	_ "github.com/lib/pq"
)

var consultations = []Consultation{
	{
		Id:   1,
		Name: "Интернет безопасность",
		Description: "Защитите свои личные данные, финансы и конфиденциальную информацию " +
			"от киберугроз с помощью современных средств и решений. " +
			"Надежная интернет-безопасность - это ключ к спокойствию и уверенности в сети.",
		Price: 9000,
		Image: "/img/security.jpg",
	},
	{
		Id:   2,
		Name: "Основы безопасного кода",
		Description: "Узнайте, как создавать программное обеспечение, которое защищено от взломов и атак. " +
			"Мы поможем вам освоить принципы безопасного программирования и создавать приложения, " +
			"которые стоят на страже данных и конфиденциальности.",
		Price: 10000,
		Image: "/img/code.jpg",
	},
	{
		Id:   3,
		Name: "Защита баз данных",
		Description: "Наши эксперты готовы помочь вам создать броню вокруг ваших ценных данных. " +
			"Мы проведем аудит безопасности, разработаем стратегию защиты и обеспечим " +
			"вашу базу данных надежными решениями. У нас в штате большое количество специалистов " +
			"с соответствующим опытом, которые готовы Вам помочь.",
		Price: 23000,
		Image: "/img/lock.jpg",
	},
	{
		Id:   4,
		Name: "Блокчейн технологии",
		Description: "Внедрение блокчейна поможет вам улучшить " +
			"эффективность бизнес-процессов, обеспечить прозрачность и безопасность данных, " +
			"а также снизить затраты. Мы создадим для вас индивидуальное решение, " +
			"которое поможет вашей компании выйти на новый уровень конкурентоспособности.",
		Price: 12000,
		Image: "/img/crypto.jpg",
	},
}

func setupDB() (*sql.DB, error) {
	// Здесь используйте значения из ваших переменных окружения или файла конфигурации
	connectionString := "user=postgres password=123987 dbname=IT_services host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Попробовать установить соединение
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.Static("/styles", "./resources/styles")
	r.Static("/js", "./resources/js")
	r.Static("/img", "./resources/img")
	r.Static("/hacker", "./resources")
	r.LoadHTMLGlob("templates/*")

	db, err := setupDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %d", err)
	}
	defer db.Close()

	// rows, err := db.Query(`SELECT consultatuion_id, name FROM public.consultation`)
	// if err != nil {
	// 	log.Fatalf("Failed to query the database: %v", err)
	// }
	// defer rows.Close()

	r.GET("/", func(c *gin.Context) {

		searchQuery := c.DefaultQuery("fsearch", "")

		if searchQuery == "" {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"services": consultations,
			})
			return
		}

		var result []Consultation

		for _, consultation := range consultations {
			if strings.Contains(strings.ToLower(consultation.Name), strings.ToLower(searchQuery)) {
				result = append(result, consultation)
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"services":    result,
			"search_text": searchQuery,
		})
	})

	// r.GET("/search", func(c *gin.Context) {

	// 	searchQuery := c.DefaultQuery("fsearch","")

	// 	var result []Service

	// 	for _, service := range services {
	// 		if strings.Contains(strings.ToLower(service.Name), strings.ToLower(searchQuery)) {
	// 			result = append(result, service)
	// 		}
	// 	}

	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H {
	// 		"services": result,
	// 		"search_text": searchQuery,
	// 	})
	// })

	r.GET("/service/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// Обработка ошибки
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		consultation := consultations[id-1]
		c.HTML(http.StatusOK, "card.tmpl", consultation)
	})

	r.Run()

	log.Println("Server down")
}
