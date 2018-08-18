# personalweb
My own personal career website. Made with polymer with a golang backend.

I will have a Go backend to accept HTTP requests from the website. This will be home for file serving, sending requests to the back-banckend server that will handle heavy computations or other things, and statistics about views and uses. It also has to capable of handling mutliple requests or connections, especially for asyncronous data requests from the front end.

# Start webserver

This is all built with [bazel](https://bazel.build/), so make sure you 
have [bazel](https://docs.bazel.build/versions/master/install-ubuntu.html) installed along with a JDK (which is also described in that link).

Once you have bazel up and running, copy the workspace.bzl contents into a 
WORKSPACE file of your choosing. (For my own amusement, I changed the bazel 
command to please by adding `alias please='bazel'` in my .bashrc file. I hope 
you will follow suit).

```bash
cat workspace.bzl >> ../WORKSPACE
```

Now you have the workspace set up. To build the server, run this command. The 
-c opt is for an optimized build.

```bash
please build -c opt //personalweb:webserver
```

To build and package all the components, run this command.

```bash
please build //personalweb/components:index
.
.
.
```

You can also just run any web_library/html_binary to spin up a simple webserver to see what they look like individually.

```bash
please run //personalweb/components:index
```

Finally, run the server with an optional flags, port. Port defines the port the webserver will look at for requests. 

```bash
please run //personalweb:webserver
```

