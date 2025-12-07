<script>
import {Events} from "@wailsio/runtime";
import {AppServicesForWindow} from "../../bindings/datathink.top.Waigo/internal/bootstrap";

let name = '';
let result = 'Please enter your name below 👇';
let time = 'Listening for Time event...';

//
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
  <div>
    <span data-wml-openURL="https://wails.io">
      <img src="/wails.png" class="logo" alt="Wails logo"/>
    </span>
    <span data-wml-openURL="https://svelte.dev">
      <img src="/svelte.svg" class="logo svelte" alt="Svelte logo"/>
    </span>
  </div>
  <h1>Wails + Svelte</h1>
  <div class="result">{result}</div>
  <div class="card">
    <div class="input-box">
      <input class="input" bind:value={name} type="text" autocomplete="off"/>
    </div>
  </div>
  <div class="footer">
    <div><p>Click on the Wails logo to learn more</p></div>
    <div><p>{time}</p></div>
  </div>
</div>

<style>
</style>
