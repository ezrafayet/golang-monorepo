module service-b

go 1.22

require shared/libs/lib-a v0.0.0

require github.com/go-chi/chi/v5 v5.0.12 // indirect

replace shared/libs/lib-a => ../../libs/golang/lib-a