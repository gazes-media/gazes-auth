module gazes-auth

go 1.21.4

require (
	github.com/golang-jwt/jwt/v5 v5.2.0 // direct
	github.com/gorilla/mux v1.8.1 // direct
	gorm.io/gorm v1.25.5 // direct
)

require gorm.io/driver/postgres v1.5.4

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.1 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
