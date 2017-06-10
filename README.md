# Filigrane
A dead simple image watermaker made with Go, with no dependencies.

Note: this does not work correctly with JPEGs rotated via EXIF.

## Usage
Place a `watermark.png` along the executable and call `filigrane` with any JPEG
you want watermarked as an argument.

Drag-and-dropping files on the executable or a shortcut to the executable
should also work.

`filigrane_.jpg` files will be saved as JPEG (quality set to 80).

## License
Licensed under the [WTFPL](LICENSE).
