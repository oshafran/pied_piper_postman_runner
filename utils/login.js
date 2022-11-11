const axios = require("axios");
const https = require("https");

const main = async () => {
  let x_xsrf_token;
  const instance = axios.create({
    httpsAgent: new https.Agent({
      rejectUnauthorized: false,
    }),
  });
  let cookie;
  try {
    const res = await instance.post(
      `https://${process.env.VMANAGEIP}:${process.env.PORT}/j_security_check`,
      new URLSearchParams({
        j_username: process.env.J_USERNAME, //gave the values directly for testing
        j_password: process.env.J_PASSWORD,
      })
    );
    cookie = res.headers["set-cookie"][0].split(";")[0];

    console.log("RES: ");
    const res2 = await instance.get(
      `https://${process.env.VMANAGEIP}/dataservice/client/token`,
      {
        headers: {
          Cookie: cookie,
        },
      }
    );
    process.env["X-XSRF-TOKEN"] = res.data;
    console.log("RES2: ", res2.data);
    x_xsrf_token = res2.data;
    // instance.headers.common["X-XSRF-TOKEN"] = res.data;
  } catch (err) {
    console.log("ERR: ", err);
  }
  return { instance, x_xsrf_token, cookie };
};

module.exports = main;
