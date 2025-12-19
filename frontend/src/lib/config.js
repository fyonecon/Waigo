// 自定义的前端配置文件
const config = {
    debug: false, // true、false
    app: {
        app_name: "Waigo",
        app_class: "waigo_sv_", // ginthon_sv_ 、waigo_sv_
        app_version: "1.5.0", // 1.0.0
    },
    sys:{
        backend: "go", // go、py
        home_route: "/home" // 主页默认页的路由 “”、"/home”
    },
    api: {
        js_call_go_url: "http://127.0.0.1:9850/api/js_call_go", // http://127.0.0.1:9750/api/js_call_py 、http://127.0.0.1:9850/api/js_call_go
    },
};
export default config;