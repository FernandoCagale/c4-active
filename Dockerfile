FROM scratch

ADD build/c4-active /c4-active

EXPOSE 8080

ENTRYPOINT ["./c4-active"]