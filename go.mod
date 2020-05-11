module github.com/Ilyes512/satis-go

go 1.14

require (
	github.com/Ilyes512/satis-go/satis v0.0.0-00010101000000-000000000000
	github.com/Ilyes512/satis-go/satis/client v0.0.0-00010101000000-000000000000 // indirect
	github.com/Ilyes512/satis-go/satis/satisphp v0.0.0-00010101000000-000000000000 // indirect
	github.com/Ilyes512/satis-go/satis/satisphp/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/Ilyes512/satis-go/satis/satisphp/db v0.0.0-00010101000000-000000000000 // indirect
	github.com/Ilyes512/satis-go/satis/satisphp/job v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace (
	github.com/Ilyes512/satis-go/satis => ./satis
	github.com/Ilyes512/satis-go/satis/client => ./satis/client
	github.com/Ilyes512/satis-go/satis/satisphp => ./satis/satisphp
	github.com/Ilyes512/satis-go/satis/satisphp/api => ./satis/satisphp/api
	github.com/Ilyes512/satis-go/satis/satisphp/db => ./satis/satisphp/db
	github.com/Ilyes512/satis-go/satis/satisphp/job => ./satis/satisphp/job
)
