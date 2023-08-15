module github.com/gksingh1910/go-lang-interface

replace github.com/gksingh1910/go-lang-interface/config => ../config

replace github.com/gksingh1910/go-lang-interface/handler => ../handler

replace github.com/gksingh1910/go-lang-interface/util => ../util

go 1.16

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/joho/godotenv v1.5.1
)
