FROM nginx:1.27-alpine
COPY default.conf /etc/nginx/conf.d/default.conf
COPY index.html /usr/share/nginx/html/index.html
COPY sliver.js /usr/share/nginx/html/sliver.js
COPY vendor/html2canvas.min.js /usr/share/nginx/html/vendor/html2canvas.min.js
COPY vendor/gsap.min.js /usr/share/nginx/html/vendor/gsap.min.js
COPY vendor/ScrollTrigger.min.js /usr/share/nginx/html/vendor/ScrollTrigger.min.js
EXPOSE 80
