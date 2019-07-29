FROM scratch

ADD build/c4-active /c4-active

ENTRYPOINT ["./c4-active"]