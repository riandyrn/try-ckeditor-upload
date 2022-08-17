run:
	docker build -t try-ckeditor-upload .
	docker run -p 9765:9765 -e BASE_IMAGE_URL=http://localhost:9765/images try-ckeditor-upload