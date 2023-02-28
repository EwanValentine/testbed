new-service:
	cookiecutter \
		--output-dir=services/ \
		--config-file=templates/grpc-service/cookiecutter.json \
		templates/grpc-service/
