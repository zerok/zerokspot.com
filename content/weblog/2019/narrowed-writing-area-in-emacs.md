---
title: "Narrowed writing area in Emacs"
date: "2019-12-13T19:25:00+01:00"
tags:
- emacs
- writing
---

When I'm on the go I tend to write blog-posts in [iA Writer][i] which has a couple of nice features:

1. It narrows down the content horizontally but doesn't do hard linebreaks.
2. It centers the content area.

I wanted to have that in Emacs, too, for when I'm writing Markdown files there. Especially on high-res displays the rendering area is simply too large. Sure, I could simply split the frame and construct a window that is just at the right width, but that's weird. Perhaps there are ways to just do soft-wrapping lines at a certain line-length and then center everything?

For this I'm currently using [`visual-line-mode`][v] in combination with [`set-window-margins`][m]:

```
(defun narrow-editor/setup-fringe ()
  (interactive)
  (let* ((content-width 80)
        (full-width (window-total-width))
        (margin (floor (/ (- full-width content-width) 2))))
    (visual-line-mode +1)
    (set-window-margins (get-buffer-window) margin margin)))

```

There are a couple of smaller hickups primarily when working with splits but so far I like it! Especially after combining it with a little minor-mode and an entry in the `window-size-change-functions` list ;-)

```
(defun narrow-editor/unset-fringe ()
  (interactive)
  (visual-line-mode +1)
  (set-window-margins (get-buffer-window) 0 0))

(define-minor-mode narrow-editor-mode
  "In this mode the buffer will be narrowed down for a better writing experience."
  nil
  " narrow"
  '()
  :group 'narrow-editor)

(defun narrow-editor/init ()
  "Initializes the narrow editor mode for the current buffer."
  (interactive)
  (if narrow-editor-mode
      (narrow-editor/setup-fringe)
    (narrow-editor/unset-fringe)))

(add-hook 'narrow-editor-mode-hook 'narrow-editor/init)

(defun narrow-editor/window-size-change (frame)
  "Windows that have narrow-editor-mode enabled should react to size changes."
  (interactive)
  (dolist (window (window-list))
    (with-current-buffer (window-buffer window)
      (when narrow-editor-mode
        (narrow-editor/setup-fringe)))))


(add-to-list 'window-size-change-functions 'narrow-editor/window-size-change)
```


[i]: https://ia.net/writer
[v]: https://www.gnu.org/software/emacs/manual/html_node/emacs/Visual-Line-Mode.html
[m]: https://www.gnu.org/software/emacs/manual/html_node/elisp/Display-Margins.html
