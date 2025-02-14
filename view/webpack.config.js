const path = require("path");
const { VueLoaderPlugin } = require("vue-loader");

module.exports = {
    entry: "./vue/app.js",
    output: {
        path: path.resolve(__dirname, "public"),
        filename: "bundle.js",
        publicPath: "/",
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: "vue-loader",
            },
            {
                test: /\.css$/,
                use: [
                    'vue-style-loader',
                    'css-loader'
                ]
            }
        ],
    },
    plugins: [
        new VueLoaderPlugin(),
    ],
    resolve: {
        alias: {
            vue$: 'vue/dist/vue.esm-bundler.js'
        },
        extensions: ['.js', '.vue', '.json']
    },
    devServer: {
        static: path.join(__dirname, "public"),
        compress: true,
        port: 8080,
        hot: true,
    },
};
