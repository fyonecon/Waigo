// 全局

import config from "../config.js";

export const csr = config.sys.csr; // 确保客户端渲染
export const ssr = config.sys.ssr; // false关闭服务端渲染，false关闭SEO。（false可以让js正确访问“window”对象）

