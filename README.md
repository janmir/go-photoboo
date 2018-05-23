# Go-PhotoBoo
> Go lang based photobooth software for Olympus cameras compatible with Olympus Image Share application

# Modules
- [ ] Interface to camera module
- [x] Image generation module
- [x] SNS Share and Upload module
    - [ ] Facebook Graph API
    - [x] Dropbox upload (github.com/stacktic/dropbox)
    - [ ] Twitter
- [x] QR Generation Module (use https://github.com/skip2/go-qrcode)
- [ ] Web Server

# Features
- [ ] live view
- [ ] live view frame
- [ ] Single photo frame
- [ ] Multiple photo grid (optional)
- [ ] Auto image generate
- [ ] Check internet connection for features like upload
- [ ] Auto image upload to facebook
- [ ] Show generated image
- [ ] Show QR code for link
- [x] wifi -> camera, ethernet -> internet

# Flow
1. START: Choose layout
2. Prompt for shutter release...Countdown
3. Shoot
4. Shoot more if needed by layout (optional)
5. Process...add the chosen frame
6. Upload
7. Make QR from upload link
8. Display picture and QR
9. Count to 10, then go to START: