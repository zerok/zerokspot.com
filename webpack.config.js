const Path = require('path');
const Webpack = require('webpack');
const {VueLoaderPlugin} = require('vue-loader');

module.exports = {
    entry: {
        archive: ['whatwg-fetch', Path.join(__dirname, 'static', 'app', 'archive.js')]
    },
    output: {
        path: Path.resolve('static', 'js'),
        filename: '[name].js'
    },
    devtool: 'source-map',
    mode: 'production',
    module: {
        rules: [
            {
                test: /\.vue$/,
                exclude: /node_modules/,
                use: "vue-loader"
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: "babel-loader"
            },
        ]
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.common.js'
        }
    },
    plugins:[
        new VueLoaderPlugin()
    ]
};
