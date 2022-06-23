## Call auth endpoint with curl
`curl -i localhost:8090/auth`

## Call secured endpoint with curl
`curl -i --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3RfdXNlciIsImlzcyI6ImdvbGFuZy1zY2hvb2wifQ.96A0yBYd8Ss0j644CRXmVcRVZRz27unXCjmgfgHInOw" localhost:8090/secured`