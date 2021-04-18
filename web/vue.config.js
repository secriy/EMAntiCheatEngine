const webpack = require('webpack');
module.exports = {
    publicPath: './',
    assetsDir: 'static',
    productionSourceMap: false,
    configureWebpack: {
        plugins: [
            new webpack.ProvidePlugin({
                $: 'jquery',
                jQuery: 'jquery'
            })
        ]
    },
    devServer: {
        proxy: {
            '^/api': {
                target: 'http://127.0.0.1:3000',
                ws: true,
                changeOrigin: true
            }
        }
    }
};
