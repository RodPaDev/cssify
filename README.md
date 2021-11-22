# CSSify

This project is written in go and will turn any png to a `HTML` page.
It's not a very efficient way of rendering an image but it was fun a project to learn go.

I learned basic things like setting up the workspace, working with `packages`, `pointers`, `interfaces` and at some point I even tried to extract the colors with `goroutines` but encountered a racing problem and the image's pixels wouldn't be written in order.

I also tried a few html packages but I settled with good old `html/template`, a little confusing coming from the traditional Web Development but good practice.

**Note: Even though the html is properly generated, the biggest image I got to render fast enough was 512x512**

## **Usage:**

```
cssify <path-to-image>
```

To run the source code: 

```
go run . <path-to-image>
```