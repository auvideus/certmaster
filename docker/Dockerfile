FROM centos:7

RUN yum -y install epel-release

RUN yum -y install \
    certbot \
    git \
    golang

RUN go get \
    github.com/auvideus/certmaster/cmd/certmaster \
    github.com/auvideus/certmaster/cmd/authhook \
    github.com/auvideus/certmaster/cmd/cleanuphook

VOLUME /etc/letsencrypt

WORKDIR /opt/certmaster

ADD pre.sh .
ADD post.sh .

ENV PATH /root/go/bin:$PATH

CMD ["certmaster"]
