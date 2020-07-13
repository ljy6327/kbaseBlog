// 在vue-config.js 中加入
const isProduction = process.env.NODE_ENV === 'production';

module.exports = {
    devServer: {
        port: 8080,     // 端口
    },
     // 配置webpack
    configureWebpack: config => {
        if (isProduction) {
            // 开启分离js
            config.optimization = {
                runtimeChunk: 'single',
                splitChunks: {
                chunks: 'all',
                maxInitialRequests: Infinity,
                minSize: 20000,
                cacheGroups: {
                    vendor: {
                    test: /[\\/]node_modules[\\/]/,
                    name (module) {
                        // get the name. E.g. node_modules/packageName/not/this/part.js
                        // or node_modules/packageName
                        const packageName = module.context.match(/[\\/]node_modules[\\/](.*?)([\\/]|$)/)[1]
                        // npm package names are URL-safe, but some servers don't like @ symbols
                        return `npm.${packageName.replace('@', '')}`
                    }
                    }
                }
                }
            };
        }
    }
};