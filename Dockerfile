FROM centos:7
COPY . /root/server
EXPOSE 8888
CMD /root/server/cmd/main