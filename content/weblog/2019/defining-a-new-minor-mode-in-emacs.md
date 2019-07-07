---
title: "Defining a new minor-mode in Emacs"
date: 2019-07-07T14:55:30+02:00
tags:
- emacs
---

Recently, I wanted to create my first little mode for Emacs which would allow me to add additional keybindings and commands to handle Kubernetes Secret definitions more efficiently. Turns out, while it’s quite simple to create new major modes and also to activate it, with minor-modes the whole process is a bit different. The working result of this whole quest can be found as the [k8s-secret-mode](https://github.com/zerok/k8s-secret-mode.el) package on Github but here I just wanted to write all things I learnt on the way about working with minor-modes in general and creating a new one in particular.

## Defining a minor-mode

Since this was my first attempt at doing something like that, I first needed to learn how to create a minor-mode. Turns out, there are already a couple of helpers available as part of the easy-mmode package, first of all the `define-minor-mode` macro:

```
(define-minor-mode k8s-secret-mode
  "Toggle k8s-secret-mode."
  
  ; Here I can set a custom keymap which should be enabled
  ; with this minor mode.
  :keymap k8s-secret-mode-map
  
  ; The text "k8s-secret" should be shown inside the
  ; minor-mode status line:
  :lighter " k8s-secret"
  
  ; Inside the mode I want to keep track if the buffer has
  ; already been decoded or not. I'm using a buffer-local
  ; variable for that which should be reset every time 
  ; the mode is enabled.
  (when k8s-secret-mode
    (setq k8s-secret--status nil)))
```

## Detecting an active minor-mode

Next, I wanted to find out, if a certain minor mode is active or not. Sadly, there isn't something like the `major-mode` variable. Instead, I opted to parse the minor mode out of the `minor-mode-alist` which holds all status-line activation functions for all known minor-modes:

```
(require 'dash)

(defun active-minor-modes ()
  "Returns a list of all active minor modes based on the state of the mode-line."
  (interactive)
  (let* ((svars (-map '-first-item minor-mode-alist))
         (actives (-filter 'symbol-value svars)))
    (progn
      (when (called-interactively-p 'any)
        (message "Active minor modes: %s" actives))
      actives)))
```

This is basically a list of cons-lists with the car being the status variable of a minor mode and the cdr the method or string that should be used for the status line. It's not perfect but usually the minor-mode variable has the same name as the minor-mode itself, so it's enough for my scenario.

## How to automatically activate a minor-mode

While acting a major-mode usually happens through the `auto-mode-alist` and `magic-mode-alist`, something similar for minor-modes is not available in the standard distribution.

You could hook into things like the `find-file-hook` and implement a `set-file-mode` like function ourselves. Luckily, though, Joe Wreschnig has already done all the work for us with his [auto-minor-mode](https://github.com/joewreschnig/auto-minor-mode) package.

```
(require 'auto-minor-mode)
(setq auto-minor-mode-magic-alist
      (list '(k8s-secret-mode-magic-matches . 
              k8s-secret-mode)))
```


All the functions and matching taking place inside the auto-minor-mode-magic-alist happens in a narrowed context controlled through the `magic-mode-regexp-match-limit` variable. This defines that the matching is restricted to the first 4000 characters of the buffer. Since in my case the `kind: Secret` string might be also at the very end of a YAML file, I have to first widen and eventually re-narrow the buffer for reliable matching inside the `k8s-secret-mode-magic-matches` function.

## That's it!

That's pretty much it 🙂 Defining a custom (minor-)mode turned out to be much easier than I had anticipated thanks to the `easy-mmode` package! 

