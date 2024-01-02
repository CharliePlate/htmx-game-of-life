styles:
	@echo "Compiling styles..."
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/styles.css
	
styles-watch:
	@echo "Watching styles..."
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/styles.css --watch

air:
	air

templ:
	@echo "Compiling templates..."
	templ generate

templ-watch:
	@echo "Watching templates..."
	templ generate --watch

dev:
	${MAKE} -j4 styles-watch air templ-watch
