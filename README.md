vkg
===

Vkg is the plugin manager for Vim. The name vkg is a [portmanteau](http://en.wikipedia.org/wiki/Portmanteau)
for Vim package.

installation
------------
First method, with ```go get```:
```bash
go get github.com/pepegar/vkg
```

Second method, just search the latest release in the [releases page](https://github.com/pepegar/vkg/releases).

features
--------
* install from github url:
```bash
vkg install github.com/tpope/fugitive
```

* or from just an user and password
```bash
vkg install tpope/fugitive
```

* or better, search the vimawesome index to find the plugin you are looking for:
```bash
vkg search less

# Output:
#  * youcompleteme - A code-completion engine for Vim
#  * vim-less-safe-and-sound - vim syntax for LESS (dynamic CSS)
#  * vim-coloresque - css/less/sass/html color preview for vim
#  * vim-less-state-of-grace - LessCSS Syntax support in Vim
#  * vim-less - Vim support for LESS CSS
#  * less-vim - Less is More - A minimal color scheme
#  * vim-less-better-than-revenge - Vim support for LESS CSS
#  * vim-lesscss - Update corresponding css files on the fly while edit less files
#  * mediummode - Vim Medium Mode - like hard mode, but less...hard
#  * vim-less-come-back-be-here - Vim Less syntax highlighting
#  * vim-less-autocompile - vim auto compile when edit
#  * vim-coloresque-fearless - css/less/sass/html color preview for vim
#  * jasmine - DOM-less simple JavaScript testing framework
#  * fecompressor-vim - Automatic executing YUI Compressor, Closure Compiler, UglifyJS, SASS, LESS
#  * darkerdesert - A darker version of desert scheme, less radiation
#  * less-bat - Now less is more on Windows, too
#  * snippets-ours - snippets for vim-snipmate, the original one is too rich for me
#  * vim-less-colorscheme - Colorscheme for vim based off of less.org's example code section
#  * vim-less-back-to-december -
#  * vim-less-cmd - Vim commands for working with Less css metalanguage)

vkg install less-vim
```

* list your plugins:
```bash
vkg list
```

* uninstall them
```bash
vkg uninstall fugitive
```

* generate a valid vkgrc file
```bash
vkg freeze
```

* install all plugins from vkgrc
```
vkg install
```
