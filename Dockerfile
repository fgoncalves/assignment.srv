FROM alpine
ADD assignment-srv /assignment-srv
ENTRYPOINT [ "/assignment-srv" ]
