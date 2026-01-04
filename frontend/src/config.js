// 自定义的前端配置文件
const config = {
    debug: false, // true、false
    app: {
        app_name: "Waigo",
        app_class: "waigo_sv_", // ginthon_sv_ 、waigo_sv_
        app_version: "1.6.4", // 1.0.0
    },
    sys:{
        backend: "go", // go、py
        home_route: "/home", // 主页默认页的路由 ""、"/home"
        ssr: false, // true，服务端渲染+SEO，js的window将对象不可访问。Wails里面false，Gthon里面随意。
        csr: true, // true，页面有js，刷新任何页面，页面有js效果。
    },
    api: {
        js_call_go_url: "http://127.0.0.1:9850/api/js_call_go", // http(s)://127.0.0.1:9750/api/js_call_py 、http(s)://127.0.0.1:9850/api/js_call_go
        api_host: "http://127.0.0.1:9850", // http(s)://127.0.0.1:9750、http(s)://127.0.0.1:9850
    },
};
export default config;