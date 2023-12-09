module main

go 1.21.5

replace github.com/t-s-w/advent-of-code-2023/utils => ../utils

replace github.com/t-s-w/advent-of-code-2023/d05/maps => ./maps

require github.com/t-s-w/advent-of-code-2023/utils v0.0.0-20231209070949-8cb09bc822e6

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/t-s-w/advent-of-code-2023/d05/maps v0.0.0-20231209074932-ccfcd8802680 // indirect
)
