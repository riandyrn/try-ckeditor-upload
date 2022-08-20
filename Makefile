run:
	docker build -t try-ckeditor-upload .
	docker run \
		-e AWS_ACCESS_KEY_ID=${GHAZLABS_AWS_ACCESS_KEY_ID} \
		-e AWS_SECRET_ACCESS_KEY=${GHAZLABS_AWS_SECRET_ACCESS_KEY} \
		-e AWS_REGION=${GHAZLABS_AWS_REGION} \
		-e BUCKET_NAME=${GHAZLABS_ALCORE_BUCKET_NAME} \
		-e UPLOAD_ACCESS_KEY=91afb696-2ec2-4ff3-b380-ffa7e3dbbaec \
		-p 9765:9765 \
		try-ckeditor-upload