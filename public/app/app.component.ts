import { Component, ViewChild, AfterViewInit } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import './rxjs-operators';

@Component({
  selector: 'my-app',
  styles: [`
    .container {
      display: block;
      width: 100%;
    }
  `],
  template: `<h1>Mandelbrot2Go</h1>
  <div #container clas="container">
    <canvas #canvas [attr.width]="width" [attr.height]="height"></canvas>
  </div>`
})

export class AppComponent {
  @ViewChild("container") container;
  @ViewChild("canvas") canvas;

  top:number = 2;
  left:number = -2;
  bottom:number = -2;
  right:number = 2;

  width: number = 100;
  height: number = 100;
  tileSize: number = 100;
  maxIteration = 99;

  constructor(private http: Http) {
  }

  set(centerX, centerY, size) {
    var d = size / Math.min(this.width, this.height);
    this.top = centerY + (this.height / 2) * d;
    this.left = centerX - (this.width / 2) * d;
    this.bottom = centerY - (this.height / 2) * d;
    this.right = centerX + (this.width / 2) * d;
  }

  ngAfterViewInit() {
    let container = this.container.nativeElement;
    this.width = container.offsetWidth;
    this.height = 500;
    let canvas = this.canvas.nativeElement;
    this.context = canvas.getContext("2d");

    this.set(-0.5, 0, 0.2);
    this.calculateTiles();
  }

  calculateTiles() {
    var y:number = 0;
    var dx:number = (this.right - this.left) / this.width;
    var dy:number = (this.top - this.bottom) / this.height;
    while (y < this.height) {
      var x:number = 0;
      while (x < this.width) {
        let w = Math.min(x + this.tileSize, this.width) - x;
        let h = Math.min(y + this.tileSize, this.height) - y;

        let l = this.left + dx * x;
        let t = this.top - dy * y;
        let r = l + dx * w;
        let b = t - dy * h;

        this.calculateTile(t, l, b, r, w, h);

        x += this.tileSize;
      }
      y += this.tileSize;
    }
  }

  calculateTile(top, left, bottom, right, width, height) {
    let command = {
      top: top, left: left, bottom: bottom, right: right,
      width: width, height: height,
      maxIteration: this.maxIteration
    }
    console.log("calculate tile", command);
    this.http.put('tile', JSON.stringify(command))
      .map(res => res.json())
      .subscribe(tile => this.addTile(tile));
  }

  addTile(tile) {
    let tileDx = (tile.right - tile.left) / tile.width;
    let tileDy = (tile.top - tile.bottom) / tile.height;

    let dx = (this.right - this.left) / this.width;
    let dy = (this.top - this.bottom) / this.height;

    let x = (tile.left - this.left) / dx;
    let y = (this.top - tile.top) / dy;

    console.log("Got tile at x=" + x + ", y=" + y, tile);
    let imageData = this.context.getImageData(x, y, tile.width, tile.height);

    imageData = this.drawTile(tile, imageData);

    this.context.putImageData(imageData, x, y);
  }

  drawTile(tile, imageData) {
    let len = Math.min(tile.iterations.length, imageData.data.length / 4);
    let i = 0;
    let hue = 1 / tile.maxIteration;
    while (i < len) {
      let iteration = tile.iterations[i];

      let x = i % tile.width;
      let y = (i - x) / tile.width;

      let p = i * 4;

      let rgb = [0, 0, 0];
      if (iteration < tile.maxIteration) {
        rgb = this.hslToRgb(hue * iteration, 0.5, 0.5);
      }
      imageData.data[p] = rgb[0];
      imageData.data[p + 1] = rgb[1];
      imageData.data[p + 2] = rgb[2];
      imageData.data[p + 3] = 255;

      i++;
    }
    return imageData;
  }

   /**
   * Converts an HSL color value to RGB. Conversion formula
   * adapted from http://en.wikipedia.org/wiki/HSL_color_space.
   * Assumes h, s, and l are contained in the set [0, 1] and
   * returns r, g, and b in the set [0, 255].
   *
   * @param   {number}  h       The hue
   * @param   {number}  s       The saturation
   * @param   {number}  l       The lightness
   * @return  {Array}           The RGB representation
   */
  hslToRgb(h, s, l){
    var r, g, b;

    if (s == 0) {
        r = g = b = l; // achromatic
    } else {
        var hue2rgb = function hue2rgb(p, q, t) {
            if(t < 0) t += 1;
            if(t > 1) t -= 1;
            if(t < 1/6) return p + (q - p) * 6 * t;
            if(t < 1/2) return q;
            if(t < 2/3) return p + (q - p) * (2/3 - t) * 6;
            return p;
        }

        var q = l < 0.5 ? l * (1 + s) : l + s - l * s;
        var p = 2 * l - q;
        r = hue2rgb(p, q, h + 1/3);
        g = hue2rgb(p, q, h);
        b = hue2rgb(p, q, h - 1/3);
    }

    return [Math.round(r * 255), Math.round(g * 255), Math.round(b * 255)];
  }

  tick() {
    requestAnimationFrame(()=> {
      this.tick()
    });

    var ctx = this.context;
    ctx.clearRect(0, 0, 400, 400);
    ctx.fillStyle = "#FF0000";
    ctx.fillRect(0, 0, 100, 100);
  }
}
