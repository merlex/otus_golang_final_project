FROM nginx:alpine
COPY integration/images/ /usr/share/nginx/html
COPY build/nginx/conf/ /etc/nginx/conf.d/
COPY build/nginx/401/  /usr/share/nginx/html/
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]