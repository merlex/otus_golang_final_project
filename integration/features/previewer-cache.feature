# file: features/previewer-cache.feature

# http://localhost:8585/
# http://previewer:8585/

Feature: Тесты превьюера изображений с проверкой работы кеша

	Scenario: Получение превью изображения gopher_50x50.jpg с удаленного сервера
		When I send "GET" request to "http://previewer:8585/fill/50/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_50x50.jpg"
		And Image get from remote server

	Scenario: Получение превью изображения gopher_50x50.jpg из кеша
		When I send "GET" request to "http://previewer:8585/fill/50/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_50x50.jpg"
		And Image get from cache
	
	Scenario: Получение превью изображения gopher_50x50.jpg из кеша
		When I send "GET" request to "http://previewer:8585/fill/50/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_50x50.jpg"
		And Image get from cache

	Scenario: Получение превью изображения gopher_200x700.jpg с удаленного сервера
		When I send "GET" request to "http://previewer:8585/fill/200/700/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 200
		And The response equivalent image "./images/gopher_200x700.jpg"