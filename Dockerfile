FROM nginx:1.25.1 as prod-stage
COPY /static/* /usr/share/nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
