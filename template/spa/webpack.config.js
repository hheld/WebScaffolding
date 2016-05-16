const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require('webpack');

const config = {
    entry: [
        './src/app/App.js'
    ],
    output: {
        filename: 'index_bundle.js',
        path: __dirname + '/../assets'
    },
    module: {
        loaders: [{
            test: /\.js$/,
            include: __dirname + '/src',
            loader: 'babel-loader'
        }
        ]
    },
    plugins: [
        new HtmlWebpackPlugin({
            title: '{{.AppName}}',
            template: 'src/index.ejs'
        })
    ]
};

if(process.env.DEBUG) {
    config.devtool = 'eval-source-map';
} else {
    config.plugins.push(new webpack.DefinePlugin({
        'process.env': {
            'NODE_ENV': JSON.stringify('production')
        }
    }));
}

module.exports = config;
