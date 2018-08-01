# personalweb
My own personal career website. Made with polymer with a golang backend.

I will have a Go backend to accept HTTP requests from the website. This will be home for file serving, sending requests to the back-banckend server that will handle heavy computations or other things, and statistics about views and uses. It also has to capable of handling mutliple requests or connections, especially for asyncronous data requests from the front end.

# Start webserver

Save code to GoLang src directory.

```bash
cp -a -r ./gosrc/* $HOME/go/src/
```
Build the go code.

```bash
cd $HOME/go/src/personal-webserver
go build
```

Finally, run the server with 2 optional flags, port and dir. Port defines the port the webserver will look at for requests. Dir is the flag that defines where the static files are placed.

```bash
./personal-webserver --port=8080 --dir=$HOME/static
```

