const path = require('path')
const sassResourcesLoader = require('craco-sass-resources-loader');

module.exports = {
    plugins: [
        {
            plugin: sassResourcesLoader,
            options: {
                resources: './src/theme/index.scss',
            },
        },
    ],
    webpack: {
        //@符号作为src文件
        alias: {
            '@': path.join(__dirname, 'src')
        },
    }
}