/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "tess-personal",
      removal: input?.stage === "production" ? "retain" : "remove",
      protect: ["production"].includes(input?.stage),
      home: "aws",
    };
  },
  async run() {
    const func = new sst.aws.Function("TessFunc", {
      handler: "./cmd",
      runtime: 'go',
      url: true,
    });

    new sst.aws.Router("TessRouter", {
      domain: {
        name: 'twflyfishing.com',
        redirects: ['www.twflyfishing.com']
      },
      routes: {
        "/*": func.url,
      }
    })
  },
});
