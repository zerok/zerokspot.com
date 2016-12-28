const Path = require('path');
const Webpack = require('webpack');

module.exports = {
    entry: {
        archive: ['whatwg-fetch', Path.join(__dirname, 'static', 'app', 'archive.js')]
    },
    output: {
        path: Path.join('static', 'js'),
        filename: '[name].js'
    },
    devtool: 'source-map',
    plugins: [
        new Webpack.DefinePlugin({
            'process.env': {
                'NODE_ENV': JSON.stringify('production')
            }
        }),
        new Webpack.optimize.UglifyJsPlugin({
            compress: {
                warnings: false
            }
        }),
        new Webpack.optimize.OccurenceOrderPlugin(),
        new Webpack.optimize.DedupePlugin(),
    ],
    module: {
        loaders: [
            {
                test: /\.vue$/,
                exclude: /node_modules/,
                loader: "vue"
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: "babel",
                query: {
                    presets: ['es2015']
                }
            },
        ]
    },
    resolve: {
        alias: {
            'vue$': 'vue/dist/vue.common.js'
        }
    }
};
