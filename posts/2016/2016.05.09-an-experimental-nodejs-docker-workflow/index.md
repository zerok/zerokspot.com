# An experimental NodeJS+Docker workflow

Recently, I've been trying hard to integrate Docker into my development
workflow. On a server where I deploy tiny tools internally I've decided that I
don't want to have to manage 10 different environments but simply deploy
cronjob-scripts, services and more as Docker images.

So problems I wanted to address are:

- I want to have a single Dockerfile for use during development as well as
  production. This means that sharing the `node_modules` folder through volumes
  has to be avoided (this will soooo break on OSX 😉).
- During development (esp. for small service with a web-interface) I want to be
  able to have some kind of file-monitoring going on so that I don't have to
  restart the container every time I change something.


-----------------

The result of this experiment has been this little Dockerfile:

```
FROM node:4

# Prepare non-root user and folders
RUN useradd --system --user-group --create-home app && \
    mkdir /app && chown app:app /app
RUN npm install -g nodemon

# Install dependency outside of the app volume
COPY package.json /opt/
RUN cd /opt && npm install
ENV NODE_PATH=/opt/node_modules

VOLUME ["/app"]
USER app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]
```

This will install all the dependencies the project has inside a folder outside
the primary working directory and therefore avoid linking issues esp. with npm3.

The `entrypoint.sh` is simply whatever I need to use this in production. To make
switching between dev and prod easier, I try to keep all the dev-configuration
inside the `docker-compose.yml` file which mostly just mounts the project's
folder into `/app` and overrides the entrypoint (or for different setups simply
sets command parameters):

```
version: '2'
services:
  app:
    build: .
    volumes:
      - .:/app
    ports:
      - 8080:8080
    entrypoint: nodemon -L main.js
```

For more complex systems it's trivial to add dependencies like PostgreSQL or
Redis here.

The big downside so far is the `-L` flag you can see in the entrypoint
specification. Without it nodemon won't notice any changes within the volume but
with it it's doing quite a bit of busy-work which wastes power. I haven't found
a good solution for this yet. If you have, please let me know 😃

For smaller things like scripts and cronjob this is a non-issue as I can execute
them directly with `docker run` 🙂

I've also created a full example on [this gist][] for your cloning pleasure 😉

Not a primary goal so far was keeping the container size down as the machine has
enough disk space available. That being said, I'll probably try to find a good
balance between the slim and non-slim base images in the future. For other
environments I'm also try to stick with Alpine where ever it makes sense 😊

[this gist]: https://gist.github.com/zerok/64d1ee4dad2c766cd06e4a610aca8f9b
