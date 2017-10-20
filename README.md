 [![Build Status](https://travis-ci.org/Arafatk/glot.svg?branch=master)](https://travis-ci.org/Arafatk/glot) [![GoDoc](https://godoc.org/github.com/arafat/glot?status.svg)](https://godoc.org/github.com/Arafatk/glot) [![Join the chat at https://gitter.im/tensorflowrb/Lobby](https://badges.gitter.im/tensorflowrb/Lobby.svg)](https://gitter.im/glot-dev/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)
# Glot
`glot` is a plotting library for Golang built on top of [gnuplot](http://www.gnuplot.info/). `glot` currently supports styles like lines, points, bars, steps, histogram, circle, and many others. We are continuously making efforts to add more features.  

## Documentation
Documentation is available at [godoc](https://godoc.org/github.com/Arafatk/glot).      

## Requirements
 - gnu plot
    - build gnu plot from [source](https://sourceforge.net/projects/gnuplot/files/gnuplot/)
    - linux users
       -  ```sudo apt-get update```
       -  ```sudo apt-get install gnuplot-x11``` 
    - mac users
       -  install homebrew
       -  ```brew cask install xquartz``` (for x-11)
       -  ```brew install gnuplot --with-x11```

## Installation     
```go get github.com/Arafatk/glot```

## Usage and Examples  
We have a blog post explaining our vision and covering some basic usage of the `glot` library. Check it out [here](https://medium.com/@Arafat./introducing-glot-the-plotting-library-for-golang-3133399948a1).

## Examples
![](https://raw.githubusercontent.com/Arafatk/plot/master/Screenshot%20-%20Saturday%2014%20October%202017%20-%2004-51-13%20%20IST.png)

## Contributing
We really encourage developers coming in, finding a bug or requesting a new feature. Want to tell us about the feature you just implemented, just raise a pull request and we'll be happy to go through it. 

