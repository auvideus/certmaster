FROM centos:7

VOLUME /etc/letsencrypt

RUN yum -y install epel-release

RUN yum -y install \
    certbot \
    git \
    golang

RUN go get \
    github.com/auvideus/certmaster/cmd/certmaster \
    github.com/auvideus/certmaster/cmd/authhook \
    github.com/auvideus/certmaster/cmd/cleanuphook

ADD pre.sh /opt/pre.sh
ADD post.sh /opt/post.sh

ENV PATH /root/go/bin:$PATH

CMD ["certmaster"]
