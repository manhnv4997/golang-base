const domain = "https://7f39-14-191-163-181.ngrok-free.app"

document.addEventListener("DOMContentLoaded", function () {
    // Tạo thẻ div để vue.js mount vào
    let appDiv = document.createElement("div")
    appDiv.id = "vue-app"
    document.body.appendChild(appDiv);

    // Tạo script để nhúng vue.js từ Webpack build
    let script = document.createElement("script");
    script.src = `${domain}/bundle.js`; // Đường dẫn tới file build của Webpack
    document.body.appendChild(script);
})