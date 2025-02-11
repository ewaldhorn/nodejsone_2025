const HtmlWebpackPlugin = require("html-webpack-plugin");
const InlineSourcePlugin = require("inline-source-webpack-plugin");
const WebpackShellPluginNext = require("webpack-shell-plugin-next");
const path = require("path");

module.exports = {
  entry: "./index.js",
  mode: "production",
  output: {
    path: path.resolve(__dirname, "./dist"),
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "index.html",
      filename: "index.html",
      inject: false,
    }),
    new InlineSourcePlugin({
      compress: true,
      rootlist: ["index.html"],
      match: /\.(js|css)$/,
    }),
    new WebpackShellPluginNext({
      onBuildEnd: {
        scripts: ["rm ./dist/main.js"],
        blocking: false,
        parallel: true,
      },
    }),
  ],
};
