docker build -t wombat . && \
  docker stop wombat && \
  docker rm wombat && \
  docker run -d --name=wombat --restart=always --mount source=certs,target=/certs -p 80:8080 -p 443:8081 wombat