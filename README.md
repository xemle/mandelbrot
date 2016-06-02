Playground to calculate and browse mandelbrot fractal. The frontent is a
simple web page built with Angular 2. The backend is a simple webserver written
in golang.

# Run Mandelbrot Browser

    $ mkdir -p mandelbrot/src
    $ export GOPATH=`pwd`/mandelbrot
    $ git clone git@github.com:xemle/mandelbrot-browser.git mandelbrot/src/mandelbrot
    $ cd mandelbrot/src/mandelbrot/public
    $ npm install
    $ npm tsc
    $ cd ..
    $ go get
    $ go build
    $ ./mandelbrot

Open your browser a [localhost:8080](http://localhost:8080)
