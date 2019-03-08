---
title: "Reusing Jupyter-Notebook tokens"
date: "2019-03-08T17:25:36+01:00"
tags:
- jupyter
---

In order to make setting up [Jupyter](https://jupyter.org/) Notebook
easier, I've recently created a simple Docker image for it which
coworkers can use without having to install Python et al. Without any
kind of wrapper script around it, it gets tedious to enter a different
login token every time you initialize a new notebook instance or just
restart the same after a reboot.

Luckily, you can fixate the token in jupyter's configuration file:

```
$ echo "c.NotebookApp.token = 'something'" > jupyter_notebook_config.py
$ docker run -i -t -p 8888:8888 \
  -v $PWD:/var/notebooks \
  -v $PWD/jupyter_notebook_config.py:/root/.jupyter/jupyter_notebook_config.py \
  reg.company.com/jupyter:latest
```

Opening up the notebook will now accept "something" as token.
