# Run jobs server
run: jobs

install:
	@pip install -r requirements.txt

jobs:
	@export $(grep -v '^#' .env | xargs)
	@python3 jobs/main.py