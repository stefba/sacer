[<img src="static/svg/logo/sacferal-c.svg" alt="Sacer Feral" width=900>](https://github.com/stefba/sacer)

This application is the website and archive system of Sacer Feral. It uses a flat-file database and makes heavy use of a directory tree structure. I share it here for educational purposes only.

## CSS

Compile [Less](http://lesscss.org/) files:

`lessc --clean-css ./css/main.less ./css/dist/main.css`

## JS

Build JavaScript bundle:

`yarn --cwd ./js build`
