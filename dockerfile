FROM ubantu:15.10

MAINTAINER      tiger "lianchonghai@qq.com"

ADD /home/tiger/workspace/go/src/library/library /usr/app/library/library

EXPOSE 8086

VOLUME /home/tiger/workspace/go/src/library /usr/app/library

CMD /usr/app/library/library


