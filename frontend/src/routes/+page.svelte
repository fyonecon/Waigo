<script>
import {Events} from "@wailsio/runtime";
import {AppServicesForWindow} from "../../bindings/datathink.top.Waigo/internal/bootstrap";

let name = '';
let result = 'Please enter your name below 👇';
let time = 'Listening for Time event...';

// js调用Go
function JSCallGo(key, data_dict){
    return new Promise(resolve => {
      AppServicesForWindow.JSCallGo(key, data_dict).then((resultValue) => {
            resolve(resultValue);
        }).catch((error) => {
            console.error("JSCallGo=Error=", error);
            resolve({
                "state": 0,
                "msg": "JSCallGo出错",
                "content": {
                    "key": key,
                    "data_dict": data_dict,
                    "error": error,
                },
            });
        });
    });
}

//
const key1 = "test";
const data_dict = {
  "data1": "11",
};
JSCallGo(key1, data_dict).then(result=>{
    console.log("JSCallGo="+key1+"=", result);
});


</script>

<div class="container">
    <h2>Page</h2>
</div>
