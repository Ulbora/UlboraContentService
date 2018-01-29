FROM ubuntu

#RUN sudo apt-get update
RUN apt-get update  
ADD main /main
ADD entrypoint.sh /entrypoint.sh
WORKDIR /

EXPOSE 3008
ENTRYPOINT ["/entrypoint.sh"]

