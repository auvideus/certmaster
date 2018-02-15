FROM centos:7

VOLUME /etc/letsencrypt

RUN yum -y install epel-release

RUN yum -y install \
    certbot \
    git \
    golang

RUN go get \
    github.com/digitalocean/godo \
    github.com/digitalocean/godo/context \
    github.com/auvideus/certmaster
    golang.org/x/oauth2

ADD hook/certbot.go /opt/certbot.go
ADD pre.sh /opt/pre.sh
ADD post.sh /opt/post.sh

CMD ["/usr/bin/go", "run", "/opt/certbot.go", "loop"]
#CMD ["/bin/bash"]
