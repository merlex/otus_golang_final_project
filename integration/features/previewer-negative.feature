# file: features/previewer-negative.feature

# http://localhost:8585/
# http://previewer:8585/

Feature: Негативные тесты превьюера изображений

	Scenario: Получение превью изображения для несуществующего файла 
		When I send "GET" request to "http://previewer:8585/fill/50/50/nginx/1.jpg"
		Then The response code should be 404
	
	Scenario: Получение превью изображения для файла не являющегося изображением 
		When I send "GET" request to "http://previewer:8585/fill/50/50/web/index.html"
		Then The response code should be 400
		And The response should match text "image must have an extension either jpg or jpeg"

	Scenario: Получение превью изображения от несуществующего сервера 
		When I send "GET" request to "http://previewer:8585/fill/50/50/myimageserver/images/_gopher_original_1024x504.jpg"
		Then The response code should be 404
		And The response should contains text
		"""
Get "http://myimageserver/images/_gopher_original_1024x504.jpg?": dial tcp: lookup myimageserver on 127.0.0.11:53: no such host
		"""

	Scenario: Получение превью изображения c неправильно заданной шириной 
		When I send "GET" request to "http://previewer:8585/fill/abc/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 400
		And The response should match text "path variable width has wrong value"

	Scenario: Получение превью изображения c неправильно заданной высотой 
		When I send "GET" request to "http://previewer:8585/fill/50/abc/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 400
		And The response should match text "path variable height has wrong value"
	
	Scenario: Получение превью изображения c неправильно количеством аргументов (нет пути до изображения) 
		When I send "GET" request to "http://previewer:8585/fill/50/50"
		Then The response code should be 400
		And The response should match text "not set width, height or image path in URL"

	Scenario: Получение превью изображения c неправильно количеством аргументов (нет высоты или ширины) 
		When I send "GET" request to "http://previewer:8585/fill/50"
		Then The response code should be 400
		And The response should match text "not set width, height or image path in URL"

	Scenario: Получение превью изображения c неправильно количеством аргументов (нет параметров) 
		When I send "GET" request to "http://previewer:8585/fill"
		Then The response code should be 404
		And The response should match text "404 page not found"
	
	Scenario: Получение превью изображения c неправильно количеством аргументов (слишком много параметров) 
		When I send "GET" request to "http://previewer:8585/fill/50/50/50/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 404
		And The response should contains text
		"""
Get "http://50/nginx/_gopher_original_1024x504.jpg?": dial tcp: lookup 50: no such host
		"""
	
	Scenario: Получение превью изображения со слишком большой шириной (width > 3840 || height > 2160) 
		When I send "GET" request to "http://previewer:8585/fill/4000/2160/nginx/_gopher_original_1024x504.jpg"
		Then The response code should be 400
		And The response should match text "width or height is very large"