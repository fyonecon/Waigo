import {dev} from '$app/environment';
import config from "../config.js";

// we don't need any JS on this page, though we'll load
// it in dev so that we get hot module replacement
export const csr = config.sys.csr; // dev页面无js、true页面有js

// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender = true;

//
export const ssr = config.sys.ssr;  // 单独设置SEO，true开启 服务端渲染+SEO