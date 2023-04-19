docker build -t wombat . && \
  docker-compose stop && \
  docker-compose up -d --remove-orphans