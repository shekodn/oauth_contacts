FROM scratch

ENV PORT 8000

EXPOSE $PORT

COPY oauth_contacts /

CMD ["/oauth_contacts"]
