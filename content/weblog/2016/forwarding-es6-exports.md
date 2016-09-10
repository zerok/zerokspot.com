---
date: '2016-08-04T20:26:59+02:00'
language: en
tags:
- javascript
- es6
title: Forwarding ES6 exports
---

Recently I stumbled upon a little feature within ES6's [export statement][] that
helps greatly with creating `index.js` files within a package. When you try to
organise your modules into packages you usually end up creating `index.js` files
that simply import elements from one sub-module and export them again.

You are probably familiar with the simple export syntax for exporting a local
variable or function to the outside world:

```
export function myFunction() {}
```

But the export-statement can also be used to forward exports from another
module:

```
export {name1, name2} from './somefile';
```

In my concrete situation I had following package:

```
# tree app/actions
app/actions
├── auth.js
├── index.js
├── project.js
└── search.js

0 directories, 4 files
```

Within the application I wanted to import actions like this:

```
import {login} from './actions';
```

Using the export-from statement, the `index.js` for the actions-package now
simply looks like this:

```
export {login, logout} from './auth';
export {addProject, updateProject} from './project';
export {executeSearch} from './search';
```

Compact and readable 😊

[export statement]: https://developer.mozilla.org/en/docs/web/javascript/reference/statements/export
