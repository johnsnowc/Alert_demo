FROM centos:7
COPY main /root/server
EXPOSE 8888
CMD /root/server