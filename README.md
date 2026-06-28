# BANER_MAKER

A CLI utility written in **Go** designed to automatically prepare, resize, and convert custom banners, emblems, and army/chapter logos for the **Warhammer 40,000** game series (such as the *Dawn of War* series).

The program processes source images, scales them to match strict in-game resolution requirements, applies necessary mirroring/flipping transformations for proper texture mapping, and exports them directly into the required **.tga** (Truevision TGA) format.

## ✨ Features
* 📷 **Multi-format Input:** Full support for popular input formats like JPEG and PNG.
* 📐 **Automated Resizing:** Rescales images to strict dimensions required by game engines (e.g., 64x96, 64x64).
* 🔄 **Texture Orientation Fix:** Automatic horizontal and vertical flipping to ensure textures render correctly on in-game models.
* 💾 **TGA Export:** Seamless conversion to optimized `.tga` files.

## 📂 Project Structure

```text
BANER_MAKER/
├── convertor/
│   └── convertor.go        # TGA conversion and file-saving logic
├── images/
│   ├── input/              # Place your source images here (jpg, png)
│   └── output/             # Processed .tga banners are saved here
├── resize/
│   └── image_resize.go     # Resizing, flipping, and transformation logic
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── main.go                 # Application entry point
└── README.md
```
## 🚀 Quick Start
1. **Requirements**
Ensure you have [Go (Golang)](https://go.dev/dl/) version 1.16 or higher installed.
2.  **Install Dependencies**
Clone the repository and install the required `imaging` library using:
```bash
go mod tidy
```
3. **Usage**
   1. Place your raw images (e.g., `space_marines.png` or `ork_logo.jpg`) into the images/input/ directory.
   2. Run the application from the root directory: 
   ```bash 
    go run main.go
    ```
    3. Retrieve your game-ready `.tga` files from the `images/output/` directory. They will already be properly sized, flipped, and formatted.

## ⚠️ Important Developer Note (Troubleshooting)
If the program encounters a `tga: invalid format` error while trying to process standard `.jpg` or `.png` files, it means Go's image decoder does not recognize the input format at runtime and mistakenly passes it to the TGA decoder.
