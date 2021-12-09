FROM nginx:1.21
COPY / /opt/blog
RUN wget https://github.com/gohugoio/hugo/releases/download/v0.90.0/hugo_0.90.0_Linux-64bit.tar.gz \
&& tar -zxvf hugo_0.90.0_Linux-64bit.tar.gz \
&& cp hugo /usr/local/bin/ \
&& cd /opt/blog \
&& hugo server -e production -t LoveIt -D
