# file: features/previewer-positive.feature

# http://localhost:8585/
# http://previewer:8585/

Feature: Позитивные тесты превьюера изображений

	Scenario: Доступность сервиса превьюера
		When I send "GET" request to "http://previewer:8585/"
		Then The response code should be 200
		And The response should match text "This is image previewer service!"

	Scenario: Получение превью изображения gopher_50x50.jpg
		When I send "GET" request to "http://previewer:8585/fill/50/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_50x50.jpg"

	Scenario: Получение превью изображения gopher_500x500.jpg 
		When I send "GET" request to "http://previewer:8585/fill/500/500/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_500x500.jpg"

	Scenario: Получение превью изображения gopher_200x700.jpg 
		When I send "GET" request to "http://previewer:8585/fill/200/700/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_200x700.jpg"

	Scenario: Получение превью изображения gopher_256x126.jpg 
		When I send "GET" request to "http://previewer:8585/fill/256/126/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_256x126.jpg"

	Scenario: Получение превью изображения gopher_333x666.jpg 
		When I send "GET" request to "http://previewer:8585/fill/333/666/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_333x666.jpg"

	Scenario: Получение превью изображения _gopher_original_1024x504.jpg
		When I send "GET" request to "http://previewer:8585/fill/1024/252/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_1024x252.jpg"

	Scenario: Получение превью изображения gopher_2000x1000.jpg 
		When I send "GET" request to "http://previewer:8585/fill/2000/1000/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_2000x1000.jpg"